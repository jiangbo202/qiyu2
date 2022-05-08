/**
 * @Author: jiangbo
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/05/05 11:09 下午
 */

package config

import (
    "github.com/zeromicro/go-zero/core/service"
    "github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
    service.ServiceConf

    Redis redis.RedisConf

    AsynqConf struct{
        Queque string
    }

    // //kq
    // PaymentUpdateStatusConf kq.KqConf
    //
    // //rpc
    // OrderRpcConf      zrpc.RpcClientConf
    // MqueueRpcConf     zrpc.RpcClientConf
    // UsercenterRpcConf zrpc.RpcClientConf
}
