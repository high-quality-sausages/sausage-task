package api

import "github.com/gin-gonic/gin"

var (
	handler *gin.Engine
)

func Init() {
	handler = gin.Default()
	handler.GET("ping", Ping)
	handler.GET("get_tasks_by_uid", GetTasksByUid)
}

func GetHandler() *gin.Engine {
	return handler
}
