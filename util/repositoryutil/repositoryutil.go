package repositoryutil

import (
	"context"
	"fmt"
	"net/url"
	"runtime"
	"strings"

	dockercliconfigtypes "github.com/docker/cli/cli/config/types"
	registrytypes "github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/registry"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Code from https://github.com/containerd/nerdctl/blob/v0.15.0/cmd/nerdctl/login.go
func loginClientSide(ctx context.Context, insecure bool, auth registrytypes.AuthConfig) (registrytypes.AuthenticateOKBody, error) {
	var insecureRegistries []string
	if insecure {
		insecureRegistries = append(insecureRegistries, auth.ServerAddress)
	}
	svc, err := registry.NewService(registry.ServiceOptions{
		InsecureRegistries: insecureRegistries,
	})

	if err != nil {
		return registrytypes.AuthenticateOKBody{}, err
	}

	userAgent := fmt.Sprintf("Docker-Client/nerdctl-%s", runtime.GOOS)

	status, token, err := svc.Auth(ctx, &auth, userAgent)

	return registrytypes.AuthenticateOKBody{
		Status:        status,
		IdentityToken: token,
	}, err
}

func convertToHostname(serverAddress string) (string, error) {
	// Ensure that URL contains scheme for a good parsing process
	if strings.Contains(serverAddress, "://") {
		u, err := url.Parse(serverAddress)
		if err != nil {
			return "", err
		}
		serverAddress = u.Host
	} else {
		u, err := url.Parse("https://" + serverAddress)
		if err != nil {
			return "", err
		}
		serverAddress = u.Host
	}

	return serverAddress, nil
}

// GetDefaultAuthConfig gets default auth config
func GetDefaultAuthConfig(serverAddress, username, password string, isDefaultRegistry bool) (*registrytypes.AuthConfig, error) {
	if !isDefaultRegistry {
		var err error
		serverAddress, err = convertToHostname(serverAddress)
		if err != nil {
			return nil, err
		}
	}
	res := registrytypes.AuthConfig(dockercliconfigtypes.AuthConfig{
		ServerAddress: serverAddress,
	})
	if username != "" {
		res.Username = username
	}
	if password != "" {
		res.Password = password
	}
	return &res, nil
}

// ConfigureAuthentication configures authentication for a registry
func ConfigureAuthentication(authConfig *registrytypes.AuthConfig, username, password string) error {
	authConfig.Username = strings.TrimSpace(authConfig.Username)
	if username = strings.TrimSpace(username); username == "" {
		username = authConfig.Username
	}
	if username == "" {
		return fmt.Errorf("error: Username is Required")
	}
	if password == "" {
		return fmt.Errorf("error: Password is Required")
	}
	authConfig.Username = username
	authConfig.Password = password
	return nil
}

// LoginRepository logs in to a image repository
// 离线环境下，wutong.me 作为 serverAddress，此函数调用不通过，会导致创建 wutongcomponent 组件失败
func LoginRepository(serverAddress, username, password string) error {
	ctx := context.Background()
	isDefaultRegistry := serverAddress == "wutong.me"

	authConfig, err := GetDefaultAuthConfig(serverAddress, username, password, isDefaultRegistry)
	if authConfig == nil {
		authConfig = &registrytypes.AuthConfig{ServerAddress: serverAddress}
	}

	if err == nil && authConfig.Username != "" && authConfig.Password != "" {
		//login With StoreCreds
		_, err = loginClientSide(ctx, false, *authConfig)
	}

	if err != nil || authConfig.Username == "" || authConfig.Password == "" {
		logrus.Infof("First login failed [%+v], login insecure repository %s with username %s and password %s", err, serverAddress, authConfig.Username, authConfig.Password)
		err = ConfigureAuthentication(authConfig, username, password)
		if err != nil {
			return errors.Wrap(err, "ConfigureAuthentication")
		}
		_, err = loginClientSide(ctx, true, *authConfig)
		if err != nil {
			return errors.Wrap(err, "loginClientSide")
		}
	}
	return nil
}
