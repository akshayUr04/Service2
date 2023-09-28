package client

import (
	"context"
	"fmt"
	"x-tentioncrew/microservice-2/pkg/config"
	"x-tentioncrew/microservice-2/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func InitUserSvc(cfg *config.Config) UserServiceClient {
	cc, err := grpc.Dial(cfg.UserServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	c := UserServiceClient{
		Client: pb.NewUserServiceClient(cc),
	}

	return c
}

func (c *UserServiceClient) Methods() (*pb.GetAllUserDataResult, error) {
	req := *&pb.GetAllUserDataReq{}
	return c.Client.GetAllUserData(context.Background(), &req)
}
