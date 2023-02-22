package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"github.com/qing-wq/easy-note/cmd/api/rpc"
	"github.com/qing-wq/easy-note/kitex_gen/note"
	"github.com/qing-wq/easy-note/pkg/constants"
	"github.com/qing-wq/easy-note/pkg/errno"
)

func CreateNote(ctx context.Context, c *app.RequestContext)  {
	var noteVar NoteParam
	if err := c.Bind(&noteVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(noteVar.Title) == 0 || len(noteVar.Content) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))
	err := rpc.CreateNote(context.Background(), &note.CreateNoteRequest{
		UserId: userID,
		Content:  noteVar.Content,
		Title: noteVar.Title,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(c, errno.Success, nil)
}
