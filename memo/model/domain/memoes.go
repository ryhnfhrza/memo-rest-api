package domain

import "time"

type Memoes struct {
	Id         int
	Title      string
	MemoText   string
	Created_At time.Time
	Updated_At time.Time
}