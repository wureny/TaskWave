package rpc

import "github.com/wureny/TaskWave/taskutils/rpc/model"

type TaskRpc struct {
	Host string
}

func (t TaskRpc) CreateTask(req *model.CreateTaskReq) (*model.CreateTaskResp, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRpc) SetTask(req *model.SetTaskReq) (*model.SetTaskResp, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRpc) HoldTasks(req *model.HoldTasksReq) (*model.HoldTasksResp, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRpc) GetTaskList(req *model.GetTaskListReq) (*model.GetTaskListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRpc) GetTask(req *model.GetTaskReq) (*model.GetTaskResp, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRpc) GetTaskScheduleCfgList() (*model.GetTaskScheduleCfgListResp, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRpc) GetTaskCountsByType(req *model.GetTaskCountsByTypeReq) (*model.GetTaskCountsByTypeResp, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskRpc) GetTaskCountsBtStatus(req *model.GetTaskCountsBtStatusReq) (*model.GetTaskCountsBtStatusResp, error) {
	//TODO implement me
	panic("implement me")
}
