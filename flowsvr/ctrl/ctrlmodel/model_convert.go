package ctrlmodel

import (
	"github.com/wureny/TaskWave/flowsvr/db"
	"github.com/wureny/TaskWave/taskutils/rpc/model"
)

// model.TaskData to db.Task
func FillTaskModel(data *model.TaskData) *db.Task {
	task := &db.Task{}
	task.TaskId = data.TaskId
	task.TaskType = data.TaskType
	task.UserId = data.UserId
	task.Status = data.Status
	task.ScheduleLog = data.ScheduleLog
	task.TaskStage = data.TaskStage
	task.CrtRetryNum = data.CrtRetryNum
	task.MaxRetryNum = data.MaxRetryNum
	task.MaxRetryInterval = data.MaxRetryInterval
	task.TaskContext = data.TaskContext
	return task
}

// db.Task to model.TaskData
func FillTaskResp(task *db.Task) *model.TaskData {
	data := &model.TaskData{}
	data.UserId = task.UserId
	data.TaskId = task.TaskId
	data.TaskType = task.TaskType
	data.Status = task.Status
	data.ScheduleLog = task.ScheduleLog
	priority := task.Priority
	data.Priority = &priority
	data.TaskStage = task.TaskStage
	data.CrtRetryNum = task.CrtRetryNum
	data.MaxRetryNum = task.MaxRetryNum
	data.MaxRetryInterval = task.MaxRetryInterval
	data.TaskContext = task.TaskContext
	data.CreateTime = task.CreateTime
	data.ModifyTime = task.ModifyTime
	return data
}
