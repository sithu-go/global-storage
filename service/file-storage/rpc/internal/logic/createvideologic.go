package logic

import (
	"context"

	"grpc-storage/service/file-storage/rpc/internal/svc"
	"grpc-storage/service/file-storage/rpc/types/filegrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatevideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatevideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatevideoLogic {
	return &CreatevideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// video grpc call
func (l *CreatevideoLogic) Createvideo(in *filegrpc.VidoeCreateRequest) (*filegrpc.VideoCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &filegrpc.VideoCreateResponse{}, nil
}
