// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthSlice) DeepCopyInto(out *BandwidthSlice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthSlice.
func (in *BandwidthSlice) DeepCopy() *BandwidthSlice {
	if in == nil {
		return nil
	}
	out := new(BandwidthSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthSlice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthSliceList) DeepCopyInto(out *BandwidthSliceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BandwidthSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthSliceList.
func (in *BandwidthSliceList) DeepCopy() *BandwidthSliceList {
	if in == nil {
		return nil
	}
	out := new(BandwidthSliceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthSliceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthSliceSpec) DeepCopyInto(out *BandwidthSliceSpec) {
	*out = *in
	if in.Slices != nil {
		in, out := &in.Slices, &out.Slices
		*out = make([]Slice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthSliceSpec.
func (in *BandwidthSliceSpec) DeepCopy() *BandwidthSliceSpec {
	if in == nil {
		return nil
	}
	out := new(BandwidthSliceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthSliceStatus) DeepCopyInto(out *BandwidthSliceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthSliceStatus.
func (in *BandwidthSliceStatus) DeepCopy() *BandwidthSliceStatus {
	if in == nil {
		return nil
	}
	out := new(BandwidthSliceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricConfig) DeepCopyInto(out *FabricConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricConfig.
func (in *FabricConfig) DeepCopy() *FabricConfig {
	if in == nil {
		return nil
	}
	out := new(FabricConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricConfigList) DeepCopyInto(out *FabricConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FabricConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricConfigList.
func (in *FabricConfigList) DeepCopy() *FabricConfigList {
	if in == nil {
		return nil
	}
	out := new(FabricConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricConfigSpec) DeepCopyInto(out *FabricConfigSpec) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConnectPoints != nil {
		in, out := &in.ConnectPoints, &out.ConnectPoints
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricConfigSpec.
func (in *FabricConfigSpec) DeepCopy() *FabricConfigSpec {
	if in == nil {
		return nil
	}
	out := new(FabricConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricConfigStatus) DeepCopyInto(out *FabricConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricConfigStatus.
func (in *FabricConfigStatus) DeepCopy() *FabricConfigStatus {
	if in == nil {
		return nil
	}
	out := new(FabricConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Flow) DeepCopyInto(out *Flow) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Flow.
func (in *Flow) DeepCopy() *Flow {
	if in == nil {
		return nil
	}
	out := new(Flow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Slice) DeepCopyInto(out *Slice) {
	*out = *in
	if in.MinRate != nil {
		in, out := &in.MinRate, &out.MinRate
		*out = new(uint)
		**out = **in
	}
	if in.MaxRate != nil {
		in, out := &in.MaxRate, &out.MaxRate
		*out = new(uint)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(uint)
		**out = **in
	}
	if in.Flows != nil {
		in, out := &in.Flows, &out.Flows
		*out = make([]Flow, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Slice.
func (in *Slice) DeepCopy() *Slice {
	if in == nil {
		return nil
	}
	out := new(Slice)
	in.DeepCopyInto(out)
	return out
}
