package model

import "time"

type Group struct {
	ID       int
	Name     string
	IsDelete int
	Created  time.Time
	Updated  time.Time
}
