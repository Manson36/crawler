package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier

	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seed ...Request) {
	out := make(chan ParserResult)

	e.Scheduler.Run()

	for i:=0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seed {
		if isDuplicate(r.Url) {
			continue
		}

		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v",itemCount, item)
			itemCount++
		}

		//URl 去重
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}

			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			//tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

//最简单的去重方法，哈希表；缺点：占用空间大
//缺点：每次重启，数据会丢失；方法：每次结束都将数据存储，或者利用外部数据库redis等工具存储
var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}

	visitedUrl[url] = true //如果输入的URL在map中不存在，则返回初始值false
	return false
}
