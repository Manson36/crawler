package config

const (
	//Parser names
	ParseCity   = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser 	 = "NilParser"

	//ElasticSearch
	ElasticIndex = "dating_profile"

	//RPC Endpoint
	ItemSaverRpc = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	//rate limiting
	Qps = 20
)
