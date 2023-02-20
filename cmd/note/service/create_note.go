package service

import (
	"context"
	"github.com/qing-wq/easy-note/cmd/note/dal/db"
	"github.com/qing-wq/easy-note/kitex_gen/note"
)

type CreateNoteService struct {
	ctx context.Context
}

func NewCreateNoteService(ctx context.Context) *CreateNoteService {
	return &CreateNoteService{ctx: ctx}
}

func (s *CreateNoteService) CreateNote(req *note.CreateNoteRequest) error {
	noteModel := &db.Note{
		UserID:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	}
	return db.CreateNote(s.ctx, []*db.Note{noteModel})
}
