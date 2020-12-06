package server

import (
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/high-quality-sausages/sausage-task/task/api"
	"github.com/high-quality-sausages/sausage-task/task/internal/service"
)

func Init(svc *service.Service) {
	listen, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	log.Println("start listening port: 9000")
	s := grpc.NewServer()
	api.RegisterTaskServer(s, svc)

	if err := s.Serve(listen); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
