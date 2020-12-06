package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/high-quality-sausages/sausage-task/task/internal/model"
)

const (
	_insertTask = "INSERT INTO tbl_user_task (uid,task_name,start_time,end_time,description) VALUES(?,?,?,?,?);"
	_selectTask = "SELECT task_name,start_time,end_time FROM tbl_user_task WHERE uid=?;"
	_updateTask = "UPDATE tbl_user_task SET task_name=?,start_time=?,end_time=? WHERE uid=? AND task_id=?;"
	_deleteTask = "DELETE FROMtbl_user_task WHERE uid=? AND task_id=?;"
)

// CreateUserTask 创建用户任务
func (d *MainDao) CreateUserTask(ctx context.Context, uid int64, task model.Task) (err error) {
	if _, err = d.DB.Exec(_insertTask, uid, task.TaskName, task.StartTime, task.EndTime); err != nil {
		fmt.Printf("failed to exec sql, err:%v\n", err)
	}
	return
}

// GetUserTask
func (d *MainDao) GetUserTasksByUid(ctx context.Context, uid int64) (resp []model.Task, err error) {
	rows, err := d.DB.Query(_selectTask, uid)
	if err != nil {
		fmt.Printf("[dao|GetUserTasksByUid] failed to get user tasks, uid:%d, err:%v", uid, err)
		return
	}
	defer rows.Close()

	resp = make([]model.Task, 0)

	for rows.Next() {
		task := model.Task{}
		var start_time string
		var end_time string
		if err = rows.Scan(&task.TaskName, &start_time, &end_time); err != nil {
			log.Println("[dao|GetUserTasksByUid] failed to scan rows, err:%v", err)
			return
		}
		resp = append(resp, task)
	}
	return
}

func (d *MainDao) UpdateUserTask(
	ctx context.Context, uid, taskID int64, task model.Task) (err error) {

	return
}

func (d *MainDao) DeleteUserTask(ctx context.Context, uid, taskID int64) (err error) {
	return
}
