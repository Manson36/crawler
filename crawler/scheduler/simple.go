package scheduler

import (
	"github.com/crawler/crawler/engine"
)

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.WorkerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	 go func() {
		 s.WorkerChan <- r
	 }()
}

