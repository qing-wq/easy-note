package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"github.com/qing-wq/easy-note/cmd/api/rpc"
	"github.com/qing-wq/easy-note/kitex_gen/note"
	"github.com/qing-wq/easy-note/pkg/constants"
	"github.com/qing-wq/easy-note/pkg/errno"
	"strconv"
)

func DeleteNote(ctx context.Context, c *app.RequestContext) {
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))
	noteIDStr := c.Param(constants.NoteID)
	noteID, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if noteID <= 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.DeleteNote(context.Background(), &note.DeleteNoteRequest{
		NoteId: noteID, UserId: userID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
