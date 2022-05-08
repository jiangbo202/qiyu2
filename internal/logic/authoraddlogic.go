package logic

import (
	"context"
	"database/sql"
	"fmt"
	"go_zero_t2/internal/model"
	"go_zero_t2/internal/svc"
	"go_zero_t2/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorAddLogic {
	return &AuthorAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorAddLogic) AuthorAdd(req *types.AuthAddReq) (resp *types.AuthAddRes, err error) {
	// todo: add your logic here and delete this line
	var author = new(model.Author)
	author.Name = sql.NullString{String: req.Name, Valid: true}
	author.Age = sql.NullInt64{Int64: int64(req.Age), Valid: true}
	author.Gender = sql.NullString{String: req.Gender, Valid: true}
	author.CreateAt = sql.NullTime{Time: time.Now(), Valid: true}
	author.UpdateAt = sql.NullTime{Time: time.Now(), Valid: true}
	_, err = l.svcCtx.Author.Insert(l.ctx, author)
	if err != nil {
		return nil, err
	}
	return &types.AuthAddRes{
		Code: 0,
		Msg: fmt.Sprintf("添加%s成功!", req.Name),
		Data: types.AuthorRes{
			Id: int(author.Id),
			Author: types.Author{
				Name:   author.Name.String,
				Age: int(author.Age.Int64),
				Gender: author.Gender.String,
			},
			CreateAt: author.CreateAt.Time.Format("2006-01-02 15:04:05"),
			UpdateAt: author.UpdateAt.Time.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
