package service

import (
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-template/model"
	"github.com/tiptophelmet/nomess-template/repo"
)

type GetItem struct {
	itemRepo *repo.Item
}

func InitGetItemService() *GetItem {
	return &GetItem{
		itemRepo: repo.InitItemRepo(),
	}
}

func (srv *GetItem) Get(ID int) (*model.Item, error) {
	item, err := srv.itemRepo.Get(ID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return item, nil
}
