package util

import "context"

type BaseQuery struct {
	OrderBy []map[string]int
	NextCur int64
	Limit   int
}

type Queriable interface {
	FormatQuery() (string, []interface{})
	QueryDB(ctx context.Context, model interface{}, result interface{})
}
