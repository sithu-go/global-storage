package logic

import (
	"context"

	"grpc-storage/service/file-storage/rpc/internal/svc"
	"grpc-storage/service/file-storage/rpc/types/filegrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetvideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetvideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetvideoLogic {
	return &GetvideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetvideoLogic) Getvideo(in *filegrpc.VideoGetRequest) (*filegrpc.VideoGetResponse, error) {
	// todo: add your logic here and delete this line

	return &filegrpc.VideoGetResponse{}, nil
}
