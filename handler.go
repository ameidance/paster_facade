package main

import (
	"context"

	"github.com/ameidance/paster_facade/model/vo"
	"github.com/ameidance/paster_facade/model/vo/kitex_gen/facade"
	"github.com/ameidance/paster_facade/service"
)

// PasterFacadeImpl implements the last service interface defined in the IDL.
type PasterFacadeImpl struct{}

// GetPost implements the PasterFacadeImpl interface.
func (s *PasterFacadeImpl) GetPost(ctx context.Context, req *facade.GetPostRequest) (resp *facade.GetPostResponse, err error) {
	return service.GetPost(ctx, &vo.GetPostRequest{GetPostRequest: req}).GetPostResponse, nil
}

// SavePost implements the PasterFacadeImpl interface.
func (s *PasterFacadeImpl) SavePost(ctx context.Context, req *facade.SavePostRequest) (resp *facade.SavePostResponse, err error) {
	return service.SavePost(ctx, &vo.SavePostRequest{SavePostRequest: req}).SavePostResponse, nil
}

// GetComments implements the PasterFacadeImpl interface.
func (s *PasterFacadeImpl) GetComments(ctx context.Context, req *facade.GetCommentsRequest) (resp *facade.GetCommentsResponse, err error) {
	return service.GetComments(ctx, &vo.GetCommentsRequest{GetCommentsRequest: req}).GetCommentsResponse, nil
}

// SaveComment implements the PasterFacadeImpl interface.
func (s *PasterFacadeImpl) SaveComment(ctx context.Context, req *facade.SaveCommentRequest) (resp *facade.SaveCommentResponse, err error) {
	return service.SaveComment(ctx, &vo.SaveCommentRequest{SaveCommentRequest: req}).SaveCommentResponse, nil
}

// Check implements the PasterFacadeImpl interface.
func (s *PasterFacadeImpl) Check(ctx context.Context, req *facade.HealthCheckRequest) (resp *facade.HealthCheckResponse, err error) {
	return &facade.HealthCheckResponse{Status: facade.ServingStatus_SERVING}, nil
}

func (s *PasterFacadeImpl) Watch(req *facade.HealthCheckRequest, stream facade.PasterFacade_WatchServer) (err error) {
	return
}
