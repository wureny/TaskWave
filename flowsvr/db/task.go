package db

import "time"

// Task 任务表
type Task struct {
	Id               int64     `gorm:"primary_key;AUTO_INCREMENT:15"`
	UserId           string    `gorm:"type:varchar(256);column:user_id;not null;default:'';index:idx_user_id"`
	TaskId           string    `gorm:"type:varchar(256);column:task_id;not null;default:'';unique_index:idx_task_id"`
	TaskType         string    `gorm:"type:varchar(128);column:task_type;not null"`
	TaskStage        string    `gorm:"type:varchar(128);column:task_stage;not null"`
	Status           int64     `gorm:"type:tinyint(3);column:status;not null;default:1;index:idx_tasktype_status_modify_time"`
	Priority         int64     `gorm:"type:int(11);column:priority;not null;default:0;comment:'优先级'"`
	CrtRetryNum      int64     `gorm:"type:int(11);column:crt_retry_num;not null;default:0;comment:'已经重试几次了'"`
	MaxRetryNum      int64     `gorm:"type:int(11);column:max_retry_num;not null;default:0;comment:'最大能重试几次'"`
	MaxRetryInterval int64     `gorm:"type:int(11);column:max_retry_interval;not null;default:0;comment:'最大重试间隔'"`
	ScheduleLog      string    `gorm:"type:varchar(4096);column:schedule_log;not null;default:'';comment:'调度信息记录'"`
	TaskContext      string    `gorm:"type:varchar(8192);column:task_context;not null;default:'';comment:'任务上下文，用户自定义'"`
	OrderTime        int64     `gorm:"type:int(20);column:order_time;not null;default:0;comment:'调度时间，越小调度越优先';index:idx_tasktype_status_modify_time"`
	CreateTime       time.Time `gorm:"type:datetime;column:create_time;not null;autoCreateTime;"`
	ModifyTime       time.Time `gorm:"type:datetime;column:modify_time;not null;autoCreateTime;"`
}
