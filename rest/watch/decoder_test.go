/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package versioned_test

import (
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/spotmaxtech/k8s-api-v02217/core/v1"
	apiequality "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/api/equality"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/apis/meta/v1"
	"github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/runtime"
	runtimejson "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/runtime/serializer/json"
	"github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/runtime/serializer/streaming"
	"github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/util/wait"
	"github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/watch"
	"github.com/spotmaxtech/k8s-client-go-v02217/kubernetes/scheme"
	restclientwatch "github.com/spotmaxtech/k8s-client-go-v02217/rest/watch"
)

// getDecoder mimics how k8s.io/client-go/rest.createSerializers creates a decoder
func getDecoder() runtime.Decoder {
	jsonSerializer := runtimejson.NewSerializer(runtimejson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, false)
	directCodecFactory := scheme.Codecs.WithoutConversion()
	return directCodecFactory.DecoderToVersion(jsonSerializer, v1.SchemeGroupVersion)
}

func TestDecoder(t *testing.T) {
	table := []watch.EventType{watch.Added, watch.Deleted, watch.Modified, watch.Error, watch.Bookmark}

	for _, eventType := range table {
		out, in := io.Pipe()

		decoder := restclientwatch.NewDecoder(streaming.NewDecoder(out, getDecoder()), getDecoder())
		expect := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "foo"}}
		encoder := json.NewEncoder(in)
		eType := eventType
		go func() {
			data, err := runtime.Encode(scheme.Codecs.LegacyCodec(v1.SchemeGroupVersion), expect)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}
			event := metav1.WatchEvent{
				Type:   string(eType),
				Object: runtime.RawExtension{Raw: json.RawMessage(data)},
			}
			if err := encoder.Encode(&event); err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			in.Close()
		}()

		done := make(chan struct{})
		go func() {
			action, got, err := decoder.Decode()
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}
			if e, a := eType, action; e != a {
				t.Errorf("Expected %v, got %v", e, a)
			}
			if e, a := expect, got; !apiequality.Semantic.DeepDerivative(e, a) {
				t.Errorf("Expected %v, got %v", e, a)
			}
			t.Logf("Exited read")
			close(done)
		}()
		<-done

		done = make(chan struct{})
		go func() {
			_, _, err := decoder.Decode()
			if err == nil {
				t.Errorf("Unexpected nil error")
			}
			close(done)
		}()
		<-done

		decoder.Close()
	}
}

func TestDecoder_SourceClose(t *testing.T) {
	out, in := io.Pipe()
	decoder := restclientwatch.NewDecoder(streaming.NewDecoder(out, getDecoder()), getDecoder())

	done := make(chan struct{})

	go func() {
		_, _, err := decoder.Decode()
		if err == nil {
			t.Errorf("Unexpected nil error")
		}
		close(done)
	}()

	in.Close()

	select {
	case <-done:
		break
	case <-time.After(wait.ForeverTestTimeout):
		t.Error("Timeout")
	}
}
