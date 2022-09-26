package logic

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"grpc-storage/service/file-storage/rpc/internal/svc"
	"grpc-storage/service/file-storage/rpc/types/filegrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateImageLogic {
	return &CreateImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// image grpc call
func (l *CreateImageLogic) CreateImage(in *filegrpc.ImageCreateRequest) (*filegrpc.ImageResponse, error) {
	// todo: add your logic here and delete this line
	file, err := ioutil.TempFile("./files/images", "img-*"+in.Ext)
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
		return &filegrpc.ImageResponse{}, err
	}
	fmt.Println(len(in.FileData))
	if _, err := file.Write(in.FileData); err != nil {
		fmt.Println("Error Writing file:", err.Error())
		return &filegrpc.ImageResponse{}, err
	}
	filenames := strings.Split(file.Name(), "/")
	fmt.Println(filenames)
	return &filegrpc.ImageResponse{
		Filename: filenames[len(filenames)-1],
	}, nil
}
