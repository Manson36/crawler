package client

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler_distributed/config"
	"github.com/crawler/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParserResult, error) {
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <- clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParserResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
