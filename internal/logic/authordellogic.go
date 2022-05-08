package logic

import (
	"context"
	"fmt"
	errorx "go_zero_t2/common"

	"go_zero_t2/internal/svc"
	"go_zero_t2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorDelLogic {
	return &AuthorDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorDelLogic) AuthorDel(req *types.AuthDelReq) (resp *types.AuthDelRes, err error) {
	// todo: add your logic here and delete this line
	findRes, err := l.svcCtx.Author.FindOne(l.ctx, int64(req.Id))
	if err != nil{
		return nil, errorx.NewDefaultError(fmt.Sprintf("将要删除的信息不存在author_id : %d ,err:%v", req.Id, err))
	}

	err = l.svcCtx.Author.Delete(l.ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &types.AuthDelRes{
		Code: 0,
		Msg: fmt.Sprintf("删除作者(%s)成功!", findRes.Name.String),
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
