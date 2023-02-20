package service

import (
	"context"
	"github.com/qing-wq/easy-note/cmd/note/dal/db"
	"github.com/qing-wq/easy-note/kitex_gen/note"
)

type UpdateNoteService struct {
	ctx context.Context
}

func NewUpdateNoteService(ctx context.Context) *UpdateNoteService {
	return &UpdateNoteService{ctx: ctx}
}

func (s *UpdateNoteService) UpdateNote(req *note.UpdateNoteRequest) error {
	return db.UpdateNote(s.ctx, req.NoteId, req.UserId, req.Title, req.Content)
}
