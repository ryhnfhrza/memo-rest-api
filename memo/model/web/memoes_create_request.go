package web

type MemoesCreateRequest struct {
	Title    string `validate:"required,min=1,max=20"`
	MemoText string
}