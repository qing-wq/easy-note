package service

import (
	"context"
	"github.com/qing-wq/easy-note/cmd/user/dal/db"
	"github.com/qing-wq/easy-note/cmd/user/pack"
	"github.com/qing-wq/easy-note/kitex_gen/user"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

func (s *MGetUserService) MGetUser(req *user.MGetUserRequest) ([]*user.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), err
}