package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/qing-wq/easy-note/kitex_gen/note"
	"github.com/qing-wq/easy-note/kitex_gen/note/noteservice"
	"github.com/qing-wq/easy-note/pkg/constants"
	"github.com/qing-wq/easy-note/pkg/errno"
	"github.com/qing-wq/easy-note/pkg/middleware"
	"time"
)

var noteClient noteservice.Client

func InitNoteRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})

	if err != nil {
		panic(err)
	}

	c, err := noteservice.NewClient(
		constants.NoteServiceName,
		client.WithMiddleware(middleware.CommonMiddleWare),
		client.WithInstanceMW(middleware.ClientMiddleWare),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	noteClient = c
}

func CreateNote(ctx context.Context, req *note.CreateNoteRequest) error {
	resp, err := noteClient.CreateNote(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return  errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// UpdateNote update note info
func UpdateNote(ctx context.Context, req *note.UpdateNoteRequest) error {
	resp, err := noteClient.UpdateNote(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// DeleteNote delete note info
func DeleteNote(ctx context.Context, req *note.DeleteNoteRequest) error {
	resp, err := noteClient.DeleteNote(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func QueryNotes(ctx context.Context, req *note.QueryNoteRequest) ([]*note.Note, int64, error) {
	resp, err := noteClient.QueryNote(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Notes, resp.Total, nil
}
