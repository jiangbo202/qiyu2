/**
 * @Author: jiangbo
 * @Description:
 * @File:  firstCronTask
 * @Version: 1.0.0
 * @Date: 2022/05/06 12:31 上午
 */

package cronMq

import (
    "context"
    "github.com/hibiken/asynq"
    "github.com/zeromicro/go-zero/core/logx"
)

func (l *AsynqCronTask) firstCronTaskMqHandler(ctx context.Context, t *asynq.Task) error {
    logx.Info("我是firstCronTaskMqHandler，我在干活了！")
    return nil
}