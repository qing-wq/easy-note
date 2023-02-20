package pack

import (
	"github.com/qing-wq/easy-note/cmd/note/dal/db"
	"github.com/qing-wq/easy-note/kitex_gen/note"
)

func Note(m *db.Note) *note.Note {
	if m == nil {
		return nil
	}

	return &note.Note{
		NoteId:     int64(m.ID),
		UserId:     m.UserID,
		Title:      m.Title,
		Content:    m.Content,
		CreateTime: m.CreatedAt.Unix(),
	}
}

func Notes(notes []*db.Note) []*note.Note {
	res := make([]*note.Note, 0)
	for _, n := range notes {
		if m := Note(n); m != nil {
			res = append(res, m)
		}
	}
	return res
}

func UserIds(ms []*db.Note) []int64 {
	uIds := make([]int64, 0)
	if len(ms) == 0 {
		return uIds
	}
	uIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			uIdMap[m.UserID] = struct{}{}
		}
	}
	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}
