package main

import (
	"github.com/high-quality-sausages/sausage-task/gateway/api"
	"github.com/high-quality-sausages/sausage-task/gateway/internal/grpc"
)

func main() {
	grpc.Init()
	api.Init()
	r := api.GetHandler()
	r.Run()
}
