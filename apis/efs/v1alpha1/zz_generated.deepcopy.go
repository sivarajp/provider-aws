//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPoint) DeepCopyInto(out *AccessPoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPoint.
func (in *AccessPoint) DeepCopy() *AccessPoint {
	if in == nil {
		return nil
	}
	out := new(AccessPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointDescription) DeepCopyInto(out *AccessPointDescription) {
	*out = *in
	if in.AccessPointARN != nil {
		in, out := &in.AccessPointARN, &out.AccessPointARN
		*out = new(string)
		**out = **in
	}
	if in.AccessPointID != nil {
		in, out := &in.AccessPointID, &out.AccessPointID
		*out = new(string)
		**out = **in
	}
	if in.ClientToken != nil {
		in, out := &in.ClientToken, &out.ClientToken
		*out = new(string)
		**out = **in
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.LifeCycleState != nil {
		in, out := &in.LifeCycleState, &out.LifeCycleState
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.PosixUser != nil {
		in, out := &in.PosixUser, &out.PosixUser
		*out = new(PosixUser)
		(*in).DeepCopyInto(*out)
	}
	if in.RootDirectory != nil {
		in, out := &in.RootDirectory, &out.RootDirectory
		*out = new(RootDirectory)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointDescription.
func (in *AccessPointDescription) DeepCopy() *AccessPointDescription {
	if in == nil {
		return nil
	}
	out := new(AccessPointDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointList) DeepCopyInto(out *AccessPointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessPoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointList.
func (in *AccessPointList) DeepCopy() *AccessPointList {
	if in == nil {
		return nil
	}
	out := new(AccessPointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointObservation) DeepCopyInto(out *AccessPointObservation) {
	*out = *in
	if in.AccessPointARN != nil {
		in, out := &in.AccessPointARN, &out.AccessPointARN
		*out = new(string)
		**out = **in
	}
	if in.AccessPointID != nil {
		in, out := &in.AccessPointID, &out.AccessPointID
		*out = new(string)
		**out = **in
	}
	if in.ClientToken != nil {
		in, out := &in.ClientToken, &out.ClientToken
		*out = new(string)
		**out = **in
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.LifeCycleState != nil {
		in, out := &in.LifeCycleState, &out.LifeCycleState
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointObservation.
func (in *AccessPointObservation) DeepCopy() *AccessPointObservation {
	if in == nil {
		return nil
	}
	out := new(AccessPointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointParameters) DeepCopyInto(out *AccessPointParameters) {
	*out = *in
	if in.PosixUser != nil {
		in, out := &in.PosixUser, &out.PosixUser
		*out = new(PosixUser)
		(*in).DeepCopyInto(*out)
	}
	if in.RootDirectory != nil {
		in, out := &in.RootDirectory, &out.RootDirectory
		*out = new(RootDirectory)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.CustomAccessPointParameters.DeepCopyInto(&out.CustomAccessPointParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointParameters.
func (in *AccessPointParameters) DeepCopy() *AccessPointParameters {
	if in == nil {
		return nil
	}
	out := new(AccessPointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointSpec) DeepCopyInto(out *AccessPointSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointSpec.
func (in *AccessPointSpec) DeepCopy() *AccessPointSpec {
	if in == nil {
		return nil
	}
	out := new(AccessPointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointStatus) DeepCopyInto(out *AccessPointStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointStatus.
func (in *AccessPointStatus) DeepCopy() *AccessPointStatus {
	if in == nil {
		return nil
	}
	out := new(AccessPointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreationInfo) DeepCopyInto(out *CreationInfo) {
	*out = *in
	if in.OwnerGid != nil {
		in, out := &in.OwnerGid, &out.OwnerGid
		*out = new(int64)
		**out = **in
	}
	if in.OwnerUid != nil {
		in, out := &in.OwnerUid, &out.OwnerUid
		*out = new(int64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreationInfo.
func (in *CreationInfo) DeepCopy() *CreationInfo {
	if in == nil {
		return nil
	}
	out := new(CreationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomAccessPointParameters) DeepCopyInto(out *CustomAccessPointParameters) {
	*out = *in
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.FileSystemIDRef != nil {
		in, out := &in.FileSystemIDRef, &out.FileSystemIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FileSystemIDSelector != nil {
		in, out := &in.FileSystemIDSelector, &out.FileSystemIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomAccessPointParameters.
func (in *CustomAccessPointParameters) DeepCopy() *CustomAccessPointParameters {
	if in == nil {
		return nil
	}
	out := new(CustomAccessPointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomFileSystemParameters) DeepCopyInto(out *CustomFileSystemParameters) {
	*out = *in
	if in.ProvisionedThroughputInMibps != nil {
		in, out := &in.ProvisionedThroughputInMibps, &out.ProvisionedThroughputInMibps
		*out = new(int64)
		**out = **in
	}
	if in.KMSKeyIDRef != nil {
		in, out := &in.KMSKeyIDRef, &out.KMSKeyIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyIDSelector != nil {
		in, out := &in.KMSKeyIDSelector, &out.KMSKeyIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomFileSystemParameters.
func (in *CustomFileSystemParameters) DeepCopy() *CustomFileSystemParameters {
	if in == nil {
		return nil
	}
	out := new(CustomFileSystemParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomMountTargetParameters) DeepCopyInto(out *CustomMountTargetParameters) {
	*out = *in
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupsRefs != nil {
		in, out := &in.SecurityGroupsRefs, &out.SecurityGroupsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroupsSelector != nil {
		in, out := &in.SecurityGroupsSelector, &out.SecurityGroupsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.FileSystemIDRef != nil {
		in, out := &in.FileSystemIDRef, &out.FileSystemIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FileSystemIDSelector != nil {
		in, out := &in.FileSystemIDSelector, &out.FileSystemIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDRef != nil {
		in, out := &in.SubnetIDRef, &out.SubnetIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomMountTargetParameters.
func (in *CustomMountTargetParameters) DeepCopy() *CustomMountTargetParameters {
	if in == nil {
		return nil
	}
	out := new(CustomMountTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystem) DeepCopyInto(out *FileSystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystem.
func (in *FileSystem) DeepCopy() *FileSystem {
	if in == nil {
		return nil
	}
	out := new(FileSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemDescription) DeepCopyInto(out *FileSystemDescription) {
	*out = *in
	if in.AvailabilityZoneID != nil {
		in, out := &in.AvailabilityZoneID, &out.AvailabilityZoneID
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZoneName != nil {
		in, out := &in.AvailabilityZoneName, &out.AvailabilityZoneName
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.CreationToken != nil {
		in, out := &in.CreationToken, &out.CreationToken
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.FileSystemARN != nil {
		in, out := &in.FileSystemARN, &out.FileSystemARN
		*out = new(string)
		**out = **in
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.LifeCycleState != nil {
		in, out := &in.LifeCycleState, &out.LifeCycleState
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumberOfMountTargets != nil {
		in, out := &in.NumberOfMountTargets, &out.NumberOfMountTargets
		*out = new(int64)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.PerformanceMode != nil {
		in, out := &in.PerformanceMode, &out.PerformanceMode
		*out = new(string)
		**out = **in
	}
	if in.SizeInBytes != nil {
		in, out := &in.SizeInBytes, &out.SizeInBytes
		*out = new(FileSystemSize)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ThroughputMode != nil {
		in, out := &in.ThroughputMode, &out.ThroughputMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemDescription.
func (in *FileSystemDescription) DeepCopy() *FileSystemDescription {
	if in == nil {
		return nil
	}
	out := new(FileSystemDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemList) DeepCopyInto(out *FileSystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FileSystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemList.
func (in *FileSystemList) DeepCopy() *FileSystemList {
	if in == nil {
		return nil
	}
	out := new(FileSystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemObservation) DeepCopyInto(out *FileSystemObservation) {
	*out = *in
	if in.AvailabilityZoneID != nil {
		in, out := &in.AvailabilityZoneID, &out.AvailabilityZoneID
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.CreationToken != nil {
		in, out := &in.CreationToken, &out.CreationToken
		*out = new(string)
		**out = **in
	}
	if in.FileSystemARN != nil {
		in, out := &in.FileSystemARN, &out.FileSystemARN
		*out = new(string)
		**out = **in
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.LifeCycleState != nil {
		in, out := &in.LifeCycleState, &out.LifeCycleState
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NumberOfMountTargets != nil {
		in, out := &in.NumberOfMountTargets, &out.NumberOfMountTargets
		*out = new(int64)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.SizeInBytes != nil {
		in, out := &in.SizeInBytes, &out.SizeInBytes
		*out = new(FileSystemSize)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemObservation.
func (in *FileSystemObservation) DeepCopy() *FileSystemObservation {
	if in == nil {
		return nil
	}
	out := new(FileSystemObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemParameters) DeepCopyInto(out *FileSystemParameters) {
	*out = *in
	if in.AvailabilityZoneName != nil {
		in, out := &in.AvailabilityZoneName, &out.AvailabilityZoneName
		*out = new(string)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
	if in.PerformanceMode != nil {
		in, out := &in.PerformanceMode, &out.PerformanceMode
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ThroughputMode != nil {
		in, out := &in.ThroughputMode, &out.ThroughputMode
		*out = new(string)
		**out = **in
	}
	in.CustomFileSystemParameters.DeepCopyInto(&out.CustomFileSystemParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemParameters.
func (in *FileSystemParameters) DeepCopy() *FileSystemParameters {
	if in == nil {
		return nil
	}
	out := new(FileSystemParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemSize) DeepCopyInto(out *FileSystemSize) {
	*out = *in
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = (*in).DeepCopy()
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
	if in.ValueInIA != nil {
		in, out := &in.ValueInIA, &out.ValueInIA
		*out = new(int64)
		**out = **in
	}
	if in.ValueInStandard != nil {
		in, out := &in.ValueInStandard, &out.ValueInStandard
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemSize.
func (in *FileSystemSize) DeepCopy() *FileSystemSize {
	if in == nil {
		return nil
	}
	out := new(FileSystemSize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemSpec) DeepCopyInto(out *FileSystemSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemSpec.
func (in *FileSystemSpec) DeepCopy() *FileSystemSpec {
	if in == nil {
		return nil
	}
	out := new(FileSystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemStatus) DeepCopyInto(out *FileSystemStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemStatus.
func (in *FileSystemStatus) DeepCopy() *FileSystemStatus {
	if in == nil {
		return nil
	}
	out := new(FileSystemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTarget) DeepCopyInto(out *MountTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTarget.
func (in *MountTarget) DeepCopy() *MountTarget {
	if in == nil {
		return nil
	}
	out := new(MountTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MountTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetDescription) DeepCopyInto(out *MountTargetDescription) {
	*out = *in
	if in.AvailabilityZoneID != nil {
		in, out := &in.AvailabilityZoneID, &out.AvailabilityZoneID
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZoneName != nil {
		in, out := &in.AvailabilityZoneName, &out.AvailabilityZoneName
		*out = new(string)
		**out = **in
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.LifeCycleState != nil {
		in, out := &in.LifeCycleState, &out.LifeCycleState
		*out = new(string)
		**out = **in
	}
	if in.MountTargetID != nil {
		in, out := &in.MountTargetID, &out.MountTargetID
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetDescription.
func (in *MountTargetDescription) DeepCopy() *MountTargetDescription {
	if in == nil {
		return nil
	}
	out := new(MountTargetDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetList) DeepCopyInto(out *MountTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MountTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetList.
func (in *MountTargetList) DeepCopy() *MountTargetList {
	if in == nil {
		return nil
	}
	out := new(MountTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MountTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetObservation) DeepCopyInto(out *MountTargetObservation) {
	*out = *in
	if in.AvailabilityZoneID != nil {
		in, out := &in.AvailabilityZoneID, &out.AvailabilityZoneID
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZoneName != nil {
		in, out := &in.AvailabilityZoneName, &out.AvailabilityZoneName
		*out = new(string)
		**out = **in
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.LifeCycleState != nil {
		in, out := &in.LifeCycleState, &out.LifeCycleState
		*out = new(string)
		**out = **in
	}
	if in.MountTargetID != nil {
		in, out := &in.MountTargetID, &out.MountTargetID
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetObservation.
func (in *MountTargetObservation) DeepCopy() *MountTargetObservation {
	if in == nil {
		return nil
	}
	out := new(MountTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetParameters) DeepCopyInto(out *MountTargetParameters) {
	*out = *in
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	in.CustomMountTargetParameters.DeepCopyInto(&out.CustomMountTargetParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetParameters.
func (in *MountTargetParameters) DeepCopy() *MountTargetParameters {
	if in == nil {
		return nil
	}
	out := new(MountTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetSpec) DeepCopyInto(out *MountTargetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetSpec.
func (in *MountTargetSpec) DeepCopy() *MountTargetSpec {
	if in == nil {
		return nil
	}
	out := new(MountTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetStatus) DeepCopyInto(out *MountTargetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetStatus.
func (in *MountTargetStatus) DeepCopy() *MountTargetStatus {
	if in == nil {
		return nil
	}
	out := new(MountTargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PosixUser) DeepCopyInto(out *PosixUser) {
	*out = *in
	if in.Gid != nil {
		in, out := &in.Gid, &out.Gid
		*out = new(int64)
		**out = **in
	}
	if in.SecondaryGids != nil {
		in, out := &in.SecondaryGids, &out.SecondaryGids
		*out = make([]*int64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(int64)
				**out = **in
			}
		}
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PosixUser.
func (in *PosixUser) DeepCopy() *PosixUser {
	if in == nil {
		return nil
	}
	out := new(PosixUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootDirectory) DeepCopyInto(out *RootDirectory) {
	*out = *in
	if in.CreationInfo != nil {
		in, out := &in.CreationInfo, &out.CreationInfo
		*out = new(CreationInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootDirectory.
func (in *RootDirectory) DeepCopy() *RootDirectory {
	if in == nil {
		return nil
	}
	out := new(RootDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
