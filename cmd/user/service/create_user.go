package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/qing-wq/easy-note/cmd/user/dal/db"
	"github.com/qing-wq/easy-note/kitex_gen/user"
	"github.com/qing-wq/easy-note/pkg/errno"
	"io"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil	{
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}

	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.Username,
		PassWord: passWord,
	}})
}