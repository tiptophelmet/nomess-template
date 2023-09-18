package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tiptophelmet/nomess-core/v4/logger"
	"github.com/tiptophelmet/nomess-template/body"
	"github.com/tiptophelmet/nomess-template/model"
	"github.com/tiptophelmet/nomess-template/repo"

	"github.com/go-playground/validator/v10"
)

type CreateItem struct {
	itemRepo *repo.Item
	validate *validator.Validate
}

func InitCreateItemService() *CreateItem {
	return &CreateItem{
		itemRepo: repo.InitItemRepo(),
		validate: validator.New(),
	}
}

func (srv *CreateItem) Validate(body body.CreateItem) error {
	err := srv.validate.Struct(body)

	if err != nil {
		errs := make([]string, 0, len(err.(validator.ValidationErrors)))

		for _, err := range err.(validator.ValidationErrors) {
			// TODO: Consider adding err.Tag() to show what exactly failed:
			// err.Tag() returns very short values like: "lte", "iscolor", "required";
			// So it is better to format err.Tag() to a more human-friendly message
			errs = append(errs, fmt.Sprintf("field: %v(%v) is invalid ", err.Field(), err.Tag()))
		}

		return errors.New(strings.Join(errs, ""))
	}

	return nil
}

func (srv *CreateItem) Create(itemBody body.CreateItem) (*model.Item, error) {
	user := &model.Item{
		Name:        itemBody.Name,
		Description: itemBody.Description,
	}

	created, err := srv.itemRepo.Save(user)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return created, nil
}
