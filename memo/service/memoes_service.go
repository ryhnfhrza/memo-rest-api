package service

import (
	"context"
	"memoAPI/model/web"
)

type MemoesService interface{
	Create(ctx context.Context,request web.MemoesCreateRequest)web.MemoesResponse
	Update(ctx context.Context,request web.MemoesUpdateRequest )web.MemoesResponse
	Delete(ctx context.Context,memoesId int)
	FindById(ctx context.Context,memoesId int)web.MemoesResponse
	FindByTitle(ctx context.Context,memoesTitle string)web.MemoesResponse
	OrderByTitleAsc(ctx context.Context)[]web.MemoesResponse
	OrderByIdDesc(ctx context.Context)[]web.MemoesResponse
	FindAll(ctx context.Context)[]web.MemoesResponse
}