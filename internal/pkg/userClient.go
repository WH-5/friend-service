package pkg

import (
	"context"
	v2 "github.com/WH-5/friend-service/api/push/v1"
	v1 "github.com/WH-5/friend-service/api/user/v1"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
)

func UserClient(consuls string) v1.UserClient {

	cg := api.DefaultConfig()
	cg.Address = consuls

	client, err := api.NewClient(cg)
	if err != nil {

		return nil
	}

	reg := consul.New(client)
	endpoint := "discovery:///user-service"

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(reg),
	)
	if err != nil {
		return nil
	}
	// Log connection details and status
	userClient := v1.NewUserClient(conn)
	return userClient
}

func PushClient(consuls string) v2.PushClient {

	cg := api.DefaultConfig()
	cg.Address = consuls

	client, err := api.NewClient(cg)
	if err != nil {
		return nil
	}

	reg := consul.New(client)
	endpoint := "discovery:///push-service"

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(reg),
	)
	if err != nil {
		return nil
	}

	pushClient := v2.NewPushClient(conn)
	return pushClient
}
