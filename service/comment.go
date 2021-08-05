package service

import (
	"context"

	"github.com/ameidance/paster_facade/client"
	"github.com/ameidance/paster_facade/constant"
	"github.com/ameidance/paster_facade/manager"
	"github.com/ameidance/paster_facade/model/vo"
	"github.com/ameidance/paster_facade/util"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetComments(ctx context.Context, req *vo.GetCommentsRequest) *vo.GetCommentsResponse {
	resp := new(vo.GetCommentsResponse)

	rpcResp, err := client.CoreClient.GetComments(ctx, req.ConvertToDTO())
	if errStatus := util.CheckRpcResponse(rpcResp, err); !util.IsStatusSuccess(errStatus) {
		klog.Errorf("[GetComments] rpc [GetComments] failed. errStatus:%v", errStatus)
		if rpcResp != nil {
			util.FillBizResp(resp, manager.ConvertToHttpStatus(&constant.ErrorStatus{
				StatusCode: rpcResp.GetStatusCode(),
				StatusMsg:  rpcResp.GetStatusMessage(),
			}))
		} else {
			util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
		}
		return resp
	}

	resp.ConvertFromDTO(rpcResp)
	return resp
}

func SaveComment(ctx context.Context, req *vo.SaveCommentRequest) *vo.SaveCommentResponse {
	resp := new(vo.SaveCommentResponse)

	if overLimit := manager.IsOverFrequencyLimit(ctx, ctx.Value("ip").(string)); overLimit {
		util.FillBizResp(resp, constant.HTTP_ERR_FREQUENCY_OVER_LIMIT)
		return resp
	}

	rpcResp, err := client.CoreClient.SaveComment(ctx, req.ConvertToDTO())
	if errStatus := util.CheckRpcResponse(rpcResp, err); !util.IsStatusSuccess(errStatus) {
		klog.Errorf("[SaveComment] rpc [SaveComment] failed. errStatus:%v", errStatus)
		if rpcResp != nil {
			util.FillBizResp(resp, manager.ConvertToHttpStatus(&constant.ErrorStatus{
				StatusCode: rpcResp.GetStatusCode(),
				StatusMsg:  rpcResp.GetStatusMessage(),
			}))
		} else {
			util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
		}
		return resp
	}

	resp.ConvertFromDTO(rpcResp)
	return resp
}
