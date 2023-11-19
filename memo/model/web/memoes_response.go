package web

import "time"

type MemoesResponse struct {
	Id         int `json:"id"`
	Title      string `json:"title"`
	MemoText   string	`json:"memo_text"`
	Created_At time.Time	`json:"created_at"`
	Updated_At time.Time	`json:"updated_at"`
}