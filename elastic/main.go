package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

type Tweet struct {
	User string
	Message string
	Retweet int
	Image string
	Created time.Time
	Tags []string
	Location string
	Suggest *elastic.SuggestField
}

const mapping =`
{
    "settings":{
        "number_of_shards":1,
        "number_of_replicas":0
    },
    "mappings":{
        "properties":{
            "user":{
                "type":"keyword"
            },
            "message":{
                "type":"text",
                "store":true,
                "fielddata":true
            },
            "image":{
                "type":"keyword"
            },
            "created":{
                "type":"date"
            },
            "tags":{
                "type":"keyword"
            },
            "location":{
                "type":"geo_point"
            },
            "suggest_field":{
                "type":"completion"
            }
        }
    }
}`

func main() {
	
}

func NewClient(ctx context.Context) *elastic.Client {
	client, err := elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("http://39.105.203.225:9200"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	esversion, err := client.ElasticsearchVersion("http://39.105.203.225:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	return client
}

func createIndex(ctx context.Context, client *elastic.Client) {
	exists, err := client.IndexExists("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}

	if !exists {
		//create a new index
		createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}

		if !createIndex.Acknowledged {
			//not acknowledged
		}
	}
}

func deleteIndex(ctx context.Context, client *elastic.Client) {
	deleteIndex, err := client.DeleteIndex("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}

	if !deleteIndex.Acknowledged {

	}
}

func addData(ctx context.Context, client *elastic.Client) {
	//使用json序列化的方式
	tweet1 := Tweet{User:"olivere", Message:"Take Five", Retweet: 0}
	put1, err := client.Index().Id("1").BodyJson(tweet1).Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Indexed tweet %d to index %s, type %s\n",
		put1.Id, put1.Index, put1.Type)

	tweet2 := `{"user": "olivere", "message":"It is a Raggy Waltz"}`
	put2, err := client.Index().Id("2").BodyString(tweet2).Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Indexed tweet %d to index %s, type %s\n",
		put2.Id, put2.Index, put2.Type)

	_, err = client.Flush().Index("twitter").Do(ctx)
	if err !=nil {
		panic(err)
	}
}

func getData(ctx context.Context, client *elastic.Client) {
	get1, err := client.Get().Index("twiteer").Id("1").Do(ctx)
	if err != nil {
		panic(err)
	}

	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n",
			get1.Id, get1.Version, get1.Index, get1.Type)
		tweet := Tweet{}
		err = json.Unmarshal(get1.Source, &tweet)
		if err != nil {
			panic(err)
		}
		fmt.Println(tweet)
	}
}