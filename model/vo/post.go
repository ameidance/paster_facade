package vo

import (
	"github.com/ameidance/paster_facade/constant"
	"github.com/ameidance/paster_facade/manager"
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/paster/core"
	"github.com/ameidance/paster_facade/model/vo/kitex_gen/paster/facade"
	"github.com/ameidance/paster_facade/util"
)

type GetPostRequest struct {
	*facade.GetPostRequest
}

type GetPostResponse struct {
	*facade.GetPostResponse
}

type SavePostRequest struct {
	*facade.SavePostRequest
}

type SavePostResponse struct {
	*facade.SavePostResponse
}

func NewGetPostRequest(req *facade.GetPostRequest) *GetPostRequest {
	vo := new(GetPostRequest)
	vo.GetPostRequest = req
	return vo
}

func NewGetPostResponse() *GetPostResponse {
	vo := new(GetPostResponse)
	vo.GetPostResponse = new(facade.GetPostResponse)
	return vo
}

func NewSavePostRequest(req *facade.SavePostRequest) *SavePostRequest {
	vo := new(SavePostRequest)
	vo.SavePostRequest = req
	return vo
}

func NewSavePostResponse() *SavePostResponse {
	vo := new(SavePostResponse)
	vo.SavePostResponse = new(facade.SavePostResponse)
	return vo
}

func (vo *GetPostRequest) ConvertToDTO() *core.GetPostRequest {
	if vo == nil {
		return nil
	}
	return &core.GetPostRequest{
		Id:       vo.Id,
		Password: vo.Passwd,
	}
}

func (vo *GetPostRequest) CheckParams() bool {
	if vo == nil {
		return false
	}
	return vo.GetId() > 0
}

func (vo *GetPostResponse) ConvertFromDTO(dto *core.GetPostResponse) {
	if dto == nil || dto.Info == nil {
		return
	}

	info := dto.GetInfo()
	vo.Content = info.GetContent()
	vo.Lang = facade.LanguageType(info.GetLanguage())
	vo.Nickname = info.GetNickname()
	vo.IsDisposable = info.GetIsDisposable()
	vo.Time = info.GetCreateTime()

	errStatus := manager.ConvertToHttpStatus(&constant.ErrorStatus{StatusCode: dto.GetStatusCode(), StatusMsg: dto.GetStatusMessage()})
	util.FillBizResp(vo, errStatus)
}

func (vo *SavePostRequest) ConvertToDTO() *core.SavePostRequest {
	if vo == nil {
		return nil
	}
	return &core.SavePostRequest{
		Info: &core.PostInfo{
			Content:      vo.Content,
			Language:     core.LanguageType(vo.Lang),
			Nickname:     vo.Nickname,
			IsDisposable: vo.IsDisposable,
		},
		Password: vo.Passwd,
	}
}

func (vo *SavePostRequest) CheckParams() bool {
	if vo == nil {
		return false
	}
	return len(vo.GetContent()) > 0 && len(vo.GetNickname()) > 0
}

func (vo *SavePostResponse) ConvertFromDTO(dto *core.SavePostResponse) {
	if dto == nil {
		return
	}

	vo.Id = dto.GetId()

	errStatus := manager.ConvertToHttpStatus(&constant.ErrorStatus{StatusCode: dto.GetStatusCode(), StatusMsg: dto.GetStatusMessage()})
	util.FillBizResp(vo, errStatus)
}
