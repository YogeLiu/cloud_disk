package model

import "time"

type FileMoveDTO struct {
	From []string `json:"from"`
	To   string   `json:"to"`
}

type FileDeleteDTO struct {
	Fids []string `json:"fids"`
}

type FileRenameDTO struct {
	Name int64 `json:"name,string"`
	Fid  int64 `json:"fid,string"`
}

type FileCreateDTO struct {
	ParentID int64  `json:"parent_id,string"`
	Path     string `json:"path"`
}

type FileQueryDTO struct {
	Fid            int64  `json:"fid,string"`
	ParentID       int64  `json:"parent_id,string"`
	OrderBy        string `json:"order_by"`
	OrderDirection string `json:"order_direction"`
	Name           string `json:"name"`
	Limit          int    `json:"limit"`
	NextCur        string `json:"next_cur"`
}

type FileDTO struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	ParentID  string    `json:"parent_id"`
	IsForlder int       `json:"is_folder"`
	Updated   time.Time `json:"update_time"`
	NextCur   string    `json:"next_cur"`
}
