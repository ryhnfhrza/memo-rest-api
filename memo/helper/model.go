package helper

import (
	"memoAPI/model/domain"
	"memoAPI/model/web"
)

func ToMemoesResponse(memoes domain.Memoes) web.MemoesResponse{
	return web.MemoesResponse{
		Id: memoes.Id,
		Title: memoes.Title,
		MemoText: memoes.MemoText,
		Created_At: memoes.Created_At,
		Updated_At: memoes.Updated_At,
	}
}

func ToMemoesResponses(memo []domain.Memoes)[]web.MemoesResponse{
	var memoesResponses []web.MemoesResponse
	for _,memoes := range memo{
		memoesResponses = append(memoesResponses, ToMemoesResponse(memoes))
	}
	return memoesResponses
}