/*
Copyright 2021 The Kubernetes Authors.

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

package v1

import (
	"fmt"

	"github.com/spotmaxtech/k8s-api-v02217/core/v1"
	policy "github.com/spotmaxtech/k8s-api-v02217/policy/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/apis/meta/v1"
	"github.com/spotmaxtech/k8s-apimachinery-v02217/pkg/labels"
	"k8s.io/klog/v2"
)

// PodDisruptionBudgetListerExpansion allows custom methods to be added to
// PodDisruptionBudgetLister.
type PodDisruptionBudgetListerExpansion interface {
	GetPodPodDisruptionBudgets(pod *v1.Pod) ([]*policy.PodDisruptionBudget, error)
}

// PodDisruptionBudgetNamespaceListerExpansion allows custom methods to be added to
// PodDisruptionBudgetNamespaceLister.
type PodDisruptionBudgetNamespaceListerExpansion interface{}

// GetPodPodDisruptionBudgets returns a list of PodDisruptionBudgets matching a pod.
func (s *podDisruptionBudgetLister) GetPodPodDisruptionBudgets(pod *v1.Pod) ([]*policy.PodDisruptionBudget, error) {
	var selector labels.Selector

	list, err := s.PodDisruptionBudgets(pod.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var pdbList []*policy.PodDisruptionBudget
	for i := range list {
		pdb := list[i]
		selector, err = metav1.LabelSelectorAsSelector(pdb.Spec.Selector)
		if err != nil {
			klog.Warningf("invalid selector: %v", err)
			continue
		}

		// Unlike the v1beta version, here we let an empty selector match everything.
		if !selector.Matches(labels.Set(pod.Labels)) {
			continue
		}
		pdbList = append(pdbList, pdb)
	}

	if len(pdbList) == 0 {
		return nil, fmt.Errorf("could not find PodDisruptionBudget for pod %s in namespace %s with labels: %v", pod.Name, pod.Namespace, pod.Labels)
	}

	return pdbList, nil
}
