package model

import "time"

type File struct {
	ID              int64
	ParentID        int64     `gorm:"index:parent_id;not null;default:0"`
	Name            string    `gorm:"idnex:name;not null;default:'';type:varchar(255)"`
	IsFolder        int       `gorm:"is_folder;not null;default:0;type:tinyint(2)"`
	FileType        string    `gorm:"file_type;not null;default:'';type:varchar(8)"`
	UserID          uint      `gorm:"index:idx_user_id;not null;default:0;"`
	Path            string    `gorm:"->:path;not null;default:'';type:varchar(255)"`
	PhysicalAddr    string    `gorm:"not null;default:'';type:varchar(128)"`
	Size            uint64    `gorm:"not null;default:0"`
	PolicyID        uint      `gorm:"not null;default:0"`
	IsDelete        int       `gorm:"not null;default:0;type:tinyint(2)"`
	UploadSessionID string    `gorm:"unique_index:idx_upload_session_id;not null;default:'';type:varchar(128)"`
	Created         time.Time `gorm:"autoCreateTime"`
	Updated         time.Time `gorm:"autoUpdateTime"`
}
