package main

import (
	"github.com/qing-wq/easy-note/cmd/note/rpc"
	"github.com/qing-wq/easy-note/cmd/user/dal"
	note "github.com/qing-wq/easy-note/kitex_gen/note/noteservice"
	"github.com/qing-wq/easy-note/pkg/constants"
	"github.com/qing-wq/easy-note/pkg/tracer"
	"log"
)

func Init() {
	tracer.InitJaeger(constants.NoteServiceName)
	dal.Init()
	rpc.InitRPC()
}

func main() {

	svr := note.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
