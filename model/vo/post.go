package vo

import (
    "github.com/ameidance/paster_facade/constant"
    "github.com/ameidance/paster_facade/manager"
    "github.com/ameidance/paster_facade/model/dto/kitex_gen/ameidance/paster/core"
    "github.com/ameidance/paster_facade/util"
    "github.com/apache/thrift/lib/go/thrift"
)

type GetPostRequest struct {
    Id     int64  `json:"id"`
    Passwd string `json:"passwd,omitempty"`
}

type GetPostResponse struct {
    Content       string `json:"content"`
    Lang          int32  `json:"lang"`
    Nickname      string `json:"nickname"`
    IsDisposable  bool   `json:"is_disposable"`
    Time          int64  `json:"time"`
    StatusCode    int32  `json:"status_code"`
    StatusMessage string `json:"status_msg"`
}

type SavePostRequest struct {
    Content      string `json:"content"`
    Lang         int32  `json:"lang"`
    Nickname     string `json:"nickname"`
    IsDisposable bool   `json:"is_disposable"`
    Passwd       string `json:"passwd,omitempty"`
}

type SavePostResponse struct {
    Id            int64  `json:"id"`
    StatusCode    int32  `json:"status_code"`
    StatusMessage string `json:"status_msg"`
}

func (m *GetPostRequest) ConvertToDTO() *core.GetPostRequest {
    if m == nil {
        return nil
    }
    return &core.GetPostRequest{
        Id:       m.Id,
        Password: thrift.StringPtr(m.Passwd),
    }
}

func (m *GetPostResponse) ConvertFromDTO(dto *core.GetPostResponse) {
    if dto == nil || !dto.IsSetInfo() {
        return
    }

    info := dto.GetInfo()
    m.Content = info.GetContent()
    m.Lang = int32(info.GetLanguage())
    m.Nickname = info.GetNickname()
    m.IsDisposable = info.GetIsDisposable()
    m.Time = info.GetCreateTime()

    errStatus := manager.ConvertToHttpStatus(&constant.ErrorStatus{StatusCode: dto.GetStatusCode(), StatusMsg: dto.GetStatusMessage()})
    util.FillBizResp(m, errStatus)
}

func (m *SavePostRequest) ConvertToDTO() *core.SavePostRequest {
    if m == nil {
        return nil
    }
    return &core.SavePostRequest{
        Info: &core.PostInfo{
            Content:      m.Content,
            Language:     core.LanguageType(m.Lang),
            Nickname:     m.Nickname,
            IsDisposable: m.IsDisposable,
        },
        Password: thrift.StringPtr(m.Passwd),
    }
}

func (m *SavePostResponse) ConvertFromDTO(dto *core.SavePostResponse) {
    if dto == nil {
        return
    }

    m.Id = dto.GetId()

    errStatus := manager.ConvertToHttpStatus(&constant.ErrorStatus{StatusCode: dto.GetStatusCode(), StatusMsg: dto.GetStatusMessage()})
    util.FillBizResp(m, errStatus)
}
