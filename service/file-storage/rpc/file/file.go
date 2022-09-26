// Code generated by goctl. DO NOT EDIT!
// Source: file.proto

package file

import (
	"context"

	"grpc-storage/service/file-storage/rpc/types/filegrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ImageCreateRequest  = filegrpc.ImageCreateRequest
	ImageDeleteRequest  = filegrpc.ImageDeleteRequest
	ImageDeleteResponse = filegrpc.ImageDeleteResponse
	ImageGetRequest     = filegrpc.ImageGetRequest
	ImageGetResponse    = filegrpc.ImageGetResponse
	ImageResponse       = filegrpc.ImageResponse
	VideoCreateResponse = filegrpc.VideoCreateResponse
	VideoDeleteRequest  = filegrpc.VideoDeleteRequest
	VideoDeleteResponse = filegrpc.VideoDeleteResponse
	VideoGetRequest     = filegrpc.VideoGetRequest
	VideoGetResponse    = filegrpc.VideoGetResponse
	VidoeCreateRequest  = filegrpc.VidoeCreateRequest

	File interface {
		//  image grpc call
		CreateImage(ctx context.Context, in *ImageCreateRequest, opts ...grpc.CallOption) (*ImageResponse, error)
		GetImage(ctx context.Context, in *ImageGetRequest, opts ...grpc.CallOption) (*ImageGetResponse, error)
		DeleteImage(ctx context.Context, in *ImageDeleteRequest, opts ...grpc.CallOption) (*ImageDeleteResponse, error)
		// video grpc call
		Createvideo(ctx context.Context, in *VidoeCreateRequest, opts ...grpc.CallOption) (*VideoCreateResponse, error)
		Getvideo(ctx context.Context, in *VideoGetRequest, opts ...grpc.CallOption) (*VideoGetResponse, error)
		Deletevideo(ctx context.Context, in *VideoDeleteRequest, opts ...grpc.CallOption) (*VideoDeleteResponse, error)
	}

	defaultFile struct {
		cli zrpc.Client
	}
)

func NewFile(cli zrpc.Client) File {
	return &defaultFile{
		cli: cli,
	}
}

// image grpc call
func (m *defaultFile) CreateImage(ctx context.Context, in *ImageCreateRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	client := filegrpc.NewFileClient(m.cli.Conn())
	return client.CreateImage(ctx, in, opts...)
}

func (m *defaultFile) GetImage(ctx context.Context, in *ImageGetRequest, opts ...grpc.CallOption) (*ImageGetResponse, error) {
	client := filegrpc.NewFileClient(m.cli.Conn())
	return client.GetImage(ctx, in, opts...)
}

func (m *defaultFile) DeleteImage(ctx context.Context, in *ImageDeleteRequest, opts ...grpc.CallOption) (*ImageDeleteResponse, error) {
	client := filegrpc.NewFileClient(m.cli.Conn())
	return client.DeleteImage(ctx, in, opts...)
}

// video grpc call
func (m *defaultFile) Createvideo(ctx context.Context, in *VidoeCreateRequest, opts ...grpc.CallOption) (*VideoCreateResponse, error) {
	client := filegrpc.NewFileClient(m.cli.Conn())
	return client.Createvideo(ctx, in, opts...)
}

func (m *defaultFile) Getvideo(ctx context.Context, in *VideoGetRequest, opts ...grpc.CallOption) (*VideoGetResponse, error) {
	client := filegrpc.NewFileClient(m.cli.Conn())
	return client.Getvideo(ctx, in, opts...)
}

func (m *defaultFile) Deletevideo(ctx context.Context, in *VideoDeleteRequest, opts ...grpc.CallOption) (*VideoDeleteResponse, error) {
	client := filegrpc.NewFileClient(m.cli.Conn())
	return client.Deletevideo(ctx, in, opts...)
}
