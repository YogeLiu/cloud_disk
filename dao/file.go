package dao

import (
	"context"
	"fmt"
	"reflect"

	"github.com/YogeLiu/CloudDisk/model"

	"github.com/YogeLiu/CloudDisk/pkg/util"
	"gorm.io/gorm"
)

type FileQuery struct {
	util.BaseQuery
	ID        int64  `gorm:"id"`
	ParentID  int64  `gorm:"parent_id"`
	UserID    uint   `gorm:"user_id"`
	FuzzyName string `gorm:"name" match_type:"fuzzy"`
	MatchName string `gorm:"name" math_type:"exact"`
	IsDelete  *int   `gorm:"is_delete"`
}

func (q *FileQuery) formatQuery() (query string, values []interface{}) {
	qType := reflect.TypeOf(q)
	qValue := reflect.ValueOf(&q).Elem()
	for i := 0; i < qType.NumField(); i++ {
		key, ok := qType.Field(i).Tag.Lookup("gorm")
		if !ok {
			continue
		}
		if qValue.Field(i).IsZero() {
			continue
		}
		matchType, matchKeyOk := qType.Field(i).Tag.Lookup("match_type") // 默认精确匹配
		if !matchKeyOk || matchType == "exact" {
			values = append(values, qValue.Field(i).Interface())
			query += fmt.Sprintf("%s = ? and ", key)
		} else {
			values = append(values, fmt.Sprintf("%%%s%%", qValue.Field(i).String()))
			query += fmt.Sprintf("%s like ? and ", key)
		}
		if len(query) > 5 {
			query = query[:len(query)-5]
		}
	}
	return
}

func (q *FileQuery) QueryDB(ctx context.Context, model interface{}, value interface{}) error {
	qs, vs := q.formatQuery()
	tx := DB.WithContext(ctx).Model(model).Where(qs, vs...)

	if q.NextCur != 0 {
		tx = tx.Where("id > ?", q.NextCur).Limit(q.Limit)
	} else {
		tx = tx.Limit(q.Limit)
	}
	if len(q.OrderBy) > 0 {
		for _, val := range q.OrderBy {
			for filed, direction := range val {
				var val string
				if direction == util.DB_Order_Asc {
					val = fmt.Sprintf("%s %s", filed, "Asc")
				} else {
					val = fmt.Sprintf("%s %s", filed, "Desc")
				}
				tx = tx.Order(val)
			}
		}
	}
	return tx.Find(value).Error
}

type FileDao struct {
	*gorm.DB
}

func NewFileDao() *FileDao {
	return &FileDao{DB}
}

func (dao *FileDao) Delete(ctx context.Context, fids []int64) error {
	tx := dao.WithContext(ctx)
	return tx.Model(model.File{}).Where(`id in (?)`, fids).Update(`is_delete`, 1).Error
}

func (dao *FileDao) Updates(ctx context.Context, fids []int64, value map[string]interface{}) error {
	tx := dao.WithContext(ctx).Model(model.File{})
	return tx.Where("id in (?)", fids).UpdateColumns(value).Error
}

func (dao *FileDao) Update(ctx context.Context, fid int64, value map[string]interface{}) error {
	tx := dao.WithContext(ctx).Model(model.File{})
	return tx.Where("id = ?", fid).Updates(value).Error
}

func (dao *FileDao) Create(ctx context.Context, files []*model.File) error {
	tx := dao.WithContext(ctx)
	return tx.Create(&files).Error
}

func (dao *FileDao) Gets(ctx context.Context, query *FileQuery) (files []*model.File, err error) {
	files = make([]*model.File, 0)
	err = query.QueryDB(ctx, &model.File{}, files)
	return
}

func (dao *FileDao) Get(ctx context.Context, query *FileQuery) (file *model.File, err error) {
	file = &model.File{}
	err = query.QueryDB(ctx, &model.File{}, file)
	return
}
