package pkg

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func RegistryConsul() registry.Registry {
	return consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
}
