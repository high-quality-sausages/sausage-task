package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"

	"github.com/high-quality-sausages/sausage-task/gateway/internal/grpc"
	taskApi "github.com/high-quality-sausages/sausage-task/task/api"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetTasksByUid(ctx *gin.Context) {
	log.Println("uid:", ctx.Query("uid"))
	uid, err := strconv.ParseInt(ctx.Query("uid"), 10, 64)
	if err != nil {
		log.Printf("[GetTasksByUid] failed to parse uid, err:%v", err)
		ctx.JSON(400, gin.H{
			"message": "known parameter",
		})
		return
	}
	req := &taskApi.GetTasksByUidReq{
		Uid: uid,
	}

	taskClient := grpc.GetTaskClient()
	r, err := taskClient.GetTasksByUid(context.Background(), req)
	if err != nil {
		log.Println("[GetTasksByUid|grpc] failed to request grpc, err:%v", err)
		ctx.JSON(500, gin.H{
			"message": "Internal server error",
		})
	}

	tasks := make([]*taskApi.TaskInfo, 0, len(r.Tasks))
	if len(r.Tasks) != 0 {
		tasks = r.Tasks
	}
	log.Printf("tasks:%v\n", r.Tasks)
	ctx.JSON(200, gin.H{
		"tasks": tasks,
	})
}
