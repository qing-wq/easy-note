package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qing-wq/easy-note/cmd/api/rpc"
	"github.com/qing-wq/easy-note/kitex_gen/user"
	"github.com/qing-wq/easy-note/pkg/errno"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserParam
	if err := c.Bind(&registerVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, errno.UserAlreadyExistErr, nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
