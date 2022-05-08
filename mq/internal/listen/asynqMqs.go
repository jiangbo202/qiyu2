/**
 * @Author: jiangbo
 * @Description:
 * @File:  asynqMqs
 * @Version: 1.0.0
 * @Date: 2022/05/05 11:13 下午
 */

package listen

import (
    "context"
    "github.com/zeromicro/go-zero/core/service"
    "go_zero_t2/mq/internal/config"
    "go_zero_t2/mq/internal/mqs/cronMq"
    "go_zero_t2/mq/internal/mqs/deferMq"
    "go_zero_t2/mq/internal/svc"
)

// AsynqMqs asynq
// 定时任务、延迟任务
func AsynqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

    return []service.Service{
        // 监听延迟队列
        deferMq.NewAsynqTask(ctx, svcContext),

        // 监听定时任务 -- 服务端
        cronMq.NewAsynqCronTask(ctx, svcContext),
        // 监听定时任务 -- 客户端
        cronMq.NewAsynqCronTaskClient(ctx, svcContext),
    }

}
