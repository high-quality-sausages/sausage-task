package grpc

import (
	taskApi "github.com/high-quality-sausages/sausage-task/task/api"
	"google.golang.org/grpc"
	"log"
)

var (
	TaskClient taskApi.TaskClient
)

func Init() {
	conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to listen connect grpc, err:%v\n", err)
	}
	TaskClient = taskApi.NewTaskClient(conn)
}

func GetTaskClient() taskApi.TaskClient {
	return TaskClient
}
