package task

import (
	"github.com/astaxie/beego/toolbox"
)

var (
	Tasks map[string]Task
)

func init() {
	Tasks = make(map[string]Task)
}

type Task interface {
	Name() string
	Spec() string
	TaskFunc() error
	Shutdown() error
}

func Register(t Task) {
	Tasks[t.Name()] = t
}

func Start() {
	for _, v := range Tasks {
		task := toolbox.NewTask(v.Name(), v.Spec(), v.TaskFunc)
		task.Run()
		toolbox.AddTask(v.Name(), task)
	}

	toolbox.StartTask()
}

func Stop() {
	toolbox.StopTask()
}

func Shutdown() {
	toolbox.StopTask()
	for _, v := range Tasks {
		v.Shutdown()
	}
}
