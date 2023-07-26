package mkafka

type Task struct {
	TaskId string
	Data   interface{}
}

type TaskHandler interface {
	Do(interface{})
}

type TaskHandlerFunc func() TaskHandler

func NewTask(taskId string, data interface{}) *Task {
	return &Task{TaskId: taskId, Data: data}
}
