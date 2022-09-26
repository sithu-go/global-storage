package logic

import (
	"context"
	"fmt"
	"os"

	"grpc-storage/service/file-storage/rpc/internal/svc"
	"grpc-storage/service/file-storage/rpc/types/filegrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImageLogic {
	return &GetImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetImageLogic) GetImage(in *filegrpc.ImageGetRequest) (*filegrpc.ImageGetResponse, error) {
	// todo: add your logic here and delete this line
	filename := in.Filename
	data, err := os.ReadFile(fmt.Sprintf("./files/images/%v", filename))
	if err != nil {
		fmt.Println("ERROR reading file", err)
		return &filegrpc.ImageGetResponse{
			FileData: data,
		}, err
	}

	return &filegrpc.ImageGetResponse{
		FileData: data,
	}, nil
}
