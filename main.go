package main

import (
    "fmt"

    "github.com/ameidance/paster_facade/client"
    "github.com/ameidance/paster_facade/conf"
    "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
    client.InitRedis()

    ginConf, err := conf.GetGinConfig()
    if err != nil {
        panic(err)
    }
    err = router.Run(fmt.Sprintf("%s:%d", ginConf.Address, ginConf.Port))
    if err != nil {
        panic(err)
    }
}
