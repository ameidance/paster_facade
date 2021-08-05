package main

import (
	"context"
	"net/http"

	"github.com/ameidance/paster_facade/constant"
	"github.com/ameidance/paster_facade/model/vo"
	"github.com/ameidance/paster_facade/service"
	"github.com/ameidance/paster_facade/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
)

func init() {
	router = gin.Default()
	router.Use(Cors())
	router.POST("/post/get", GetPost)
	router.POST("/post/save", SavePost)
	router.GET("/comment/get", GetComment)
	router.POST("/comment/save", SaveComment)
}

func GetPost(requests *gin.Context) {
	req := new(vo.GetPostRequest)
	resp := new(vo.GetPostResponse)
	if err := requests.ShouldBindJSON(&req); err != nil {
		klog.Errorf("[GetPost] bind json failed. err:%v", err)
		util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
		requests.JSON(http.StatusBadRequest, util.GetJsonMapFromStruct(resp))
		return
	}

	resp = service.GetPost(context.Background(), req)
	requests.JSON(http.StatusOK, util.GetJsonMapFromStruct(resp))
}

func SavePost(requests *gin.Context) {
	req := new(vo.SavePostRequest)
	resp := new(vo.SavePostResponse)
	if err := requests.ShouldBindJSON(&req); err != nil {
		klog.Errorf("[SavePost] bind json failed. err:%v", err)
		util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
		requests.JSON(http.StatusBadRequest, util.GetJsonMapFromStruct(resp))
		return
	}

	ctx := context.WithValue(context.Background(), "ip", requests.ClientIP())
	resp = service.SavePost(ctx, req)
	requests.JSON(http.StatusOK, util.GetJsonMapFromStruct(resp))
}

func GetComment(requests *gin.Context) {
	req := new(vo.GetCommentsRequest)
	resp := new(vo.GetCommentsResponse)
	if err := requests.ShouldBindQuery(&req); err != nil {
		klog.Errorf("[GetComment] bind json failed. err:%v", err)
		util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
		requests.JSON(http.StatusBadRequest, util.GetJsonMapFromStruct(resp))
		return
	}

	resp = service.GetComments(context.Background(), req)
	requests.JSON(http.StatusOK, util.GetJsonMapFromStruct(resp))
}

func SaveComment(requests *gin.Context) {
	req := new(vo.SaveCommentRequest)
	resp := new(vo.SaveCommentResponse)
	if err := requests.ShouldBindJSON(&req); err != nil {
		klog.Errorf("[SaveComment] bind json failed. err:%v", err)
		util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
		requests.JSON(http.StatusBadRequest, util.GetJsonMapFromStruct(resp))
		return
	}

	ctx := context.WithValue(context.Background(), "ip", requests.ClientIP())
	resp = service.SaveComment(ctx, req)
	requests.JSON(http.StatusOK, util.GetJsonMapFromStruct(resp))
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
