package threadpool

import (
	"../util"
	"math/rand"
	"sync"
	"time"
)

var Tasks = []*Task{}
/*var Tasks = &TaskMutex{
	Mutex: sync.Mutex{},
	Tasks: []*Task{},
	Free:	true,
}*/
var Workers = &WorkerMutex{
	Mutex: sync.Mutex{},
	Map:   map[int]bool{},
}
var SigWork = false
var Handler func(map[string]interface{})
var WorkerCount = 0

func Init() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < util.DefaultConfig.GoroutineCount; i++ {
		go newWorker()
	}
}

func newWorker() {
	workerId := rand.Int()
	Workers.Mutex.Lock()
	Workers.Map[workerId] = true
	Workers.Mutex.Unlock()

	WorkerCount++

	for true {
		tasksLen := len(Tasks)

		if tasksLen > 0 && SigWork {
			SigWork = false

			task := Tasks[tasksLen - 1]
			Tasks = Tasks[:tasksLen - 1]

			Workers.Mutex.Lock()
			Workers.Map[workerId] = false
			Workers.Mutex.Unlock()

			Handler(task.Task)

			Workers.Mutex.Lock()
			Workers.Map[workerId] = true
			Workers.Mutex.Unlock()
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func CountFreeWorkers() int {
	amount := 0
	for _, state := range Workers.Map {
		if state {
			amount++
		}
	}

	return amount
}
