package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
}

func (e ConcurrentEngine) Run(seed ...Request) {
	for _, r := range seed {
		e.Scheduler.Submit(r)
	}

	in := make(chan Request)
	out := make(chan ParserResult)

	for i:=0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for {
		result := <- out
		for _, item := range result.Items {
			fmt.Printf("Got item: %v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}
