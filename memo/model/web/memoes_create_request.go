package web

type MemoesCreateRequest struct {
	Title    string `validate:"required,min=1,max=100" json:"title"`
	MemoText string `json:"memo_text"`
}