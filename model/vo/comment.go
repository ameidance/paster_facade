package vo

import (
	"github.com/ameidance/paster_facade/constant"
	"github.com/ameidance/paster_facade/manager"
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/core"
	"github.com/ameidance/paster_facade/util"
)

type CommentInfo struct {
	Content  string `json:"content"`
	Nickname string `json:"nickname"`
	Time     int64  `json:"time"`
}

type GetCommentsRequest struct {
	PostId int64  `json:"post_id" form:"post_id"`
	Passwd string `json:"passwd,omitempty" form:"passwd,omitempty"`
}

type GetCommentsResponse struct {
	Info          []*CommentInfo `json:"info"`
	StatusCode    int32          `json:"status_code"`
	StatusMessage string         `json:"status_msg"`
}

type SaveCommentRequest struct {
	Content  string `json:"content"`
	Nickname string `json:"nickname"`
	PostId   int64  `json:"post_id"`
	Passwd   string `json:"passwd,omitempty"`
}

type SaveCommentResponse struct {
	StatusCode    int32  `json:"status_code"`
	StatusMessage string `json:"status_msg"`
}

func (m *GetCommentsRequest) ConvertToDTO() *core.GetCommentsRequest {
	if m == nil {
		return nil
	}
	return &core.GetCommentsRequest{
		PostId:   m.PostId,
		Password: m.Passwd,
	}
}

func (m *GetCommentsResponse) ConvertFromDTO(dto *core.GetCommentsResponse) {
	if dto == nil || dto.Info == nil {
		return
	}

	info := dto.GetInfo()
	m.Info = make([]*CommentInfo, 0)
	for _, each := range info {
		m.Info = append(m.Info, &CommentInfo{
			Content:  each.GetContent(),
			Nickname: each.GetNickname(),
			Time:     each.GetCreateTime(),
		})
	}

	errStatus := manager.ConvertToHttpStatus(&constant.ErrorStatus{StatusCode: dto.GetStatusCode(), StatusMsg: dto.GetStatusMessage()})
	util.FillBizResp(m, errStatus)
}

func (m *SaveCommentRequest) ConvertToDTO() *core.SaveCommentRequest {
	if m == nil {
		return nil
	}
	return &core.SaveCommentRequest{
		Info: &core.CommentInfo{
			Content:  m.Content,
			Nickname: m.Nickname,
		},
		PostId:   m.PostId,
		Password: m.Passwd,
	}
}

func (m *SaveCommentResponse) ConvertFromDTO(dto *core.SaveCommentResponse) {
	if dto == nil {
		return
	}

	errStatus := manager.ConvertToHttpStatus(&constant.ErrorStatus{StatusCode: dto.GetStatusCode(), StatusMsg: dto.GetStatusMessage()})
	util.FillBizResp(m, errStatus)
}
