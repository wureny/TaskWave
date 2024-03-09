package task

import "time"

type Itask interface {
	CreateTask() (string, error)
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

// each type of task should register here
var TaskHandlerMap = make(map[string]*TaskHandler, 0)

func RegisterHandler(handler *TaskHandler) {
	TaskHandlerMap[handler.TaskType] = handler
}

func GetHandler(taskType string) (*TaskHandler, int) {
	if _, ok := TaskHandlerMap[taskType]; ok {
		return TaskHandlerMap[taskType], 0
	}
	return nil, -1
}

func NewHandler(taskType string, newProc func() Itask) *TaskHandler {
	return &TaskHandler{
		TaskType: taskType,
		NewProc:  newProc,
	}
}

func (t *TaskBase) SetContextLocal(context interface{}) {
	t.ContextIntf = context
}

func (t *TaskBase) Base() *TaskBase {
	return t
}

// TODO:Need to support configurable initialization
func NewTaskBase() *TaskBase {
	return &TaskBase{}
}

func (t *TaskBase) CreateTask() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TaskBase) SetTask() error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskBase) ContextLoad() error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskBase) HandleProcess() error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskBase) HandleFinish() {
	//TODO implement me
	panic("implement me")
}

func (t *TaskBase) HandleFinishError() error {
	//TODO implement me
	panic("implement me")
}

func (t *TaskBase) HandleFailedMust() error {
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

type TaskHandler struct {
	TaskType string
	NewProc  func() Itask
}
