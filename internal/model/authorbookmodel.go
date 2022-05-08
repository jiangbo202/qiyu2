package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AuthorBookModel = (*customAuthorBookModel)(nil)

type (
	// AuthorBookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuthorBookModel.
	AuthorBookModel interface {
		authorBookModel
	}

	customAuthorBookModel struct {
		*defaultAuthorBookModel
	}
)

// NewAuthorBookModel returns a model for the database table.
func NewAuthorBookModel(conn sqlx.SqlConn) AuthorBookModel {
	return &customAuthorBookModel{
		defaultAuthorBookModel: newAuthorBookModel(conn),
	}
}
