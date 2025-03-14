package controllers

import (
	"context"
	"errors"

	wutongv1alpha1 "github.com/wutong-paas/wutong-operator/v2/api/v1alpha1"
	"github.com/wutong-paas/wutong-operator/v2/controllers/plugin"
	"github.com/wutong-paas/wutong-operator/v2/controllers/plugin/aliyunclouddisk"
	"github.com/wutong-paas/wutong-operator/v2/controllers/plugin/aliyunnas"
	"github.com/wutong-paas/wutong-operator/v2/controllers/plugin/nfs"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewCSIPlugin creates a new csi plugin
func NewCSIPlugin(ctx context.Context, cli client.Client, volume *wutongv1alpha1.WutongVolume) (plugin.CSIPlugin, error) {
	cp := volume.Spec.CSIPlugin
	var p plugin.CSIPlugin
	switch {
	case cp.AliyunCloudDisk != nil:
		p = aliyunclouddisk.CSIPlugins(ctx, cli, volume)
	case cp.AliyunNas != nil:
		p = aliyunnas.CSIPlugins(ctx, cli, volume)
	case cp.NFS != nil:
		p = nfs.CSIPlugins(ctx, cli, volume)
	}
	if p == nil {
		return nil, errors.New("unsupported csi plugin")
	}
	return p, nil
}
