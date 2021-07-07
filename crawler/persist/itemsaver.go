package persist

import (
	"context"
	"crawler/crawler/engine"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver(indexName string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	elasticClient, err := elastic.NewClient(
		elastic.SetURL("http://192.168.58.130:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSave Got Item #%d :%v", itemCount, item)
			err := save(elasticClient, indexName, item)
			if err != nil {
				log.Printf("Item Saver Error(%v): %v", item, err)
			}
			itemCount++
		}
	}()
	return out, nil
}

func save(client *elastic.Client, indexName string, item engine.Item) (err error) {

	indexService := client.Index()
	indexService.Index(indexName).Type("_doc")
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
