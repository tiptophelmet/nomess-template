package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-template/body"
	"github.com/tiptophelmet/nomess-template/model"
	"github.com/tiptophelmet/nomess-template/repo"

	"github.com/go-playground/validator/v10"
)

type UpdateItem struct {
	itemRepo *repo.Item
	validate *validator.Validate
}

func InitUpdateItemService() *UpdateItem {
	return &UpdateItem{
		itemRepo: repo.InitItemRepo(),
		validate: validator.New(),
	}
}

func (srv *UpdateItem) Validate(body body.UpdateItem) error {
	err := srv.validate.Struct(body)

	if err != nil {
		errs := make([]string, 0, len(err.(validator.ValidationErrors)))

		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Sprintf("field: %v(%v) is invalid ", err.Field(), err.Tag()))
		}

		return errors.New(strings.Join(errs, ""))
	}

	return nil
}

func (srv *UpdateItem) Update(itemId int, itemBody body.UpdateItem) (*model.Item, error) {
	foundItem, err := srv.itemRepo.Get(itemId)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	foundItem.Name = itemBody.Name
	foundItem.Description = itemBody.Description

	updated, err := srv.itemRepo.Save(foundItem)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return updated, nil
}
