package service

import (
	"context"
	"database/sql"
	"memoAPI/helper"
	"memoAPI/model/domain"
	"memoAPI/model/web"
	"memoAPI/repository"
)

type MemoesServiceImpl struct{
	memoesRepository repository.MemoesRepository
	Db *sql.DB
}

func(memoesService *MemoesServiceImpl)Create(ctx context.Context,request web.MemoesCreateRequest)web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	memoes := domain.Memoes{
		Title: request.Title,
		MemoText: request.MemoText,
	}

	memoes = memoesService.memoesRepository.Create(ctx,tx,memoes)

	return helper.ToMemoesResponse(memoes)
}

func(memoesService *MemoesServiceImpl)Update(ctx context.Context,request web.MemoesUpdateRequest )web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	memoes,err := memoesService.memoesRepository.FindById(ctx,tx,request.Id)
	helper.PanicIfError(err)

	memoes.Title = request.Title
	memoes.MemoText = request.MemoText

	memoes = memoesService.memoesRepository.Update(ctx,tx,memoes)
	return helper.ToMemoesResponse(memoes)
}

func(memoesService *MemoesServiceImpl)Delete(ctx context.Context,memoesId int){
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	memoes,err := memoesService.memoesRepository.FindById(ctx,tx,memoesId)
	helper.PanicIfError(err)

	memoesService.memoesRepository.Delete(ctx,tx,memoes)
}

func(memoesService *MemoesServiceImpl)FindById(ctx context.Context,memoesId int)web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	memoes,err := memoesService.memoesRepository.FindById(ctx,tx,memoesId)
	helper.PanicIfError(err)

	return helper.ToMemoesResponse(memoes)
}

func(memoesService *MemoesServiceImpl)FindByTitle(ctx context.Context,memoesTitle string)web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	memoes,err := memoesService.memoesRepository.FindByTitle(ctx,tx,memoesTitle)
	helper.PanicIfError(err)

	return helper.ToMemoesResponse(memoes)
}

func(memoesService *MemoesServiceImpl)OrderByTitleAsc(ctx context.Context)[]web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	memoes := memoesService.memoesRepository.OrderByTitleAsc(ctx,tx)

	return helper.ToMemoesResponses(memoes)
}

func(memoesService *MemoesServiceImpl)OrderByIdDesc(ctx context.Context)[]web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	memoes := memoesService.memoesRepository.OrderByIdDesc(ctx,tx)

	return helper.ToMemoesResponses(memoes)
}

func(memoesService *MemoesServiceImpl)FindAll(ctx context.Context)[]web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	helper.CommitOrRollback(tx)

	memoes := memoesService.memoesRepository.FindAll(ctx,tx)

	return helper.ToMemoesResponses(memoes)
}
