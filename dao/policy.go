package dao

import "time"

type Policy struct {
	ID        int
	Name      string    `gorm:"not null;default:'';type:varchar(32)"`
	IsDelete  int       `gorm:"not null;default:0;type:tinyint(2)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpateTime"`
}
