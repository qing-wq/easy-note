package main

import (
	"context"
	"github.com/qing-wq/easy-note/cmd/note/pack"
	"github.com/qing-wq/easy-note/cmd/note/service"
	note "github.com/qing-wq/easy-note/kitex_gen/note"
	"github.com/qing-wq/easy-note/pkg/errno"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *note.CreateNoteRequest) (resp *note.CreateNoteResponse, err error) {
	// TODO: Your code here...
	resp = new(note.CreateNoteResponse)
	if req.UserId <= 0 || len(req.Title) == 0 || len(req.Content) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewCreateNoteService(ctx).CreateNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) MGetNote(ctx context.Context, req *note.MGetNoteRequest) (resp *note.MGetNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteNote(ctx context.Context, req *note.DeleteNoteRequest) (resp *note.DeleteNoteResponse, err error) {
	// TODO: Your code here...
	resp = new(note.DeleteNoteResponse)

	if req.NoteId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	if err = service.NewDelNoteService(ctx).DelNote(req); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, err
}

// QueryNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) QueryNote(ctx context.Context, req *note.QueryNoteRequest) (resp *note.QueryNoteResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req *note.UpdateNoteRequest) (resp *note.UpdateNoteResponse, err error) {
	// TODO: Your code here...
	resp = new(note.UpdateNoteResponse)

	if req.UserId <= 0 || req.NoteId <= 0 || req.Title == nil || req.Content == nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	if err = service.NewUpdateNoteService(ctx).UpdateNote(req); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
