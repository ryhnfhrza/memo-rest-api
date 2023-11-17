package web

import "time"

type MemoesResponse struct {
	Id         int
	Title      string
	MemoText   string
	Created_At time.Time
	Updated_At time.Time
}