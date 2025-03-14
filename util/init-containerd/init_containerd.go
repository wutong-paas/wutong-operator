package init_containerd

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/namespaces"
	"github.com/wutong-paas/wutong-operator/v2/util/k8sutil"
)

const ()

// ContainerdAPI -
type ContainerdAPI struct {
	ImageService     images.Store
	CCtx             context.Context
	ContainerdClient *containerd.Client
}

func InitContainerd() (*ContainerdAPI, error) {
	cr := k8sutil.GetContainerRuntime()
	containerdClient, err := containerd.New(cr.Endpoint)
	if err != nil {
		return nil, err
	}

	cctx := namespaces.WithNamespace(context.Background(), "k8s.io")
	imageService := containerdClient.ImageService()
	return &ContainerdAPI{
		ImageService:     imageService,
		CCtx:             cctx,
		ContainerdClient: containerdClient,
	}, nil
}
