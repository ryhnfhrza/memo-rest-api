package web

type MemoesUpdateRequest struct {
	Id       int    `validate:"required,number"`
	Title    string `validate:"required,min=1,max=20"`
	MemoText string
}