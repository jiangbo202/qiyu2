package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go_zero_t2/internal/svc"
	"go_zero_t2/internal/types"
)

type AuthorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorListLogic {
	return &AuthorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorListLogic) AuthorList(req *types.AuthListReq) (resp *types.AuthListRes, err error) {
	// todo: add your logic here and delete this line

	fmt.Printf("%+v\n", *req)
	var condition = make(map[string]interface{})
	if req.Name != "" {
		condition["name"] = req.Name
	}
	if req.Age != 0 {
		condition["age"] = req.Age
	}
	if req.Gender != "" {
		condition["gender"] = req.Gender
	}
	if req.Page != 0 {
		condition["page"] = req.Page
	} else {
		condition["page"] = 1
	}
	if req.PageSize != 0 {
		condition["page_size"] = req.PageSize
	} else {
		condition["page_size"] = 10
	}
	authorList, err := l.svcCtx.Author.FindListByCondition(l.ctx, condition)
	if err != nil {
		return nil, err
	}

	var list []types.AuthorRes
	for _, author := range authorList {
		var tmp types.AuthorRes
		tmp.Id = int(author.Id)
		tmp.Name = author.Name.String
		tmp.Age = int(author.Age.Int64)
		tmp.Gender = author.Gender.String
		tmp.CreateAt = author.CreateAt.Time.Format("2006-01-02 15:04:05")
		tmp.UpdateAt = author.UpdateAt.Time.Format("2006-01-02 15:04:05")
		list = append(list, tmp)
	}

	total, err := l.svcCtx.Author.FindListByConditionCount(l.ctx, condition)

	return &types.AuthListRes{
		Code: 0,
		Msg: "查询成功",
		Data: types.ListPaging{
			List: list,
			Page: condition["page"].(int),
			PageSize: condition["page_size"].(int),
			Total: total,
		},
	}, nil
}
