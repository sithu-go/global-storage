package logic

import (
	"context"

	"grpc-storage/service/file-storage/rpc/internal/svc"
	"grpc-storage/service/file-storage/rpc/types/filegrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletevideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletevideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletevideoLogic {
	return &DeletevideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletevideoLogic) Deletevideo(in *filegrpc.VideoDeleteRequest) (*filegrpc.VideoDeleteResponse, error) {
	// todo: add your logic here and delete this line

	return &filegrpc.VideoDeleteResponse{}, nil
}
