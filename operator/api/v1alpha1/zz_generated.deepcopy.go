//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
func (in *ModelDeployment) DeepCopyInto(out *ModelDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeployment.
func (in *ModelDeployment) DeepCopy() *ModelDeployment {
	if in == nil {
		return nil
	}
	out := new(ModelDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentList) DeepCopyInto(out *ModelDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentList.
func (in *ModelDeploymentList) DeepCopy() *ModelDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpec) DeepCopyInto(out *ModelDeploymentSpec) {
	*out = *in
	out.Optimization = in.Optimization
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpec.
func (in *ModelDeploymentSpec) DeepCopy() *ModelDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentStatus) DeepCopyInto(out *ModelDeploymentStatus) {
	*out = *in
	out.ModelOptimizationJob = in.ModelOptimizationJob
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentStatus.
func (in *ModelDeploymentStatus) DeepCopy() *ModelDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptimizationSpec) DeepCopyInto(out *OptimizationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptimizationSpec.
func (in *OptimizationSpec) DeepCopy() *OptimizationSpec {
	if in == nil {
		return nil
	}
	out := new(OptimizationSpec)
	in.DeepCopyInto(out)
	return out
}
