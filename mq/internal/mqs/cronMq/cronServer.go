/**
 * @Author: jiangbo
 * @Description:
 * @File:  cronServer
 * @Version: 1.0.0
 * @Date: 2022/05/06 12:26 上午
 */

package cronMq

import (
    "context"
    "fmt"
    "github.com/hibiken/asynq"
    "go_zero_t2/common/asynqTasks"
    "go_zero_t2/mq/internal/svc"
    "log"
)

type AsynqCronTask struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewAsynqCronTask(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqCronTask {
    return &AsynqCronTask{
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *AsynqCronTask) Start() {

    fmt.Println("AsynqCronTask start ")

    srv := asynq.NewServer(
        asynq.RedisClientOpt{Addr: l.svcCtx.Config.Redis.Host, Password: l.svcCtx.Config.Redis.Pass},
        asynq.Config{
            Concurrency: 10,
            Queues: map[string]int{
                l.svcCtx.Config.AsynqConf.Queque: 6,
            },
        },
    )

    mux := asynq.NewServeMux()

    // 任务注册
    mux.HandleFunc(asynqTasks.FirstCronTask, l.firstCronTaskMqHandler)

    if err := srv.Run(mux); err != nil {
        log.Fatalf("could not run server: %v", err)
    }
}

func (l *AsynqCronTask) Stop() {
    fmt.Println("AsynqCronTask stop")
}

