/**
 * @Author: jiangbo
 * @Description:
 * @File:  serviceContext
 * @Version: 1.0.0
 * @Date: 2022/05/05 11:28 下午
 */

package svc

import (
    "go_zero_t2/mq/internal/config"
)

type ServiceContext struct {
    Config config.Config

}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config: c,
    }
}

