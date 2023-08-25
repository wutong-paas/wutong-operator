package init_containerd

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/namespaces"
	"github.com/pelletier/go-toml"
)

// ContainerdAPI -
type ContainerdAPI struct {
	ImageService     images.Store
	CCtx             context.Context
	ContainerdClient *containerd.Client
}

func InitContainerd() (*ContainerdAPI, error) {
	// 通过读取配置文件，获取 containerd 的 socket 地址
	containerdConf, err := toml.LoadFile("/etc/containerd/config.toml")
	if err != nil {
		return nil, err
	}
	address := containerdConf.Get("grpc.address").(string)
	if address == "" {
		return nil, err
	}

	containerdClient, err := containerd.New(address)
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
