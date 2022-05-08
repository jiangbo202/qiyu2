/**
 * @Author: jiangbo
 * @Description:
 * @File:  mq
 * @Version: 1.0.0
 * @Date: 2022/05/05 11:11 下午
 */

package main

import (
    "flag"
    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/core/service"
    "go_zero_t2/mq/internal/config"
    "go_zero_t2/mq/internal/listen"
)

var configFile = flag.String("f", "etc/qiyu-mq.yaml", "Specify the config file")

func main() {
    flag.Parse()
    var c config.Config

    conf.MustLoad(*configFile, &c)

    // log、prometheus、trace、metricsUrl.
    if err := c.SetUp(); err != nil {
        panic(err)
    }

    serviceGroup := service.NewServiceGroup()
    defer serviceGroup.Stop()

    for _, mq := range listen.Mqs(c) {
        serviceGroup.Add(mq)
    }

    serviceGroup.Start()

}
