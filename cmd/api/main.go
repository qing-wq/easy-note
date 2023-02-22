package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/jwt"
	"github.com/qing-wq/easy-note/cmd/api/rpc"
	"github.com/qing-wq/easy-note/pkg/constants"
	"github.com/qing-wq/easy-note/pkg/tracer"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := server.New(
		server.WithHostPorts(constants.ApiServiceAddr),
		server.WithHandleMethodNotAllowed(true),
	)
	authorMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
		Key:         nil,
		Timeout:     0,
		MaxRefresh:  0,
		PayloadFunc: nil,
	})
}