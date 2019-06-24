package scheduler

import (
	"github.com/crawler/crawler/engine"
)

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	s.WorkerChan <- r
}

