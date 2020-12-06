package model

import "time"

type Task struct {
	TaskName  string    `json:"task_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Desc      string    `json:"desc"`
}
