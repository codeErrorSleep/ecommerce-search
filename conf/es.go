package conf

import (
	es "github.com/elastic/go-elasticsearch/v7"
)

var (
	EsClient *es.Client
)

// NewEs 初始化es
func NewEs() error {
	var err error
	EsClient, err = es.NewClient(es.Config{
		Addresses: []string{"http://localhost:9200"},
		// Username:  "username",
		// Password:  "password",
	})
	return err
}
