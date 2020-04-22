package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FabricConfigSpec defines the desired state of FabricConfig
type FabricConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Device list
	Devices []string `json:"devices"`

	// Connect point list
	ConnectPoints map[string]string `json:"connectPoints"`
}

// FabricConfigStatus defines the observed state of FabricConfig
type FabricConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FabricConfig is the Schema for the fabricconfigs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=fabricconfigs,scope=Namespaced
type FabricConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FabricConfigSpec   `json:"spec,omitempty"`
	Status FabricConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FabricConfigList contains a list of FabricConfig
type FabricConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FabricConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FabricConfig{}, &FabricConfigList{})
}
