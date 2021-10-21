/*
Copyright 2021.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppSetSpec defines the desired state of AppSet
type AppSetSpec struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	Replicas int32 `json:"replicas"`
}

// AppSetStatus defines the observed state of AppSet
type AppSetStatus struct {
	AppNames          []string `json:"appNames,omitempty"`
	AvailableReplicas int32    `json:"availableReplicas,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AppSet is the Schema for the appsets API
type AppSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSetSpec   `json:"spec,omitempty"`
	Status AppSetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AppSetList contains a list of AppSet
type AppSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppSet{}, &AppSetList{})
}
