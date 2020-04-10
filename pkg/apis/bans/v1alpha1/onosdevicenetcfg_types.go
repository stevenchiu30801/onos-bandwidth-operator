package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// P4 Runtime
type P4Runtime struct {
	// IP address
	Ip string `json:"ip"`

	// Device key ID
	DeviceKeyId string `json:"deviceKeyId"`

	// Port, default 50051
	Port uint16 `json:"port"`

	// Device ID, defualt 0
	DeviceId int `json:"deviceId"`
}

// Thrift
type Thrift struct {
	// Port, default 9090
	Port uint16 `json:"port"`
}

// General provider
type GeneralProvider struct {
	// P4 Runtime provider
	P4Runtime P4Runtime `json:"p4runtime"`

	// Thrift provider
	Thrift Thrift `json:"thrift"`
}

// PI pipeconf
type PiPipeconf struct {
	PiPipeconfId string `json:"piPipeconfId"`
}

// Device port
type Port struct {
	// Name
	Name string `json:"name"`

	// Speed
	Speed int `json:"speed"`

	// Enabled
	Enabled bool `json:"enabled"`

	// Number
	Number int `json:"number"`

	// Removed
	Removed bool `json:"removed"`

	// Type
	Type string `json:"type"`
}

// Basic
type Basic struct {
	// Driver
	Driver string `json:"driver"`
}

// ONOS device
type Device struct {
	// General provider
	GeneralProvider GeneralProvider `json:"generalprovider"`

	// PI pipeconf
	PiPipeconf PiPipeconf `json:"piPipeconf"`

	// Device ports
	Ports map[string]Port `json:"ports"`

	// Basic
	Basic Basic `json:"basic"`
}

// OnosDeviceNetcfgSpec defines the desired state of OnosDeviceNetcfg
type OnosDeviceNetcfgSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// ONOS devices
	Devices map[string]Device `json:"devices"`

	// AMF connect point
	AmfConnectPoint string `json:"amfConnectPoint,omitempty"`

	// UPF connect point
	UpfConnectPoint string `json:"upfConnectPoint,omitempty"`
}

// OnosDeviceNetcfgStatus defines the observed state of OnosDeviceNetcfg
type OnosDeviceNetcfgStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OnosDeviceNetcfg is the Schema for the onosdevicenetcfgs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=onosdevicenetcfgs,scope=Namespaced
type OnosDeviceNetcfg struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OnosDeviceNetcfgSpec   `json:"spec,omitempty"`
	Status OnosDeviceNetcfgStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OnosDeviceNetcfgList contains a list of OnosDeviceNetcfg
type OnosDeviceNetcfgList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OnosDeviceNetcfg `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OnosDeviceNetcfg{}, &OnosDeviceNetcfgList{})
}
