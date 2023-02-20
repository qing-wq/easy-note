package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/qing-wq/easy-note/kitex_gen/user"
	"github.com/qing-wq/easy-note/kitex_gen/user/userservice"
	"github.com/qing-wq/easy-note/pkg/constants"
	"github.com/qing-wq/easy-note/pkg/errno"
	"github.com/qing-wq/easy-note/pkg/middleware"
	"time"
)

var userClient userservice.Client

func InitUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})

	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleWare),
		client.WithInstanceMW(middleware.ClientMiddleWare),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func MGetUser(ctx context.Context, req *user.MGetUserRequest) (map[int64]*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := make(map[int64]*user.User, 0)
	for _, u := range resp.Users {
		res[u.UserId] = u
	}
	return res, nil
}
