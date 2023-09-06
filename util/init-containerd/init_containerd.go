package init_containerd

import (
	"context"
	"os"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/namespaces"
	"github.com/pelletier/go-toml"
)

const (
	ContainerdConfigPath     = "/etc/containerd/config.toml"
	DefaultContainerdAddress = "/run/containerd/containerd.sock"
	K3sContainerdAddress     = "/run/k3s/containerd/containerd.sock"
)

var (
	ContainerdSockAddressList = []string{
		DefaultContainerdAddress,
		K3sContainerdAddress,
	}
)

// ContainerdAPI -
type ContainerdAPI struct {
	ImageService     images.Store
	CCtx             context.Context
	ContainerdClient *containerd.Client
}

func InitContainerd() (*ContainerdAPI, error) {
	containerdClient, err := containerd.New(GetRuntimeSocketAddress())
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

func GetRuntimeSocketAddress() string {
	var result string

	containerdConf, err := toml.LoadFile(ContainerdConfigPath)
	if err == nil {
		result = containerdConf.Get("grpc.address").(string)
	}

	if result == "" {
		for _, sock := range ContainerdSockAddressList {
			if _, err := os.Stat(sock); err == nil {
				result = sock
				break
			}
		}
	}

	if result == "" {
		result = DefaultContainerdAddress
	}
	return result
}
