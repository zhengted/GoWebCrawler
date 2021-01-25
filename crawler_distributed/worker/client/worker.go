package client

import (
	"GoWebCrawler/crawler/engine"
	"GoWebCrawler/crawler_distributed/config"
	"GoWebCrawler/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(request engine.Request) (engine.ParseResult, error) {

		sReq := worker.SerializeRequest(request)
		var sRes worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sRes)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sRes), nil

	}
}
