package imageutil

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/distribution/reference"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
)

// CheckIfImageExists -
func CheckIfImageExists(pctx context.Context, dockerClient *client.Client, imageName string) (bool, error) {
	repo, err := reference.Parse(imageName)
	if err != nil {
		return false, fmt.Errorf("parse image %s: %v", imageName, err)
	}
	named := repo.(reference.Named)
	tag := "latest"
	if t, ok := repo.(reference.Tagged); ok {
		tag = t.Tag()
	}
	imageFullName := named.Name() + ":" + tag

	ctx, cancel := context.WithCancel(pctx)
	defer cancel()

	imageSummarys, err := dockerClient.ImageList(ctx, image.ListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: "reference", Value: imageFullName}),
	})
	if err != nil {
		return false, fmt.Errorf("list images: %v", err)
	}

	return len(imageSummarys) > 0, nil
}

// ImagePull -
func ImagePull(ctx context.Context, dockerClient *client.Client, imageName string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rf, err := reference.ParseAnyReference(imageName)
	if err != nil {
		return err
	}

	res, err := dockerClient.ImagePull(ctx, rf.String(), image.PullOptions{})
	if err != nil {
		return fmt.Errorf("pull image %s failure %s", imageName, err.Error())
	}
	if res != nil {
		defer res.Close()
		dec := json.NewDecoder(res)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			var jm jsonmessage.JSONMessage
			if err := dec.Decode(&jm); err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("failed to decode json message: %v", err)
			}
			if jm.Error != nil {
				return fmt.Errorf("error detail: %v", jm.Error)
			}
		}
	}
	return nil
}

// ImagePush -
func ImagePush(ctx context.Context, dockerClient *client.Client, imageName, repo, user, pass string) error {
	var opts image.PushOptions
	authConfig := registry.AuthConfig{
		ServerAddress: repo,
	}
	authConfig.Username = user
	authConfig.Password = pass

	registryAuth, err := encodeAuthToBase64(authConfig)
	if err != nil {
		return fmt.Errorf("failed to encode auth config: %v", err)
	}
	opts.RegistryAuth = registryAuth
	var res io.ReadCloser
	res, err = dockerClient.ImagePush(ctx, imageName, opts)
	if err != nil {
		return err
	}
	if res != nil {
		defer res.Close()

		dec := json.NewDecoder(res)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			var jm jsonmessage.JSONMessage
			if err := dec.Decode(&jm); err != nil {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("failed to decode json message: %v", err)
			}
			if jm.Error != nil {
				return fmt.Errorf("error detail: %v", jm.Error)
			}
		}
	}
	return nil
}

func encodeAuthToBase64(authConfig registry.AuthConfig) (string, error) {
	buf, err := json.Marshal(authConfig)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buf), nil
}
