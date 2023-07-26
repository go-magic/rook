package mkafka

import "sync"

var (
	msgCenter     *MsgCenter
	onceMsgCenter sync.Once
)

type MsgCenter struct {
	lock              sync.Mutex
	taskHandleFuncMap map[string]TaskHandlerFunc
}

func GetMsgCenter() *MsgCenter {
	onceMsgCenter.Do(func() {
		msgCenter = &MsgCenter{
			taskHandleFuncMap: make(map[string]TaskHandlerFunc),
		}
	})
	return msgCenter
}

func (m *MsgCenter) Register(taskId string, taskHandlerFunc TaskHandlerFunc) {
	m.lock.Lock()
	m.taskHandleFuncMap[taskId] = taskHandlerFunc
	m.lock.Unlock()
}

func (m *MsgCenter) Handle(task *Task) {
	m.lock.Lock()
	h, ok := m.taskHandleFuncMap[task.TaskId]
	m.lock.Unlock()
	if ok {
		go m.do(task.Data, h)
	}
}

func (m *MsgCenter) do(data interface{}, handleFunc TaskHandlerFunc) {
	handleFunc().Do(data)
}
