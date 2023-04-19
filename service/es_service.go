package service

import (
	"bytes"
	"context"
	"ecommerce-search/conf"
	"encoding/json"
	"fmt"
	"log"
)

func GetInfoBySpuID(spuID string) (map[string]any, error) {
	// ctx := context.TODO()
	// 连接es直接查出来
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"spu_id": "c11a8df3-5eb",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	// Perform the search request.
	es := conf.EsClient
	res, err := conf.EsClient.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("spu"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return nil, err
	}

	r := map[string]any{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	fmt.Println(r["hits"])
	// spu := entity.Spu{}
	// if err := json.NewDecoder(r["hits"]).Decode(&spu); err != nil {
	// 	log.Fatalf("Error parsing the spu: %s", err)
	// }

	return r, nil
}
