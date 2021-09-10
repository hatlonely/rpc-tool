package service

import (
	"context"

	"github.com/hatlonely/rpc-tool/api/gen/go/api"
)

type Options struct{}

type ToolService struct {
	api.UnsafeToolServiceServer
	options *Options
}

func NewToolServiceWithOptions(options *Options) (*ToolService, error) {
	return &ToolService{
		options: options,
	}, nil
}

func (s *ToolService) Ping(ctx context.Context, req *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}
