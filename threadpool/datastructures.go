package threadpool

import "sync"

type Task struct {
	Task map[string]interface{}
}

type TaskMutex struct {
	Mutex sync.Mutex
	Tasks []*Task
	Free bool
}

type WorkerMutex struct {
	Mutex sync.Mutex
	Map map[int]bool
}
