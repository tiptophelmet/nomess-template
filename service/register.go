package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tiptophelmet/nomess-core/logger"
	"github.com/tiptophelmet/nomess-core/password"
	"github.com/tiptophelmet/nomess-template/body"
	"github.com/tiptophelmet/nomess-template/errs"
	"github.com/tiptophelmet/nomess-template/model"
	"github.com/tiptophelmet/nomess-template/repo"
	"github.com/tiptophelmet/nomess-template/util"

	"github.com/go-playground/validator/v10"
)

type Register struct {
	userRepo             *repo.User
	userVerificationRepo *repo.UserVerification
	validate             *validator.Validate
}

func InitRegisterService() *Register {
	return &Register{
		userRepo:             repo.InitUserRepo(),
		userVerificationRepo: repo.InitUserVerificationRepo(),
		validate:             validator.New(),
	}
}

func (srv *Register) sendVerificationEmail(mailTo string, code string) error {
	// Use your email client to send verification email
	return nil
}

func (srv *Register) Validate(body body.Register) error {
	err := srv.validate.Struct(body)

	if err != nil {
		errs := make([]string, 0, len(err.(validator.ValidationErrors)))

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()

			// TODO: Consider adding err.Tag() to show what exactly failed:
			// err.Tag() returns very short values like: "lte", "iscolor", "required";
			// So it is better to format err.Tag() to a more human-friendly message
			errs = append(errs, "field: %v is invalid, ", err.Field())
		}

		return errors.New(strings.Join(errs, ""))
	}

	return nil
}

func (srv *Register) Register(body body.Register) error {
	passwHash, err := password.HashAndSalt(body.Password)
	if err != nil {
		logger.Fatal(err.Error())
		return err
	}

	user := &model.User{
		Email:        body.Email,
		PasswordHash: passwHash,
		AuthProvider: "default",
		Verified:     false,
	}

	if err := srv.userRepo.Save(user); err != nil {
		logger.Error(err.Error())
		return errs.ErrUserInsert
	}

	userVerification := &model.UserVerification{
		User: *user,
		Code: util.RandStringBytes(6),
	}

	err = srv.userVerificationRepo.Save(userVerification)
	if err != nil {
		logger.Error(err.Error())
		return errs.ErrUserVerificationInsert
	}

	err = srv.sendVerificationEmail(body.Email, userVerification.Code)
	if err != nil {
		logger.Error(err.Error())
		return errs.ErrVerificationEmailNotSent
	}

	return nil
}
