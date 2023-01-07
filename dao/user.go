package dao

import (
	"context"
	"time"

	"github.com/YogeLiu/CloudDisk/pkg/util"
)

type User struct {
	ID        int
	Name      string    `gorm:"uniqueIndex:idx_name;not null;default:'';varchar(11)"`
	Password  string    `gorm:"not null;default:'';type:varchar(64)"`
	Avator    string    `gorm:"not null;default:'';type:varchar(48)"`
	Email     string    `gorm:"not null;default:'';type:varchar(32)"`
	IsDelete  int       `gorm:"not null;default:0;type:tinyint(2)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpateTime"`
}

func GetUserByName(ctx context.Context, name string) (u User, err error) {
	err = DB.Model(User{}).Where("name = ? AND is_delete = ?", name, util.DB_Not_Delete_Flag).Take(&u).Error
	return
}

func GetUserByID(ctx context.Context, id int) (u User, err error) {
	u = User{}
	err = DB.Model(User{}).Where("id = ?", id).Take(&u).Error
	return
}

func DeleteUser(ctx context.Context, id int) error {
	return DB.Model(User{}).Where("id = ?", id).Update("is_delete", 1).Error
}

func AddUser(ctx context.Context, u *User) error {
	return DB.Model(u).Create(u).Error
}

func Update(ctx context.Context, id int, value map[string]interface{}) error {
	return DB.Model(User{}).Where("id = ?", id).UpdateColumns(value).Error
}
