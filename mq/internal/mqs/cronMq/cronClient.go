/**
 * @Author: jiangbo
 * @Description:
 * @File:  cronClient
 * @Version: 1.0.0
 * @Date: 2022/05/06 12:36 上午
 */

package cronMq

import (
    "context"
    "fmt"
    "github.com/hibiken/asynq"
    "go_zero_t2/common/asynqTasks"
    "go_zero_t2/mq/internal/svc"
    "log"
    "time"
)

type AsynqCronTaskClient struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewAsynqCronTaskClient(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqCronTaskClient {
    return &AsynqCronTaskClient{
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *AsynqCronTaskClient) Start() {

    fmt.Println("AsynqCronTaskClient start ")

    loc, err := time.LoadLocation("Asia/Shanghai")
    // 周期性任务
    scheduler := asynq.NewScheduler(
        asynq.RedisClientOpt{Addr: l.svcCtx.Config.Redis.Host, Password: l.svcCtx.Config.Redis.Pass},
        &asynq.SchedulerOpts{
            Location: loc,
        })

    firstTask, _ := asynqTasks.NewFirstCronTask()
    // 每隔5分钟执行一次
    entryID, err := scheduler.Register("*/1 * * * *", firstTask, asynq.Queue(l.svcCtx.Config.AsynqConf.Queque))
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("registered1 an entry: %q\n", entryID)

    if err := scheduler.Run(); err != nil {
        log.Fatal(err)
    }
}

func (l *AsynqCronTaskClient) Stop() {
    fmt.Println("AsynqCronTaskClient stop")
}


