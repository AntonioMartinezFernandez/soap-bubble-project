/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SoapBubbleMachineSpec defines the desired state of SoapBubbleMachine.
type SoapBubbleMachineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	MachineName string `json:"machineName,omitempty"`
	StartURL    string `json:"startURL,omitempty" validate:"url"`
	StopURL     string `json:"stopURL,omitempty" validate:"url"`
	MakeBubbles bool   `json:"makeBubbles,omitempty"`
}

// SoapBubbleMachineStatus defines the observed state of SoapBubbleMachine.
type SoapBubbleMachineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	MakingBubbles bool `json:"makingBubbles,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SoapBubbleMachine is the Schema for the soapbubblemachines API.
type SoapBubbleMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SoapBubbleMachineSpec   `json:"spec,omitempty"`
	Status SoapBubbleMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SoapBubbleMachineList contains a list of SoapBubbleMachine.
type SoapBubbleMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SoapBubbleMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SoapBubbleMachine{}, &SoapBubbleMachineList{})
}
