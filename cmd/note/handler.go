package main

import (
	"context"
	"github.com/qing-wq/easy-note/cmd/note/pack"
	"github.com/qing-wq/easy-note/cmd/note/service"
	note "github.com/qing-wq/easy-note/kitex_gen/note"
	"github.com/qing-wq/easy-note/pkg/constants"
	"github.com/qing-wq/easy-note/pkg/errno"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *note.CreateNoteRequest) (resp *note.CreateNoteResponse, err error) {
	resp = new(note.CreateNoteResponse)

	if req.UserId <= 0 || len(req.Title) == 0 || len(req.Content) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateNoteService(ctx).CreateNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) MGetNote(ctx context.Context, req *note.MGetNoteRequest) (resp *note.MGetNoteResponse, err error) {
	resp = new(note.MGetNoteResponse)

	if len(req.NoteIds) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	notes, err := service.NewMGetNoteService(ctx).MGetNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Notes = notes
	return resp, nil
}

// DeleteNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteNote(ctx context.Context, req *note.DeleteNoteRequest) (resp *note.DeleteNoteResponse, err error) {
	resp = new(note.DeleteNoteResponse)

	if req.NoteId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewDelNoteService(ctx).DelNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// QueryNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) QueryNote(ctx context.Context, req *note.QueryNoteRequest) (resp *note.QueryNoteResponse, err error) {
	resp = new(note.QueryNoteResponse)

	if req.UserId <= 0 || req.Limit < 0 || req.Offset < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	if req.Limit == 0 {
		req.Limit = constants.DefaultLimit
	}

	notes, total, err := service.NewQueryNoteService(ctx).QueryNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Notes = notes
	resp.Total = total
	return resp, nil
}

// UpdateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req *note.UpdateNoteRequest) (resp *note.UpdateNoteResponse, err error) {
	resp = new(note.UpdateNoteResponse)

	if req.NoteId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateNoteService(ctx).UpdateNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
