/**
 * @Author: jiangbo
 * @Description:
 * @File:  listen
 * @Version: 1.0.0
 * @Date: 2022/05/05 11:12 下午
 */

package listen

import (
    "context"
    "github.com/zeromicro/go-zero/core/service"
    "go_zero_t2/mq/internal/config"
    "go_zero_t2/mq/internal/svc"
)

// 返回所有消费者
func Mqs(c config.Config) []service.Service {

    svcContext := svc.NewServiceContext(c)
    ctx := context.Background()

    var services []service.Service

    // kq ：消息队列.
    // services = append(services, KqMqs(c, ctx, svcContext)...)
    // asynq ： 延迟队列、定时任务
    services = append(services, AsynqMqs(c, ctx, svcContext)...)
    // other mq ....

    return services
}
