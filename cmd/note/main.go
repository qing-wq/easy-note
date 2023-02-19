package main

import (
	note "github.com/qing-wq/easy-note/kitex_gen/note/noteservice"
	"log"
)

func main() {
	svr := note.NewServer(new(NoteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
