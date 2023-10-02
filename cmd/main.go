package main

import (
	"log"
	"net"
	"x-tentioncrew/microservice-2/pkg/client"
	"x-tentioncrew/microservice-2/pkg/config"
	"x-tentioncrew/microservice-2/pkg/pb"
	"x-tentioncrew/microservice-2/pkg/service"

	"google.golang.org/grpc"
)

func main() {
	cfg, cfgErr := config.LoadConfig()

	if cfgErr != nil {
		log.Fatalln("Failed at config", cfgErr)
	}
	lis, lisErr := net.Listen("tcp", cfg.M2PORT)
	if lisErr != nil {
		log.Fatalln("Failed to listing:", lisErr)
	}

	userSvc := client.InitUserSvc(&cfg)

	s := service.Server{
		UserSvc: userSvc,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterService2Server(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
