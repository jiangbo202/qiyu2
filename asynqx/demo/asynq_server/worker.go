package main

import (
    "context"
    "github.com/hibiken/asynq"
    "go_zero_t2/asynqx/demo/task"
    "golang.org/x/sys/unix"
    "log"
    "os"
    "os/signal"
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

// loggingMiddleware 记录任务日志中间件
func loggingMiddleware(h asynq.Handler) asynq.Handler {
    return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
        start := time.Now()
        log.Printf("Start processing %q", t.Type())
        err := h.ProcessTask(ctx, t)
        if err != nil {
            return err
        }
        log.Printf("Finished processing %q: Elapsed Time = %v", t.Type(), time.Since(start))
        return nil
    })
}


func main() {

    srv := asynq.NewServer(
        RedisConnOpt,
        asynq.Config{Concurrency: 20},
    )

    mux := asynq.NewServeMux()
    mux.Use(loggingMiddleware)
    // 任务执行时的handle
    mux.HandleFunc(task.WorkTest1, task.HandleMyFirstTask)
    mux.HandleFunc(task.WorkTest2, task.HandleMySecondTask)

    // start server
    if err := srv.Start(mux); err != nil {
        log.Fatalf("could not start server: %v", err)
    }

    // Wait for termination signal.
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, unix.SIGTERM, unix.SIGINT, unix.SIGTSTP)
    for {
        s := <-sigs
        if s == unix.SIGTSTP {
            srv.Shutdown()
            continue
        }
        break
    }

    // Stop worker server.
    srv.Stop()
}
