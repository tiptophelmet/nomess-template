package service

import (
	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-template/repo"
)

type DeleteItem struct {
	itemRepo *repo.Item
}

func InitDeleteItemService() *DeleteItem {
	return &DeleteItem{
		itemRepo: repo.InitItemRepo(),
	}
}

func (srv *DeleteItem) Delete(ID int) error {
	err := srv.itemRepo.Delete(ID)
	if err != nil {
		logger.Error(err.Error())
	}

	return err
}
