package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"GoStudy/Chapter17/crawlerPro/engine"
	"GoStudy/Chapter17/crawlerPro/persist"
	"fmt"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	fmt.Println("saving...")
	err := persist.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	fmt.Println(result)
	return err
}
