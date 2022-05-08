/**
 * @Author: jiangbo
 * @Description:
 * @File:  asynqTask
 * @Version: 1.0.0
 * @Date: 2022/05/05 11:14 下午
 */

package deferMq

import (
    "context"
    "fmt"
    "github.com/hibiken/asynq"
    "go_zero_t2/common/asynqTasks"
    "go_zero_t2/mq/internal/svc"
    "log"
)

type AsynqTask struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewAsynqTask(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqTask {
    return &AsynqTask{
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *AsynqTask) Start() {

    fmt.Println("AsynqTask start ")

    srv := asynq.NewServer(
        asynq.RedisClientOpt{Addr: l.svcCtx.Config.Redis.Host, Password: l.svcCtx.Config.Redis.Pass},
        asynq.Config{
            Concurrency: 10,
            Queues: map[string]int{
                "critical": 6,
                "default":  3,
                "low":      1,
            },
        },
    )

    mux := asynq.NewServeMux()

    // 任务注册
    mux.HandleFunc(asynqTasks.FirstDeferTask, l.firstDeferTaskMqHandler)

    if err := srv.Run(mux); err != nil {
        log.Fatalf("could not run server: %v", err)
    }
}

func (l *AsynqTask) Stop() {
    fmt.Println("AsynqTask stop")
}
