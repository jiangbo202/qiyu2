/**
 * @Author: jiangbo
 * @Description:
 * @File:  firstCronTask
 * @Version: 1.0.0
 * @Date: 2022/05/06 12:28 上午
 */

package asynqTasks

import (
    "github.com/hibiken/asynq"
)

const (
    FirstCronTask = "qiyu:task:firstCronTask"
)

func NewFirstCronTask() (*asynq.Task, error) {
    return asynq.NewTask(FirstCronTask, []byte("abc")), nil
}

