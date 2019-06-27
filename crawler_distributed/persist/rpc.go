package persist

import (
	"github.com/crawler/crawler/engine"
	"github.com/crawler/crawler/persist"
)

type ItemSaverService struct {
	Client *elastice.client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	if err == nil {
		*result = "ok"
	}
	return err
}