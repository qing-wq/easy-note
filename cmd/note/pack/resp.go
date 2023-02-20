package pack

import (
	"errors"
	"github.com/qing-wq/easy-note/kitex_gen/note"
	"github.com/qing-wq/easy-note/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *note.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *note.BaseResp {
	return &note.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
