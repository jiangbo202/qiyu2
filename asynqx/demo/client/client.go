package main

import (
    "github.com/hibiken/asynq"
    "go_zero_t2/asynqx/demo/task"
    "log"
    "time"
)

var RedisConnOpt = &asynq.RedisClientOpt{
    Addr: "127.0.0.1:6379",
    // Omit if no password is required
    Password: "3hcml2gaYn&Q93hM",
    // Use a dedicated db number for asynq.
    // By default, Redis offers 16 databases (0..15)
    DB: 0,
}

func main() {
    loc, err := time.LoadLocation("Asia/Beijing")
    // 周期性任务
    scheduler := asynq.NewScheduler(
        RedisConnOpt,
        &asynq.SchedulerOpts{
            Location: loc,
        })

    handleData := map[string]string{
        "a": "aa",
        "b": "bb",
    }

    firstTask := task.NewMyFirstTask(handleData)
    SecondTask := task.NewMySecondTask(handleData)
    // 每隔5分钟执行一次
    entryID, err := scheduler.Register("*/2 * * * *", firstTask, asynq.Queue("my_queue"))
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("registered1 an entry: %q\n", entryID)
    entryID, err = scheduler.Register("*/1 * * * *", SecondTask)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("registered2 an entry: %q\n", entryID)

    if err := scheduler.Run(); err != nil {
        log.Fatal(err)
    }
}
