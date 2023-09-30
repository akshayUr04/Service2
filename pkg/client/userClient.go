package client

import (
	"context"
	"fmt"
	"x-tentioncrew/microservice-2/pkg/config"
	"x-tentioncrew/microservice-2/pkg/pb/userproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	Client userproto.UserSvcClient
}

func InitUserSvc(cfg *config.Config) UserServiceClient {
	cc, err := grpc.Dial(cfg.UserServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	c := UserServiceClient{
		Client: userproto.NewUserSvcClient(cc),
	}

	return c
}

func (c *UserServiceClient) GetUserData() (*userproto.GetAllUserDataResult, error) {
	fmt.Println("in getallmethod")
	val, err := c.Client.GetUserData(context.Background(), &userproto.GetAllUserDataReq{})
	return val, err
}
