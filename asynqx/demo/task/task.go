package task

import (
    "context"
    "encoding/json"
    "github.com/hibiken/asynq"
    "log"
)

const (
    WorkTest1  = "worker:test:test1"
    WorkTest2  = "worker:test:test2"
)

// NewMyFirstTask 打印
func NewMyFirstTask(conf map[string]string) *asynq.Task {
    payload, err := json.Marshal(conf)
    if err != nil {
        panic(err)
    }
    return asynq.NewTask(WorkTest1, payload)
}

// NewMySecondTask 打印
func NewMySecondTask(conf map[string]string) *asynq.Task {
    payload, err := json.Marshal(conf)
    if err != nil {
        panic(err)
    }
    return asynq.NewTask(WorkTest2, payload)
}

func HandleMyFirstTask(ctx context.Context, t *asynq.Task) error {

    var a map[string]string
    if err := json.Unmarshal(t.Payload(), &a); err != nil {
        return err
    }

    log.Printf("payload: %+v", a)

    log.Printf("This is my first asynq task...")
    return nil
}

func HandleMySecondTask(ctx context.Context, t *asynq.Task) error {

    var a map[string]string
    if err := json.Unmarshal(t.Payload(), &a); err != nil {
        return err
    }

    log.Printf("payload2: %+v", a)

    log.Printf("This is my second asynq task...")
    return nil
}