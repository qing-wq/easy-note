package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/qing-wq/easy-note/cmd/user/dal"
	user "github.com/qing-wq/easy-note/kitex_gen/user/userservice"
	"github.com/qing-wq/easy-note/pkg/bound"
	"github.com/qing-wq/easy-note/pkg/constants"
	"github.com/qing-wq/easy-note/pkg/middleware"
	"github.com/qing-wq/easy-note/pkg/tracer"
	"net"
)

func Init() {
	tracer.InitJaeger(constants.UserServiceName)
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.UserServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
		server.WithMiddleware(middleware.CommonMiddleWare),
		server.WithMiddleware(middleware.ServerMiddleWare),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithBoundHandler(bound.NewCpuLimitHandler()),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
