package service

import (
	"context"
	"github.com/qing-wq/easy-note/cmd/note/dal/db"
	"github.com/qing-wq/easy-note/kitex_gen/note"
)

type DelNoteService struct {
	ctx context.Context
}

func NewDelNoteService(ctx context.Context) *DelNoteService {
	return &DelNoteService{
		ctx: ctx,
	}
}

func (s *DelNoteService) DelNote(req *note.DeleteNoteRequest) error {
	return db.DeleteNote(s.ctx, req.NoteId, req.UserId)
}
