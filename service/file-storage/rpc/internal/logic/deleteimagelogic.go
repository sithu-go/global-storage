package logic

import (
	"context"
	"fmt"
	"os"

	"grpc-storage/service/file-storage/rpc/internal/svc"
	"grpc-storage/service/file-storage/rpc/types/filegrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteImageLogic {
	return &DeleteImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteImageLogic) DeleteImage(in *filegrpc.ImageDeleteRequest) (*filegrpc.ImageDeleteResponse, error) {
	// todo: add your logic here and delete this line

	filename := in.Filename

	err := os.Remove("./files/images/" + filename)

	if err != nil {
		fmt.Println("Error deleting file", err.Error())
		return &filegrpc.ImageDeleteResponse{
			Deleted: false,
		}, nil
	}

	return &filegrpc.ImageDeleteResponse{
		Deleted: true,
	}, nil
}
