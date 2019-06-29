package client

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler_distributed/config"
	"github.com/crawler/crawler_distributed/rpcsupport"
	"github.com/crawler/crawler_distributed/worker"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(config.WorkerPort0)
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParserResult, error) {
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParserResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
