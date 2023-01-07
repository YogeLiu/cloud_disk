package model

import "time"

type Policy struct {
	ID       int
	Name     string
	IsDelete int
	Created  time.Time
	Updated  time.Time
}
