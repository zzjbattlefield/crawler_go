package persist

import (
	"context"
	"crawler/crawler/engine"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("ItemSave Got Item #%d :%v", itemCount, item)
			err := save(item)
			if err != nil {
				log.Printf("Item Saver Error(%v): %v", item, err)
			}
			itemCount++
		}
	}()
	return out
}

func save(item engine.Item) (err error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.58.130:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	indexService := client.Index()
	indexService.Index("dating_profile").Type("_doc")
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
