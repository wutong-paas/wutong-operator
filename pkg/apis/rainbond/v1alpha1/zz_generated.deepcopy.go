// +build !ignore_autogenerated

// RAINBOND, Application Management Platform
// Copyright (C) 2014-2020 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunCloudDiskCSIPluginSource) DeepCopyInto(out *AliyunCloudDiskCSIPluginSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliyunCloudDiskCSIPluginSource.
func (in *AliyunCloudDiskCSIPluginSource) DeepCopy() *AliyunCloudDiskCSIPluginSource {
	if in == nil {
		return nil
	}
	out := new(AliyunCloudDiskCSIPluginSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunNasCSIPluginSource) DeepCopyInto(out *AliyunNasCSIPluginSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliyunNasCSIPluginSource.
func (in *AliyunNasCSIPluginSource) DeepCopy() *AliyunNasCSIPluginSource {
	if in == nil {
		return nil
	}
	out := new(AliyunNasCSIPluginSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailableNodes) DeepCopyInto(out *AvailableNodes) {
	*out = *in
	if in.SpecifiedNodes != nil {
		in, out := &in.SpecifiedNodes, &out.SpecifiedNodes
		*out = make([]*K8sNode, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(K8sNode)
				**out = **in
			}
		}
	}
	if in.MasterNodes != nil {
		in, out := &in.MasterNodes, &out.MasterNodes
		*out = make([]*K8sNode, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(K8sNode)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailableNodes.
func (in *AvailableNodes) DeepCopy() *AvailableNodes {
	if in == nil {
		return nil
	}
	out := new(AvailableNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIPluginSource) DeepCopyInto(out *CSIPluginSource) {
	*out = *in
	if in.AliyunCloudDisk != nil {
		in, out := &in.AliyunCloudDisk, &out.AliyunCloudDisk
		*out = new(AliyunCloudDiskCSIPluginSource)
		**out = **in
	}
	if in.AliyunNas != nil {
		in, out := &in.AliyunNas, &out.AliyunNas
		*out = new(AliyunNasCSIPluginSource)
		**out = **in
	}
	if in.NFS != nil {
		in, out := &in.NFS, &out.NFS
		*out = new(NFSCSIPluginSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIPluginSource.
func (in *CSIPluginSource) DeepCopy() *CSIPluginSource {
	if in == nil {
		return nil
	}
	out := new(CSIPluginSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdConfig) DeepCopyInto(out *EtcdConfig) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdConfig.
func (in *EtcdConfig) DeepCopy() *EtcdConfig {
	if in == nil {
		return nil
	}
	out := new(EtcdConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageHub) DeepCopyInto(out *ImageHub) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageHub.
func (in *ImageHub) DeepCopy() *ImageHub {
	if in == nil {
		return nil
	}
	out := new(ImageHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallPackageConfig) DeepCopyInto(out *InstallPackageConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallPackageConfig.
func (in *InstallPackageConfig) DeepCopy() *InstallPackageConfig {
	if in == nil {
		return nil
	}
	out := new(InstallPackageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sNode) DeepCopyInto(out *K8sNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sNode.
func (in *K8sNode) DeepCopy() *K8sNode {
	if in == nil {
		return nil
	}
	out := new(K8sNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NFSCSIPluginSource) DeepCopyInto(out *NFSCSIPluginSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NFSCSIPluginSource.
func (in *NFSCSIPluginSource) DeepCopy() *NFSCSIPluginSource {
	if in == nil {
		return nil
	}
	out := new(NFSCSIPluginSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageCondition) DeepCopyInto(out *PackageCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageCondition.
func (in *PackageCondition) DeepCopy() *PackageCondition {
	if in == nil {
		return nil
	}
	out := new(PackageCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondCluster) DeepCopyInto(out *RainbondCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(RainbondClusterStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondCluster.
func (in *RainbondCluster) DeepCopy() *RainbondCluster {
	if in == nil {
		return nil
	}
	out := new(RainbondCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondClusterCondition) DeepCopyInto(out *RainbondClusterCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondClusterCondition.
func (in *RainbondClusterCondition) DeepCopy() *RainbondClusterCondition {
	if in == nil {
		return nil
	}
	out := new(RainbondClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondClusterList) DeepCopyInto(out *RainbondClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RainbondCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondClusterList.
func (in *RainbondClusterList) DeepCopy() *RainbondClusterList {
	if in == nil {
		return nil
	}
	out := new(RainbondClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondClusterSpec) DeepCopyInto(out *RainbondClusterSpec) {
	*out = *in
	if in.GatewayIngressIPs != nil {
		in, out := &in.GatewayIngressIPs, &out.GatewayIngressIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodesForGateway != nil {
		in, out := &in.NodesForGateway, &out.NodesForGateway
		*out = make([]*K8sNode, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(K8sNode)
				**out = **in
			}
		}
	}
	if in.NodesForChaos != nil {
		in, out := &in.NodesForChaos, &out.NodesForChaos
		*out = make([]*K8sNode, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(K8sNode)
				**out = **in
			}
		}
	}
	if in.ImageHub != nil {
		in, out := &in.ImageHub, &out.ImageHub
		*out = new(ImageHub)
		**out = **in
	}
	if in.RegionDatabase != nil {
		in, out := &in.RegionDatabase, &out.RegionDatabase
		*out = new(Database)
		**out = **in
	}
	if in.UIDatabase != nil {
		in, out := &in.UIDatabase, &out.UIDatabase
		*out = new(Database)
		**out = **in
	}
	if in.EtcdConfig != nil {
		in, out := &in.EtcdConfig, &out.EtcdConfig
		*out = new(EtcdConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RainbondVolumeSpecRWX != nil {
		in, out := &in.RainbondVolumeSpecRWX, &out.RainbondVolumeSpecRWX
		*out = new(RainbondVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RainbondVolumeSpecRWO != nil {
		in, out := &in.RainbondVolumeSpecRWO, &out.RainbondVolumeSpecRWO
		*out = new(RainbondVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondClusterSpec.
func (in *RainbondClusterSpec) DeepCopy() *RainbondClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RainbondClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondClusterStatus) DeepCopyInto(out *RainbondClusterStatus) {
	*out = *in
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = make([]*StorageClass, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StorageClass)
				**out = **in
			}
		}
	}
	if in.GatewayAvailableNodes != nil {
		in, out := &in.GatewayAvailableNodes, &out.GatewayAvailableNodes
		*out = new(AvailableNodes)
		(*in).DeepCopyInto(*out)
	}
	if in.ChaosAvailableNodes != nil {
		in, out := &in.ChaosAvailableNodes, &out.ChaosAvailableNodes
		*out = new(AvailableNodes)
		(*in).DeepCopyInto(*out)
	}
	out.ImagePullSecret = in.ImagePullSecret
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RainbondClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondClusterStatus.
func (in *RainbondClusterStatus) DeepCopy() *RainbondClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RainbondClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondPackage) DeepCopyInto(out *RainbondPackage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(RainbondPackageStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondPackage.
func (in *RainbondPackage) DeepCopy() *RainbondPackage {
	if in == nil {
		return nil
	}
	out := new(RainbondPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondPackage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondPackageImage) DeepCopyInto(out *RainbondPackageImage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondPackageImage.
func (in *RainbondPackageImage) DeepCopy() *RainbondPackageImage {
	if in == nil {
		return nil
	}
	out := new(RainbondPackageImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondPackageList) DeepCopyInto(out *RainbondPackageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RainbondPackage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondPackageList.
func (in *RainbondPackageList) DeepCopy() *RainbondPackageList {
	if in == nil {
		return nil
	}
	out := new(RainbondPackageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondPackageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondPackageSpec) DeepCopyInto(out *RainbondPackageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondPackageSpec.
func (in *RainbondPackageSpec) DeepCopy() *RainbondPackageSpec {
	if in == nil {
		return nil
	}
	out := new(RainbondPackageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondPackageStatus) DeepCopyInto(out *RainbondPackageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PackageCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagesPushed != nil {
		in, out := &in.ImagesPushed, &out.ImagesPushed
		*out = make([]RainbondPackageImage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondPackageStatus.
func (in *RainbondPackageStatus) DeepCopy() *RainbondPackageStatus {
	if in == nil {
		return nil
	}
	out := new(RainbondPackageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondVolume) DeepCopyInto(out *RainbondVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondVolume.
func (in *RainbondVolume) DeepCopy() *RainbondVolume {
	if in == nil {
		return nil
	}
	out := new(RainbondVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondVolumeCondition) DeepCopyInto(out *RainbondVolumeCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondVolumeCondition.
func (in *RainbondVolumeCondition) DeepCopy() *RainbondVolumeCondition {
	if in == nil {
		return nil
	}
	out := new(RainbondVolumeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondVolumeList) DeepCopyInto(out *RainbondVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RainbondVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondVolumeList.
func (in *RainbondVolumeList) DeepCopy() *RainbondVolumeList {
	if in == nil {
		return nil
	}
	out := new(RainbondVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RainbondVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondVolumeSpec) DeepCopyInto(out *RainbondVolumeSpec) {
	*out = *in
	if in.StorageClassParameters != nil {
		in, out := &in.StorageClassParameters, &out.StorageClassParameters
		*out = new(StorageClassParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.CSIPlugin != nil {
		in, out := &in.CSIPlugin, &out.CSIPlugin
		*out = new(CSIPluginSource)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageRequest != nil {
		in, out := &in.StorageRequest, &out.StorageRequest
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondVolumeSpec.
func (in *RainbondVolumeSpec) DeepCopy() *RainbondVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(RainbondVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RainbondVolumeStatus) DeepCopyInto(out *RainbondVolumeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RainbondVolumeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RainbondVolumeStatus.
func (in *RainbondVolumeStatus) DeepCopy() *RainbondVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(RainbondVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdComponent) DeepCopyInto(out *RbdComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(RbdComponentStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdComponent.
func (in *RbdComponent) DeepCopy() *RbdComponent {
	if in == nil {
		return nil
	}
	out := new(RbdComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RbdComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdComponentCondition) DeepCopyInto(out *RbdComponentCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdComponentCondition.
func (in *RbdComponentCondition) DeepCopy() *RbdComponentCondition {
	if in == nil {
		return nil
	}
	out := new(RbdComponentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdComponentList) DeepCopyInto(out *RbdComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RbdComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdComponentList.
func (in *RbdComponentList) DeepCopy() *RbdComponentList {
	if in == nil {
		return nil
	}
	out := new(RbdComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RbdComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdComponentSpec) DeepCopyInto(out *RbdComponentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdComponentSpec.
func (in *RbdComponentSpec) DeepCopy() *RbdComponentSpec {
	if in == nil {
		return nil
	}
	out := new(RbdComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbdComponentStatus) DeepCopyInto(out *RbdComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RbdComponentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbdComponentStatus.
func (in *RbdComponentStatus) DeepCopy() *RbdComponentStatus {
	if in == nil {
		return nil
	}
	out := new(RbdComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClass) DeepCopyInto(out *StorageClass) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClass.
func (in *StorageClass) DeepCopy() *StorageClass {
	if in == nil {
		return nil
	}
	out := new(StorageClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClassParameters) DeepCopyInto(out *StorageClassParameters) {
	*out = *in
	if in.MountOptions != nil {
		in, out := &in.MountOptions, &out.MountOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClassParameters.
func (in *StorageClassParameters) DeepCopy() *StorageClassParameters {
	if in == nil {
		return nil
	}
	out := new(StorageClassParameters)
	in.DeepCopyInto(out)
	return out
}
