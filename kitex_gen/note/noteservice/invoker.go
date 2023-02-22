// Code generated by Kitex v0.4.4. DO NOT EDIT.

package noteservice

import (
	server "github.com/cloudwego/kitex/server"
	note "github.com/qing-wq/easy-note/kitex_gen/note"
)

// NewInvoker creates a server.Invoker with the given handlers and options.
func NewInvoker(handler note.NoteService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
