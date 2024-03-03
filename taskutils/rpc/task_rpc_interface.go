package rpc

import "github.com/wureny/TaskWave/taskutils/rpc/model"

type ITaskRpc interface {
	CreateTask(req *model.CreateTaskReq) (*model.CreateTaskResp, error)
	SetTask(req *model.SetTaskReq) (*model.SetTaskResp, error)
	HoldTasks(req *model.HoldTasksReq) (*model.HoldTasksResp, error)
	GetTaskList(req *model.GetTaskListReq) (*model.GetTaskListResp, error)
	GetTask(req *model.GetTaskReq) (*model.GetTaskResp, error)
	GetTaskScheduleCfgList() (*model.GetTaskScheduleCfgListResp, error)
	GetTaskCountsByType(req *model.GetTaskCountsByTypeReq) (*model.GetTaskCountsByTypeResp, error)
	GetTaskCountsBtStatus(req *model.GetTaskCountsBtStatusReq) (*model.GetTaskCountsBtStatusResp, error)
}
