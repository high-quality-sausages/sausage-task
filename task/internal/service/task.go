package service

import (
	"context"
	"log"

	"github.com/high-quality-sausages/sausage-task/task/api"
)

// GetTasksByUid(ctx context.Context, in *GetTasksByUidReq, opts ...grpc.CallOption) (*GetTasksByUidResp, error)

func (s *Service) GetTasksByUid(ctx context.Context, req *api.GetTasksByUidReq) (resp *api.GetTasksByUidResp, err error) {
	resp = new(api.GetTasksByUidResp)
	tasks, err := s.dao.GetUserTasksByUid(context.Background(), req.Uid)
	if err != nil {
		log.Printf("[GetTasksByUid] failed to get tasks by uid, uid:%d, err:%v", req.Uid, err)
		return
	}

	for _, task := range tasks {
		resp.Tasks = append(resp.Tasks, &api.TaskInfo{
			TaskName: task.TaskName,
			Desc:     task.Desc,
			// StartTime: task.StartTime,
		})
	}

	return
}
