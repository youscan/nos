//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Nebuly.ai.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GpuPartitionerConfig) DeepCopyInto(out *GpuPartitionerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
	out.NvidiaDevicePluginConfigMap = in.NvidiaDevicePluginConfigMap
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GpuPartitionerConfig.
func (in *GpuPartitionerConfig) DeepCopy() *GpuPartitionerConfig {
	if in == nil {
		return nil
	}
	out := new(GpuPartitionerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GpuPartitionerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAgentConfig) DeepCopyInto(out *MigAgentConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAgentConfig.
func (in *MigAgentConfig) DeepCopy() *MigAgentConfig {
	if in == nil {
		return nil
	}
	out := new(MigAgentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigAgentConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedObject) DeepCopyInto(out *NamespacedObject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedObject.
func (in *NamespacedObject) DeepCopy() *NamespacedObject {
	if in == nil {
		return nil
	}
	out := new(NamespacedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorConfig) DeepCopyInto(out *OperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorConfig.
func (in *OperatorConfig) DeepCopy() *OperatorConfig {
	if in == nil {
		return nil
	}
	out := new(OperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSlicingAgentConfig) DeepCopyInto(out *TimeSlicingAgentConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ControllerManagerConfigurationSpec.DeepCopyInto(&out.ControllerManagerConfigurationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSlicingAgentConfig.
func (in *TimeSlicingAgentConfig) DeepCopy() *TimeSlicingAgentConfig {
	if in == nil {
		return nil
	}
	out := new(TimeSlicingAgentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TimeSlicingAgentConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
