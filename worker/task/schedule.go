package task

import (
	"github.com/wureny/TaskWave/taskutils/rpc"
	"github.com/wureny/TaskWave/taskutils/rpc/model"
)

var taskRpc rpc.TaskRpc

var (
	taskSvrHost string // new is host: for example http://127.0.0.1:41555
)

var scheduleCfgDict map[string]*model.TaskScheduleCfg

func init() {
	scheduleCfgDict = make(map[string]*model.TaskScheduleCfg)
}

// LoadCfg func load cfg
func LoadCfg() error {
	cfgList, err := taskRpc.GetTaskScheduleCfgList()
	if err != nil {
		//TODO log
		return err
	}
	for _, cfg := range cfgList.ScheduleCfgList {
		scheduleCfgDict[cfg.TaskType] = cfg
	}
	return nil
}
