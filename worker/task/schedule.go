package task

import "github.com/wureny/TaskWave/taskutils/rpc/model"

var scheduleCfgDict map[string]*model.TaskScheduleCfg

func init() {
	scheduleCfgDict = make(map[string]*model.TaskScheduleCfg)
}
