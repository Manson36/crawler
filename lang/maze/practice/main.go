package practice

type request struct {

}

type queueEngine struct {
	workerChan chan chan request
	requestChan chan request
}

func (e queueEngine) Run() {
	e.workerChan = make(chan chan request)
	e.requestChan = make(chan request)

	go func() {
		var workerQ []chan request
		var requestQ []request

		for {
			var activerequest request
			var activeworker chan request

			if len(workerQ) > 0 && len(requestQ) > 0 {
				activerequest = requestQ[0]
				activeworker = workerQ[0]
			}

			select {
			case r := <- e.requestChan:
				requestQ = append(requestQ, r)
			case w := <-e.workerChan:
				workerQ = append(workerQ, w)
			case activeworker <- activerequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
