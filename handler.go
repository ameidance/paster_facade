package main

import (
    "context"
    "net/http"

    "github.com/ameidance/paster_facade/constant"
    "github.com/ameidance/paster_facade/model/vo"
    "github.com/ameidance/paster_facade/service"
    "github.com/ameidance/paster_facade/util"
    "github.com/bytedance/gopkg/util/logger"
    "github.com/fatih/structs"
    "github.com/gin-gonic/gin"
)

func init() {
    router = gin.Default()
    router.POST("/post/get", GetPost)
    router.POST("/post/save", SavePost)
    router.GET("/comment/get", GetComment)
    router.POST("/comment/save", SaveComment)
}

func GetPost(requests *gin.Context) {
    req := new(vo.GetPostRequest)
    resp := new(vo.GetPostResponse)
    if err := requests.ShouldBindJSON(&req); err != nil {
        logger.Errorf("[GetPost] bind json failed. err:%v", err)
        util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
        requests.JSON(http.StatusBadRequest, structs.Map(resp))
        return
    }

    resp = service.GetPost(context.Background(), req)
    requests.JSON(http.StatusOK, structs.Map(resp))
}

func SavePost(requests *gin.Context) {
    req := new(vo.SavePostRequest)
    resp := new(vo.SavePostResponse)
    if err := requests.ShouldBindJSON(&req); err != nil {
        logger.Errorf("[SavePost] bind json failed. err:%v", err)
        util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
        requests.JSON(http.StatusBadRequest, structs.Map(resp))
        return
    }

    ctx := context.WithValue(context.Background(), "ip", requests.ClientIP())
    resp = service.SavePost(ctx, req)
    requests.JSON(http.StatusOK, structs.Map(resp))
}

func GetComment(requests *gin.Context) {
    req := new(vo.GetCommentsRequest)
    resp := new(vo.GetCommentsResponse)
    if err := requests.ShouldBindQuery(&req); err != nil {
        logger.Errorf("[GetComment] bind json failed. err:%v", err)
        util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
        requests.JSON(http.StatusBadRequest, structs.Map(resp))
        return
    }

    resp = service.GetComments(context.Background(), req)
    requests.JSON(http.StatusOK, structs.Map(resp))
}

func SaveComment(requests *gin.Context) {
    req := new(vo.SaveCommentRequest)
    resp := new(vo.SaveCommentResponse)
    if err := requests.ShouldBindJSON(&req); err != nil {
        logger.Errorf("[SaveComment] bind json failed. err:%v", err)
        util.FillBizResp(resp, constant.HTTP_ERR_SERVICE_INTERNAL)
        requests.JSON(http.StatusBadRequest, structs.Map(resp))
        return
    }

    ctx := context.WithValue(context.Background(), "ip", requests.ClientIP())
    resp = service.SaveComment(ctx, req)
    requests.JSON(http.StatusOK, structs.Map(resp))
}
