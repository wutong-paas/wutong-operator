// +build !ignore_autogenerated

// RAINBOND, Application Management Platform
// Copyright (C) 2020-2021 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Wutong,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliyunCloudDiskCSIPluginSource) DeepCopyInto(out *AliyunCloudDiskCSIPluginSource) {
	*out = *in
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
func (in *WutongCluster) DeepCopyInto(out *WutongCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongCluster.
func (in *WutongCluster) DeepCopy() *WutongCluster {
	if in == nil {
		return nil
	}
	out := new(WutongCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WutongCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongClusterCondition) DeepCopyInto(out *WutongClusterCondition) {
	*out = *in
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongClusterCondition.
func (in *WutongClusterCondition) DeepCopy() *WutongClusterCondition {
	if in == nil {
		return nil
	}
	out := new(WutongClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongClusterList) DeepCopyInto(out *WutongClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WutongCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongClusterList.
func (in *WutongClusterList) DeepCopy() *WutongClusterList {
	if in == nil {
		return nil
	}
	out := new(WutongClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WutongClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongClusterSpec) DeepCopyInto(out *WutongClusterSpec) {
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
	if in.WutongVolumeSpecRWX != nil {
		in, out := &in.WutongVolumeSpecRWX, &out.WutongVolumeSpecRWX
		*out = new(WutongVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.WutongVolumeSpecRWO != nil {
		in, out := &in.WutongVolumeSpecRWO, &out.WutongVolumeSpecRWO
		*out = new(WutongVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongClusterSpec.
func (in *WutongClusterSpec) DeepCopy() *WutongClusterSpec {
	if in == nil {
		return nil
	}
	out := new(WutongClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongClusterStatus) DeepCopyInto(out *WutongClusterStatus) {
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
	if in.ImagePullSecret != nil {
		in, out := &in.ImagePullSecret, &out.ImagePullSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]WutongClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongClusterStatus.
func (in *WutongClusterStatus) DeepCopy() *WutongClusterStatus {
	if in == nil {
		return nil
	}
	out := new(WutongClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongPackage) DeepCopyInto(out *WutongPackage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongPackage.
func (in *WutongPackage) DeepCopy() *WutongPackage {
	if in == nil {
		return nil
	}
	out := new(WutongPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WutongPackage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongPackageImage) DeepCopyInto(out *WutongPackageImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongPackageImage.
func (in *WutongPackageImage) DeepCopy() *WutongPackageImage {
	if in == nil {
		return nil
	}
	out := new(WutongPackageImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongPackageList) DeepCopyInto(out *WutongPackageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WutongPackage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongPackageList.
func (in *WutongPackageList) DeepCopy() *WutongPackageList {
	if in == nil {
		return nil
	}
	out := new(WutongPackageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WutongPackageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongPackageSpec) DeepCopyInto(out *WutongPackageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongPackageSpec.
func (in *WutongPackageSpec) DeepCopy() *WutongPackageSpec {
	if in == nil {
		return nil
	}
	out := new(WutongPackageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongPackageStatus) DeepCopyInto(out *WutongPackageStatus) {
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
		*out = make([]WutongPackageImage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongPackageStatus.
func (in *WutongPackageStatus) DeepCopy() *WutongPackageStatus {
	if in == nil {
		return nil
	}
	out := new(WutongPackageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongVolume) DeepCopyInto(out *WutongVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongVolume.
func (in *WutongVolume) DeepCopy() *WutongVolume {
	if in == nil {
		return nil
	}
	out := new(WutongVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WutongVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongVolumeCondition) DeepCopyInto(out *WutongVolumeCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongVolumeCondition.
func (in *WutongVolumeCondition) DeepCopy() *WutongVolumeCondition {
	if in == nil {
		return nil
	}
	out := new(WutongVolumeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongVolumeList) DeepCopyInto(out *WutongVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WutongVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongVolumeList.
func (in *WutongVolumeList) DeepCopy() *WutongVolumeList {
	if in == nil {
		return nil
	}
	out := new(WutongVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WutongVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongVolumeSpec) DeepCopyInto(out *WutongVolumeSpec) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongVolumeSpec.
func (in *WutongVolumeSpec) DeepCopy() *WutongVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(WutongVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongVolumeStatus) DeepCopyInto(out *WutongVolumeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]WutongVolumeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongVolumeStatus.
func (in *WutongVolumeStatus) DeepCopy() *WutongVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(WutongVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongComponent) DeepCopyInto(out *WutongComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongComponent.
func (in *WutongComponent) DeepCopy() *WutongComponent {
	if in == nil {
		return nil
	}
	out := new(WutongComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WutongComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongComponentCondition) DeepCopyInto(out *WutongComponentCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongComponentCondition.
func (in *WutongComponentCondition) DeepCopy() *WutongComponentCondition {
	if in == nil {
		return nil
	}
	out := new(WutongComponentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongComponentList) DeepCopyInto(out *WutongComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WutongComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongComponentList.
func (in *WutongComponentList) DeepCopy() *WutongComponentList {
	if in == nil {
		return nil
	}
	out := new(WutongComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WutongComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongComponentSpec) DeepCopyInto(out *WutongComponentSpec) {
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
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongComponentSpec.
func (in *WutongComponentSpec) DeepCopy() *WutongComponentSpec {
	if in == nil {
		return nil
	}
	out := new(WutongComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WutongComponentStatus) DeepCopyInto(out *WutongComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]WutongComponentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WutongComponentStatus.
func (in *WutongComponentStatus) DeepCopy() *WutongComponentStatus {
	if in == nil {
		return nil
	}
	out := new(WutongComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClass) DeepCopyInto(out *StorageClass) {
	*out = *in
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
