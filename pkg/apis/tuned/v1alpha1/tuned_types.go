// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TunedSpec defines the desired state of Tuned
type TunedSpec struct {
	Profiles  []TunedProfile   `json:"profiles"`
	Recommend []TunedRecommend `json:"recommend"`
}

type TunedProfile struct {
	Name *string `json:"name"`
	Data *string `json:"data"`
}

type TunedRecommend struct {
	Profile  *string     `json:"profile"`
	Priority *uint64     `json:"priority"`
	Label    *TunedLabel `json:"label"`
}

type TunedLabel struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

// TunedStatus defines the observed state of Tuned
type TunedStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Tuned is the Schema for the tuneds API
// +k8s:openapi-gen=true
type Tuned struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TunedSpec   `json:"spec,omitempty"`
	Status TunedStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TunedList contains a list of Tuned
type TunedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tuned `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tuned{}, &TunedList{})
}
