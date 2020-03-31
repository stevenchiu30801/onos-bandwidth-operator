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
func (in *Basic) DeepCopyInto(out *Basic) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Basic.
func (in *Basic) DeepCopy() *Basic {
	if in == nil {
		return nil
	}
	out := new(Basic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	out.GeneralProvider = in.GeneralProvider
	out.PiPipeconf = in.PiPipeconf
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make(map[string]Port, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Basic = in.Basic
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
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
func (in *GeneralProvider) DeepCopyInto(out *GeneralProvider) {
	*out = *in
	out.P4Runtime = in.P4Runtime
	out.Thrift = in.Thrift
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralProvider.
func (in *GeneralProvider) DeepCopy() *GeneralProvider {
	if in == nil {
		return nil
	}
	out := new(GeneralProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnosDeviceNetcfg) DeepCopyInto(out *OnosDeviceNetcfg) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnosDeviceNetcfg.
func (in *OnosDeviceNetcfg) DeepCopy() *OnosDeviceNetcfg {
	if in == nil {
		return nil
	}
	out := new(OnosDeviceNetcfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnosDeviceNetcfg) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnosDeviceNetcfgList) DeepCopyInto(out *OnosDeviceNetcfgList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OnosDeviceNetcfg, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnosDeviceNetcfgList.
func (in *OnosDeviceNetcfgList) DeepCopy() *OnosDeviceNetcfgList {
	if in == nil {
		return nil
	}
	out := new(OnosDeviceNetcfgList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnosDeviceNetcfgList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnosDeviceNetcfgSpec) DeepCopyInto(out *OnosDeviceNetcfgSpec) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make(map[string]Device, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnosDeviceNetcfgSpec.
func (in *OnosDeviceNetcfgSpec) DeepCopy() *OnosDeviceNetcfgSpec {
	if in == nil {
		return nil
	}
	out := new(OnosDeviceNetcfgSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnosDeviceNetcfgStatus) DeepCopyInto(out *OnosDeviceNetcfgStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnosDeviceNetcfgStatus.
func (in *OnosDeviceNetcfgStatus) DeepCopy() *OnosDeviceNetcfgStatus {
	if in == nil {
		return nil
	}
	out := new(OnosDeviceNetcfgStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnosQueueNetcfg) DeepCopyInto(out *OnosQueueNetcfg) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnosQueueNetcfg.
func (in *OnosQueueNetcfg) DeepCopy() *OnosQueueNetcfg {
	if in == nil {
		return nil
	}
	out := new(OnosQueueNetcfg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnosQueueNetcfg) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnosQueueNetcfgList) DeepCopyInto(out *OnosQueueNetcfgList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OnosQueueNetcfg, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnosQueueNetcfgList.
func (in *OnosQueueNetcfgList) DeepCopy() *OnosQueueNetcfgList {
	if in == nil {
		return nil
	}
	out := new(OnosQueueNetcfgList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OnosQueueNetcfgList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnosQueueNetcfgSpec) DeepCopyInto(out *OnosQueueNetcfgSpec) {
	*out = *in
	if in.QueueDevices != nil {
		in, out := &in.QueueDevices, &out.QueueDevices
		*out = make(map[string]Queues, len(*in))
		for key, val := range *in {
			var outVal []Queue
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Queues, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnosQueueNetcfgSpec.
func (in *OnosQueueNetcfgSpec) DeepCopy() *OnosQueueNetcfgSpec {
	if in == nil {
		return nil
	}
	out := new(OnosQueueNetcfgSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnosQueueNetcfgStatus) DeepCopyInto(out *OnosQueueNetcfgStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnosQueueNetcfgStatus.
func (in *OnosQueueNetcfgStatus) DeepCopy() *OnosQueueNetcfgStatus {
	if in == nil {
		return nil
	}
	out := new(OnosQueueNetcfgStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *P4Runtime) DeepCopyInto(out *P4Runtime) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new P4Runtime.
func (in *P4Runtime) DeepCopy() *P4Runtime {
	if in == nil {
		return nil
	}
	out := new(P4Runtime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PiPipeconf) DeepCopyInto(out *PiPipeconf) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PiPipeconf.
func (in *PiPipeconf) DeepCopy() *PiPipeconf {
	if in == nil {
		return nil
	}
	out := new(PiPipeconf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queue) DeepCopyInto(out *Queue) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queue.
func (in *Queue) DeepCopy() *Queue {
	if in == nil {
		return nil
	}
	out := new(Queue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Queues) DeepCopyInto(out *Queues) {
	{
		in := &in
		*out = make(Queues, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queues.
func (in Queues) DeepCopy() Queues {
	if in == nil {
		return nil
	}
	out := new(Queues)
	in.DeepCopyInto(out)
	return *out
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Thrift) DeepCopyInto(out *Thrift) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Thrift.
func (in *Thrift) DeepCopy() *Thrift {
	if in == nil {
		return nil
	}
	out := new(Thrift)
	in.DeepCopyInto(out)
	return out
}
