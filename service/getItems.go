package service

import (
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-template/model"
	"github.com/tiptophelmet/nomess-template/repo"
)

type GetItems struct {
	itemRepo *repo.Item
}

func InitGetItemsService() *GetItems {
	return &GetItems{
		itemRepo: repo.InitItemRepo(),
	}
}

func (srv *GetItems) GetAll() ([]*model.Item, error) {
	items, err := srv.itemRepo.GetAll()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return items, nil
}
