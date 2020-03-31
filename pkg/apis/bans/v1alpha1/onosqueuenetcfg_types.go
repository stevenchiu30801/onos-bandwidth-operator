package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Queue
type Queue struct {
	// Port
	Port uint16 `json:"port"`

	// Queue ID
	Qid int `json:"qid"`

	// Priority
	Priority int `json:"priority"`

	// Weight
	Weight int `json:"weight"`
}

// Queues
type Queues []Queue

// OnosQueueNetcfgSpec defines the desired state of OnosQueueNetcfg
type OnosQueueNetcfgSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Queue devices
	QueueDevices map[string]Queues `json:"queueDevices"`
}

// OnosQueueNetcfgStatus defines the observed state of OnosQueueNetcfg
type OnosQueueNetcfgStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OnosQueueNetcfg is the Schema for the onosqueuenetcfgs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=onosqueuenetcfgs,scope=Namespaced
type OnosQueueNetcfg struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OnosQueueNetcfgSpec   `json:"spec,omitempty"`
	Status OnosQueueNetcfgStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OnosQueueNetcfgList contains a list of OnosQueueNetcfg
type OnosQueueNetcfgList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OnosQueueNetcfg `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OnosQueueNetcfg{}, &OnosQueueNetcfgList{})
}
