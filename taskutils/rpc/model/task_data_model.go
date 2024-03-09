package model

import "time"

type TaskData struct {
	UserId           string    `json:"user_id"`
	TaskId           string    `json:"task_id"`
	TaskType         string    `json:"task_type"`
	TaskStage        string    `json:"task_stage"`
	Status           int64     `json:"status"`
	Priority         *int64    `json:"priority"`
	CrtRetryNum      int64     `json:"crt_retry_num"`
	MaxRetryNum      int64     `json:"max_retry_num"`
	MaxRetryInterval int64     `json:"max_retry_interval"`
	ScheduleLog      string    `json:"schedule_log"`
	TaskContext      string    `json:"context"`
	OrderTime        int64     `json:"order_time"`
	CreateTime       time.Time `json:"create_time"`
	ModifyTime       time.Time `json:"modify_time"`
}
