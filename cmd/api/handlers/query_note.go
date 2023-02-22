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

func QueryNote(ctx context.Context, c *app.RequestContext) {
	claims := jwt.ExtractClaims(ctx,c)
	userID := int64(claims[constants.IdentityKey].(float64))
	var queryVar struct {
		Limit         int64  `json:"limit" form:"limit" query:"limit"`
		Offset        int64  `json:"offset" form:"offset" query:"offset"`
		SearchKeyword string `json:"search_keyword" form:"search_keyword" query:"search_keyword"`
	}
	if err := c.Bind(&queryVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if queryVar.Limit < 0 || queryVar.Offset < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &note.QueryNoteRequest{
		UserId:    userID,
		Offset:    queryVar.Offset,
		Limit:     queryVar.Limit,
	}

	if len(queryVar.SearchKeyword) != 0 {
		req.SearchKey = &queryVar.SearchKeyword
	}
	notes, total, err := rpc.QueryNotes(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{constants.Total: total, constants.Notes: notes})
}
