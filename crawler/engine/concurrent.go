package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
	//分布式  在这里添加字段，外面就可以去配置了
	RequestProcessor Processor
}
//我们在这里把Worker的函数类型定义在这里
type Processor func(Request) (ParserResult, error)

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
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)//e 的肚子里有Processor的函数
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
			go func(i Item) {
				e.ItemChan <- i
			}(item)
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

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			//tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <- in
			result, err := e.RequestProcessor(request)//分布式，将这里的worker函数
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
