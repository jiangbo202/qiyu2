package logic

import (
    "context"
	"fmt"
	"github.com/hibiken/asynq"
    "time"

    "go_zero_t2/internal/svc"
    "go_zero_t2/internal/types"

    "github.com/zeromicro/go-zero/core/logx"
    "go_zero_t2/common/asynqTasks"
)

type AsynqDeferLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewAsynqDeferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqDeferLogic {
    return &AsynqDeferLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *AsynqDeferLogic) AsynqDefer(req *types.AsynqDeferTask1Req) (resp *types.AsynqDeferTask1Res, err error) {
    // todo: add your logic here and delete this line
    task, err := asynqTasks.NewFirstDeferTask(req.Payload)
    if err != nil {
        return nil, err
    }
    _, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.ProcessIn(20*time.Second), asynq.Queue("low"))
	if err != nil {
		return &types.AsynqDeferTask1Res{
			Code: 0,
			Msg: fmt.Sprintf("触发延迟任务(FirstDeferTask)失败: %s", err.Error()),
		}, fmt.Errorf("触发延迟任务(sn: %s)失败: %s", req.Payload, err.Error())
	}
    return &types.AsynqDeferTask1Res{
		Code: 0,
		Msg: "触发延迟任务(FirstDeferTask)成功",
    }, nil
}
