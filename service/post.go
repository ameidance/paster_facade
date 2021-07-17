package service

import (
    "context"

    "github.com/ameidance/paster_facade/client"
    "github.com/ameidance/paster_facade/constant"
    "github.com/ameidance/paster_facade/manager"
    "github.com/ameidance/paster_facade/model/dto/kitex_gen/ameidance/paster/core"
    "github.com/ameidance/paster_facade/model/vo"
    "github.com/ameidance/paster_facade/util"
    "github.com/bytedance/gopkg/util/logger"
)

func GetPost(ctx context.Context, req *vo.GetPostRequest) *vo.GetPostResponse {
    resp := new(vo.GetPostResponse)
    postResp, err := client.CoreClient.GetPost(ctx, req.ConvertToDTO())
    if errStatus := util.CheckRpcResponse(ctx, postResp, err); !util.IsStatusSuccess(errStatus) {
        logger.CtxErrorf(ctx, "[GetPost] rpc [GetPost] failed. errStatus:%v", errStatus)
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
    if !postResp.IsSetInfo() {
        util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
        return resp
    }

    postInfo := postResp.GetInfo()
    if postInfo.GetIsDisposable() {
        commentResp, err := client.CoreClient.DeletePost(ctx, &core.DeletePostRequest{
            Id: req.Id,
        })
        if errStatus := util.CheckRpcResponse(ctx, commentResp, err); !util.IsStatusSuccess(errStatus) {
            logger.CtxErrorf(ctx, "[GetPost] rpc [DeletePost] failed. errStatus:%v", errStatus)
        }
    }

    resp.ConvertFromDTO(postResp)
    return resp
}

func SavePost(ctx context.Context, req *vo.SavePostRequest) *vo.SavePostResponse {
    resp := new(vo.SavePostResponse)

    if overLimit := manager.IsOverFrequencyLimit(ctx, ctx.Value("ip").(string)); overLimit {
        util.FillBizResp(resp, constant.HTTP_ERR_FREQUENCY_OVER_LIMIT)
        return resp
    }

    rpcResp, err := client.CoreClient.SavePost(ctx, req.ConvertToDTO())
    if errStatus := util.CheckRpcResponse(ctx, rpcResp, err); !util.IsStatusSuccess(errStatus) {
        logger.CtxErrorf(ctx, "[SavePost] rpc [SavePost] failed. errStatus:%v", errStatus)
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
