package elk

import (
	"context"
	"jettster/provider/config"

	"github.com/olivere/elastic"
)

var client *elastic.Client

func init() {
	var cl, err = elastic.NewSimpleClient(elastic.SetURL(config.GetString("elastic.url")))
	if err != nil {
		panic(err)
	}
	client = cl
}

func Connect() *elastic.Client {
	return client
}

func CreateIndex(name string) {
	_, _ = client.CreateIndex(name).Do(context.Background())
}

func AddData(index string, docType string, body interface{}) *elastic.IndexResponse {
	result, err := client.Index().
		Index(index).
		Type(docType).
		BodyJson(body).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return result
}

func UpsertData(index string, docType string, id string, body interface{}) *elastic.IndexResponse {
	result, err := client.Index().
		Index(index).
		Type(docType).
		Id(id).
		BodyJson(body).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	return result
}
