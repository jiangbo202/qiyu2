package logic

import (
	"context"
	"database/sql"
	"fmt"
	errorx "go_zero_t2/common"
	"time"

	"go_zero_t2/internal/svc"
	"go_zero_t2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorUpdateLogic {
	return &AuthorUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorUpdateLogic) AuthorUpdate(req *types.AuthUpdateReq) (resp *types.AuthUpdateRes, err error) {
	// todo: add your logic here and delete this line
	findRes, err := l.svcCtx.Author.FindOne(l.ctx, int64(req.Id))
	if err != nil {
		return nil, errorx.NewDefaultError(fmt.Sprintf("将要更新的作者信息不存在author_id : %d ,err:%v", req.Id, err))
	}
	flag := 0
	if req.Author.Name != "" && req.Author.Name != findRes.Name.String {
		findRes.Name = sql.NullString{String: req.Author.Name, Valid: true}
		flag = 1
	}
	if req.Author.Age != 0 && req.Author.Age != int(findRes.Age.Int64) {
		findRes.Age = sql.NullInt64{Int64: int64(req.Author.Age), Valid: true}
		flag = 1
	}
	if req.Author.Gender != "" && req.Author.Gender != findRes.Gender.String {
		findRes.Gender = sql.NullString{String: req.Author.Gender, Valid: true}
		flag = 1
	}
	if flag == 0 {
		return nil, errorx.NewDefaultError("更新的信息无变化!")
	} else {
		findRes.UpdateAt = sql.NullTime{Time: time.Now(), Valid: true}
	}

	err = l.svcCtx.Author.Update(l.ctx, findRes)
	if err != nil {
		return nil, err
	}
	return &types.AuthUpdateRes{
		Code: 0,
		Msg: "修改成功!",
		Data: types.AuthorRes{
			Id: int(findRes.Id),
			Author: types.Author{
				Name:   findRes.Name.String,
				Age: int(findRes.Age.Int64),
				Gender: findRes.Gender.String,
			},
			CreateAt: findRes.CreateAt.Time.Format("2006-01-02 15:04:05"),
			UpdateAt: findRes.UpdateAt.Time.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
