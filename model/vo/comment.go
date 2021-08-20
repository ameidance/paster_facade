package vo

import (
	"github.com/ameidance/paster_facade/constant"
	"github.com/ameidance/paster_facade/manager"
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/core"
	"github.com/ameidance/paster_facade/model/vo/kitex_gen/facade"
	"github.com/ameidance/paster_facade/util"
)

type GetCommentsRequest struct {
	*facade.GetCommentsRequest
}

type GetCommentsResponse struct {
	*facade.GetCommentsResponse
}

type SaveCommentRequest struct {
	*facade.SaveCommentRequest
}

type SaveCommentResponse struct {
	*facade.SaveCommentResponse
}

func NewGetCommentsRequest() *GetCommentsRequest {
	vo := new(GetCommentsRequest)
	vo.GetCommentsRequest = new(facade.GetCommentsRequest)
	return vo
}

func NewGetCommentsResponse() *GetCommentsResponse {
	vo := new(GetCommentsResponse)
	vo.GetCommentsResponse = new(facade.GetCommentsResponse)
	return vo
}

func NewSaveCommentRequest() *SaveCommentRequest {
	vo := new(SaveCommentRequest)
	vo.SaveCommentRequest = new(facade.SaveCommentRequest)
	return vo
}

func NewSaveCommentResponse() *SaveCommentResponse {
	vo := new(SaveCommentResponse)
	vo.SaveCommentResponse = new(facade.SaveCommentResponse)
	return vo
}

func (vo *GetCommentsRequest) ConvertToDTO() *core.GetCommentsRequest {
	if vo == nil {
		return nil
	}
	return &core.GetCommentsRequest{
		PostId:   vo.PostId,
		Password: vo.Passwd,
	}
}

func (vo *GetCommentsRequest) CheckParams() bool {
	if vo == nil {
		return false
	}
	return vo.GetPostId() > 0
}

func (vo *GetCommentsResponse) ConvertFromDTO(dto *core.GetCommentsResponse) {
	if dto == nil || dto.Info == nil {
		return
	}

	info := dto.GetInfo()
	vo.Info = make([]*facade.CommentInfo, 0)
	for _, each := range info {
		vo.Info = append(vo.Info, &facade.CommentInfo{
			Content:  each.GetContent(),
			Nickname: each.GetNickname(),
			Time:     each.GetCreateTime(),
		})
	}

	errStatus := manager.ConvertToHttpStatus(&constant.ErrorStatus{StatusCode: dto.GetStatusCode(), StatusMsg: dto.GetStatusMessage()})
	util.FillBizResp(vo, errStatus)
}

func (vo *SaveCommentRequest) ConvertToDTO() *core.SaveCommentRequest {
	if vo == nil {
		return nil
	}
	return &core.SaveCommentRequest{
		Info: &core.CommentInfo{
			Content:  vo.Content,
			Nickname: vo.Nickname,
		},
		PostId:   vo.PostId,
		Password: vo.Passwd,
	}
}

func (vo *SaveCommentRequest) CheckParams() bool {
	if vo == nil {
		return false
	}
	return vo.GetPostId() > 0 && len(vo.GetContent()) > 0 && len(vo.GetNickname()) > 0
}

func (vo *SaveCommentResponse) ConvertFromDTO(dto *core.SaveCommentResponse) {
	if dto == nil {
		return
	}

	errStatus := manager.ConvertToHttpStatus(&constant.ErrorStatus{StatusCode: dto.GetStatusCode(), StatusMsg: dto.GetStatusMessage()})
	util.FillBizResp(vo, errStatus)
}
