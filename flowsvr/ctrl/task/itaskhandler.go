package task

type ITaskHandler interface {
	HandleInput() error
	HandleProcess() error
}
