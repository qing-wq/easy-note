package service

import (
	"context"
	"github.com/qing-wq/easy-note/cmd/note/dal/db"
	"github.com/qing-wq/easy-note/cmd/note/pack"
	"github.com/qing-wq/easy-note/cmd/note/rpc"
	"github.com/qing-wq/easy-note/kitex_gen/note"
	"github.com/qing-wq/easy-note/kitex_gen/user"
)

type MGetNoteService struct {
	ctx context.Context
}

// NewMGetNoteService new MGetNoteService
func NewMGetNoteService(ctx context.Context) *MGetNoteService {
	return &MGetNoteService{ctx: ctx}
}

func (s *MGetNoteService) MGetNote(req *note.MGetNoteRequest) ([]*note.Note, error) {
	noteModel, err := db.MGetNote(s.ctx, req.NoteIds)
	if err != nil {
		return nil, err
	}
	uIds := pack.UserIds(noteModel)
	userMap, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserIds: uIds})
	if err != nil {
		return nil, err
	}
	notes := pack.Notes(noteModel)
	for i := 0; i < len(notes); i++ {
		if u, ok := userMap[notes[i].UserId]; ok {
			notes[i].UserName = u.Username
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, nil
}
