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
	"github.com/keyval-dev/odigos/common"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsertClusterAttributes) DeepCopyInto(out *InsertClusterAttributes) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsertClusterAttributes.
func (in *InsertClusterAttributes) DeepCopy() *InsertClusterAttributes {
	if in == nil {
		return nil
	}
	out := new(InsertClusterAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InsertClusterAttributes) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsertClusterAttributesList) DeepCopyInto(out *InsertClusterAttributesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InsertClusterAttributes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsertClusterAttributesList.
func (in *InsertClusterAttributesList) DeepCopy() *InsertClusterAttributesList {
	if in == nil {
		return nil
	}
	out := new(InsertClusterAttributesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InsertClusterAttributesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsertClusterAttributesSpec) DeepCopyInto(out *InsertClusterAttributesSpec) {
	*out = *in
	if in.Signals != nil {
		in, out := &in.Signals, &out.Signals
		*out = make([]common.ObservabilitySignal, len(*in))
		copy(*out, *in)
	}
	if in.ClusterAttributes != nil {
		in, out := &in.ClusterAttributes, &out.ClusterAttributes
		*out = make([]OtelAttributeWithValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsertClusterAttributesSpec.
func (in *InsertClusterAttributesSpec) DeepCopy() *InsertClusterAttributesSpec {
	if in == nil {
		return nil
	}
	out := new(InsertClusterAttributesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsertClusterAttributesStatus) DeepCopyInto(out *InsertClusterAttributesStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsertClusterAttributesStatus.
func (in *InsertClusterAttributesStatus) DeepCopy() *InsertClusterAttributesStatus {
	if in == nil {
		return nil
	}
	out := new(InsertClusterAttributesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OtelAttributeWithValue) DeepCopyInto(out *OtelAttributeWithValue) {
	*out = *in
	if in.AttributeStringValue != nil {
		in, out := &in.AttributeStringValue, &out.AttributeStringValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OtelAttributeWithValue.
func (in *OtelAttributeWithValue) DeepCopy() *OtelAttributeWithValue {
	if in == nil {
		return nil
	}
	out := new(OtelAttributeWithValue)
	in.DeepCopyInto(out)
	return out
}
