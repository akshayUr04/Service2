package service

import (
	"context"
	"fmt"
	"time"
	"x-tentioncrew/microservice-2/pkg/client"
	"x-tentioncrew/microservice-2/pkg/pb"
)

type Server struct {
	UserSvc client.UserServiceClient
	pb.UnimplementedService2Server
}

func (s *Server) Methods(ctx context.Context, req *pb.MethodRequest) (*pb.MethodResponce, error) {
	if req.Method == 1 {
		fmt.Println("in method")
		res, err := s.UserSvc.GetUserData()
		if err != nil {
			return nil, err
		}
		names := res.Name
		time.Sleep(time.Duration(req.WaitTime) * time.Second)
		return &pb.MethodResponce{UserName: names}, nil
	} else {
		time.Sleep(time.Duration(req.WaitTime) * time.Second)
		res, err := s.UserSvc.GetUserData()
		if err != nil {
			return nil, err
		}
		names := res.Name
		return &pb.MethodResponce{UserName: names}, nil
	}
}
