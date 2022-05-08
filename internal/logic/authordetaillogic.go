package logic

import (
	"context"

	"go_zero_t2/internal/svc"
	"go_zero_t2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorDetailLogic {
	return &AuthorDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorDetailLogic) AuthorDetail(req *types.AuthDetailReq) (resp *types.AuthDetailRes, err error) {
	// todo: add your logic here and delete this line
	findRes, err := l.svcCtx.Author.FindOne(l.ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &types.AuthDetailRes{
		Code: 0,
		Msg: "查询成功",
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
