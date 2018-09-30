package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"GoStudy/Chapter17/crawlerPro/engine"
	"errors"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://100.100.16.55:9200"), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("saver recieved:%d %v \n", itemCount, item)
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Save error:saving item %v,error:%v", item, err)
			}
			itemCount++
		}
	}()
	return out, err
}

func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}

	indexService := client.Index().Index(index).
		Type(item.Type).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf("%+v",resp)

	return nil
}
