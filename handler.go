package main

import (
	"context"

	"github.com/ameidance/paster_facade/model/vo"
	"github.com/ameidance/paster_facade/model/vo/kitex_gen/paster/facade"
	"github.com/ameidance/paster_facade/service"
	"github.com/ameidance/paster_facade/util"
	"github.com/cloudwego/kitex/pkg/klog"
)

// FacadeImpl implements the last service interface defined in the IDL.
type FacadeImpl struct{}

// GetPost implements the FacadeImpl interface.
func (s *FacadeImpl) GetPost(ctx context.Context, req *facade.GetPostRequest) (resp *facade.GetPostResponse, err error) {
	klog.Infof("[GetPost] req:%v", util.GetJsonString(req))
	resp = service.GetPost(ctx, vo.NewGetPostRequest(req)).GetPostResponse
	klog.Infof("[GetPost] resp:%v", util.GetJsonString(resp))
	return
}

// SavePost implements the FacadeImpl interface.
func (s *FacadeImpl) SavePost(ctx context.Context, req *facade.SavePostRequest) (resp *facade.SavePostResponse, err error) {
	klog.Infof("[SavePost] req:%v", util.GetJsonString(req))
	resp = service.SavePost(ctx, vo.NewSavePostRequest(req)).SavePostResponse
	klog.Infof("[SavePost] resp:%v", util.GetJsonString(resp))
	return
}

// GetComments implements the FacadeImpl interface.
func (s *FacadeImpl) GetComments(ctx context.Context, req *facade.GetCommentsRequest) (resp *facade.GetCommentsResponse, err error) {
	klog.Infof("[GetComments] req:%v", util.GetJsonString(req))
	resp = service.GetComments(ctx, vo.NewGetCommentsRequest(req)).GetCommentsResponse
	klog.Infof("[GetComments] resp:%v", util.GetJsonString(resp))
	return
}

// SaveComment implements the FacadeImpl interface.
func (s *FacadeImpl) SaveComment(ctx context.Context, req *facade.SaveCommentRequest) (resp *facade.SaveCommentResponse, err error) {
	klog.Infof("[SaveComment] req:%v", util.GetJsonString(req))
	resp = service.SaveComment(ctx, vo.NewSaveCommentRequest(req)).SaveCommentResponse
	klog.Infof("[SaveComment] resp:%v", util.GetJsonString(resp))
	return
}

// Check implements the FacadeImpl interface.
func (s *FacadeImpl) Check(ctx context.Context, req *facade.HealthCheckRequest) (resp *facade.HealthCheckResponse, err error) {
	return &facade.HealthCheckResponse{Status: facade.ServingStatus_SERVING}, nil
}

func (s *FacadeImpl) Watch(req *facade.HealthCheckRequest, stream facade.Facade_WatchServer) (err error) {
	return
}
