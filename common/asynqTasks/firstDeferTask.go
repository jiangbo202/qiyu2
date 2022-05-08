/**
 * @Author: jiangbo
 * @Description:
 * @File:  firstDeferTask
 * @Version: 1.0.0
 * @Date: 2022/05/05 11:48 下午
 */

package asynqTasks

import (
    "encoding/json"
    "fmt"
    "github.com/hibiken/asynq"
)

const (
    FirstDeferTask = "qiyu:task:firstDeferTask"
)

// FirstDeferTaskPayload 延迟任务payload
type FirstDeferTaskPayload struct {
    Sn string
}

func NewFirstDeferTask(sn string) (*asynq.Task, error) {
    payload, err := json.Marshal(FirstDeferTaskPayload{Sn: sn})
    if err != nil {
        return nil, fmt.Errorf("NewFirstDeferTask异常: %v", err.Error())
    }
    return asynq.NewTask(FirstDeferTask, payload), nil
}
