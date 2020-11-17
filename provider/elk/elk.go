package elk

import (
	"context"

	"github.com/olivere/elastic"
)

var client *elastic.Client

func init() {
	var cl, err = elastic.NewSimpleClient(elastic.SetURL("https://67yzvvqhca:bwljjiblg1@gin-6369977968.us-east-1.bonsaisearch.net:443"))
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
