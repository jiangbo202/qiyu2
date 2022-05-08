/**
 * @Author: jiangbo
 * @Description:
 * @File:  firstDeferTask
 * @Version: 1.0.0
 * @Date: 2022/05/05 11:19 下午
 */

package deferMq

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/hibiken/asynq"
    "go_zero_t2/common/asynqTasks"
)



func (l *AsynqTask) firstDeferTaskMqHandler(ctx context.Context, t *asynq.Task) error {
    // var p firstDeferTaskPayload
    var p asynqTasks.FirstDeferTaskPayload
    if err := json.Unmarshal(t.Payload(), &p); err != nil {
        return fmt.Errorf("firstDeferTaskMqHandler异常: %v", err.Error())
    }
    fmt.Printf("收到延迟任务%+v，并处理-OK", p)
    return nil
}