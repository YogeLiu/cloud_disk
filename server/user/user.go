package user

import (
	"context"
	"errors"

	"github.com/YogeLiu/CloudDisk/dao"
	"github.com/YogeLiu/CloudDisk/pkg/secret"
	"github.com/YogeLiu/CloudDisk/pkg/util"
)

type UserService struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (svc *UserService) Create(ctx context.Context) error {
	password, err := secret.SetPassword(svc.Password)
	if err != nil {
		util.Log().Error(err.Error())
		return err
	}
	err = dao.AddUser(ctx, &dao.User{Name: svc.Account, Password: password})
	if err != nil {
		util.Log().Error(err.Error())
	}
	return err
}

func (svc *UserService) Get(ctx context.Context) (dao.User, error) {
	user, err := dao.GetUserByName(ctx, svc.Account)
	if err != nil {
		util.Log().Error(err.Error())
	}
	return user, err
}

func (svc *UserService) Login(ctx context.Context, user *dao.User) (string, error) {
	ok, err := secret.CheckPassword(svc.Password, user.Password)
	if err != nil {
		util.Log().Error(err.Error())
		return "", err
	}
	if !ok {
		return "", errors.New("check password failure")
	}
	token := secret.EnJWT(user.ID)
	return token, nil
}
