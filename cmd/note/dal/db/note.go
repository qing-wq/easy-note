package db

import (
	"context"
	"github.com/qing-wq/easy-note/pkg/constants"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n *Note) TableName() string {
	return constants.NoteTableName
}

// CreateNote create note info
func CreateNote(ctx context.Context, notes []*Note) error {
	if err := DB.WithContext(ctx).Create(notes).Error; err != nil {
		return err
	}
	return nil
}

func DeleteNote(ctx context.Context, noteID, userID int64) error {
	return DB.WithContext(ctx).Where("id = ? and user_id = ?", noteID, userID).Delete(&Note{}).Error
}

func MGetNote(ctx context.Context, noteIDs []int64) ([]*Note, error) {
	var notes []*Note
	if len(notes) == 0 {
		return notes, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", noteIDs).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func UpdateNote(ctx context.Context, noteID, UserID int64, title *string, content *string) error {
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = title
	}
	if content != nil {
		params["content"] = content
	}

	return DB.WithContext(ctx).Model(&Note{}).Where("note_id = ? and user_id = ?", noteID, UserID).Updates(params).Error
}

func QueryNote(ctx context.Context, userID int64, searchKey *string, limit, offset int) ([]*Note, int64, error) {
	var notes []*Note
	var total int64
	conn := DB.WithContext(ctx).Model(&Note{}).Where("user_id = ?", userID)

	if searchKey != nil {
		conn = conn.Where("title like ?", "%"+*searchKey+"%")
	}

	var err error
	if err = conn.Count(&total).Error; err != nil {
		return notes, total, err
	}

	if err = conn.Limit(limit).Offset(offset).Find(&notes).Error; err != nil {
		return notes, total, err
	}

	return notes, total, nil
}
