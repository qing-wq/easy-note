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

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (int64,error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}

	passWord := fmt.Sprintf("%x",h.Sum(nil))
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	u := users[0]
	if u.PassWord != passWord {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(u.ID), nil
}
