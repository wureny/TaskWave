package model

import "time"

type RespBase struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CreateTaskResp struct {
	RespBase
	TaskId string `json:"task_id"`
}

type CreateTaskReq struct {
	TaskData TaskData `json:"task_data"`
}

type GetTaskListReq struct {
	TaskType string `json:"task_type"`
	Status   int    `json:"status"`
	Limit    int    `json:"limit"`
}

type GetTaskListResp struct {
	RespBase
	TaskList []*TaskData `json:"task_list"`
}

type HoldTasksReq struct {
	TaskType string `json:"task_type"`
	Limit    int    `json:"limit"`
}

type HoldTasksResp struct {
	RespBase
	TaskList []*TaskData `json:"task_list"`
}

type GetTaskCountsByTypeReq struct {
	TaskType string `json:"task_type"`
}

type GetTaskCountsByTypeResp struct {
	RespBase
	Count int `json:"count"`
}

type GetTaskReq struct {
	TaskId string `json:"task_id"`
}

type GetTaskResp struct {
	RespBase
	TaskData *TaskData `json:"task_data"`
}

type GetTaskCountsBtStatusReq struct {
	Status int `json:"status"`
}

type GetTaskCountsBtStatusResp struct {
	RespBase
	Count int `json:"count"`
}

// GetTaskScheduleCfgListReq 获取任务配置信息 请求体（空）
type GetTaskScheduleCfgListReq struct {
}

// GetTaskScheduleCfgListResp 获取任务配置信息 响应体
type GetTaskScheduleCfgListResp struct {
	RespBase
	ScheduleCfgList []*TaskScheduleCfg `json:"task_schedule_cfg_list"`
}

type RegisterTaskReq struct {
	TaskType string `json:"task_type"`
	//TODO: Host is a url, so it should be a url type
	//Host     string `json:"host"`
}

type RegisterTaskResp struct {
	RespBase
}

// TaskScheduleCfg 任务调度信息
type TaskScheduleCfg struct {
	TaskType          string     `json:"task_type"`
	ScheduleLimit     int        `json:"schedule_limit"`
	ScheduleInterval  int        `json:"schedule_interval"`
	MaxProcessingTime int64      `json:"max_processing_time"`
	MaxRetryNum       int        `json:"max_retry_num"`
	RetryInterval     int        `json:"retry_interval"`
	MaxRetryInterval  int        `json:"max_retry_interval"`
	CreateTime        *time.Time `json:"create_time"`
	ModifyTime        *time.Time `json:"modify_time"`
}

// SetTaskReq 请求消息
type SetTaskReq struct {
	TaskData `json:"task_data"`
}

// SetTaskResp 响应消息
type SetTaskResp struct {
	RespBase
}
