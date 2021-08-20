package service

import (
	"context"

	"github.com/ameidance/paster_facade/client"
	"github.com/ameidance/paster_facade/constant"
	"github.com/ameidance/paster_facade/manager"
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/core"
	"github.com/ameidance/paster_facade/model/vo"
	"github.com/ameidance/paster_facade/util"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetPost(ctx context.Context, req *vo.GetPostRequest) *vo.GetPostResponse {
	resp := vo.NewGetPostResponse()

	if !req.CheckParams() {
		util.FillBizResp(resp, constant.HTTP_ERR_WRONG_PARAMS)
		return resp
	}

	postResp, err := client.CoreClient.GetPost(ctx, req.ConvertToDTO())
	if errStatus := util.CheckRpcResponse(postResp, err); !util.IsStatusSuccess(errStatus) {
		klog.Errorf("[GetPost] rpc [GetPost] failed. errStatus:%v", util.GetJsonString(errStatus))
		if postResp != nil {
			util.FillBizResp(resp, manager.ConvertToHttpStatus(&constant.ErrorStatus{
				StatusCode: postResp.GetStatusCode(),
				StatusMsg:  postResp.GetStatusMessage(),
			}))
		} else {
			util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
		}
		return resp
	}
	if postResp.Info == nil {
		util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
		return resp
	}

	postInfo := postResp.GetInfo()
	if postInfo.GetIsDisposable() {
		commentResp, err := client.CoreClient.DeletePost(ctx, &core.DeletePostRequest{Id: req.Id})
		if errStatus := util.CheckRpcResponse(commentResp, err); !util.IsStatusSuccess(errStatus) {
			klog.Errorf("[GetPost] rpc [DeletePost] failed. errStatus:%v", util.GetJsonString(errStatus))
		}
	}

	resp.ConvertFromDTO(postResp)
	return resp
}

func SavePost(ctx context.Context, req *vo.SavePostRequest) *vo.SavePostResponse {
	resp := vo.NewSavePostResponse()

	if !req.CheckParams() {
		util.FillBizResp(resp, constant.HTTP_ERR_WRONG_PARAMS)
		return resp
	}

	if overLimit := manager.IsOverFrequencyLimit(ctx, ctx.Value("ip").(string)); overLimit {
		util.FillBizResp(resp, constant.HTTP_ERR_FREQUENCY_OVER_LIMIT)
		return resp
	}

	rpcResp, err := client.CoreClient.SavePost(ctx, req.ConvertToDTO())
	if errStatus := util.CheckRpcResponse(rpcResp, err); !util.IsStatusSuccess(errStatus) {
		klog.Errorf("[SavePost] rpc [SavePost] failed. errStatus:%v", errStatus)
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
