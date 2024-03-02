package task

import "time"

type Itask interface {
	SetTask() error
	ContextLoad() error
	HandleProcess() error
	HandleFinish()
	HandleFinishError() error
	//HandleFailedMust if err, then change status from failed to processing
	HandleFailedMust() error
}

type TaskBase struct {
	Id               int64
	TaskId           string
	UserId           string
	Status           int64
	TaskType         string
	TaskStage        string
	TaskContext      string
	CrtRetryNum      int64
	MaxRetryNum      int64
	MaxRetryInterval int64
	ScheduleLog      *ScheduleLog
	ContextIntf      interface{}
	Priority         int64
	OrderTime        int64
	//TODO: handle time type
	CreateTime time.Time
	ModifyTime time.Time
}

func (t TaskBase) SetTask() error {
	//TODO implement me
	panic("implement me")
}

func (t TaskBase) ContextLoad() error {
	//TODO implement me
	panic("implement me")
}

func (t TaskBase) HandleProcess() error {
	//TODO implement me
	panic("implement me")
}

func (t TaskBase) HandleFinish() {
	//TODO implement me
	panic("implement me")
}

func (t TaskBase) HandleFinishError() error {
	//TODO implement me
	panic("implement me")
}

func (t TaskBase) HandleFailedMust() error {
	//TODO implement me
	panic("implement me")
}

type ScheduleLog struct {
	LastData     ScheduleData
	HistoryDatas []ScheduleData
}

// ScheduleData struct ScheduleData
type ScheduleData struct {
	TraceId string
	ErrMsg  string
	Cost    string
}
