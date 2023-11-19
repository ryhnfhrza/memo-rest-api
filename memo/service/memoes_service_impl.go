package service

import (
	"context"
	"database/sql"
	"memoAPI/exception"
	"memoAPI/helper"
	"memoAPI/model/domain"
	"memoAPI/model/web"
	"memoAPI/repository"

	"github.com/go-playground/validator/v10"
)

type MemoesServiceImpl struct{
	memoesRepository repository.MemoesRepository
	Db *sql.DB
	Validate *validator.Validate
}

func NewMemoesService(repository repository.MemoesRepository,DB *sql.DB, validate *validator.Validate)MemoesService{
	return &MemoesServiceImpl{
		memoesRepository: repository,
		Db: DB,
		Validate: validate,
	}
}

func(memoesService *MemoesServiceImpl)Create(ctx context.Context,request web.MemoesCreateRequest)web.MemoesResponse{
	err := memoesService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	memoes := domain.Memoes{
		Title: request.Title,
		MemoText: request.MemoText,
	}

	memoes = memoesService.memoesRepository.Create(ctx,tx,memoes)

	return helper.ToMemoesResponse(memoes)
}

func(memoesService *MemoesServiceImpl)Update(ctx context.Context,request web.MemoesUpdateRequest )web.MemoesResponse{
	err := memoesService.Validate.Struct(request)
	helper.PanicIfError(err)

	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	memoes,err := memoesService.memoesRepository.FindById(ctx,tx,request.Id)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	memoes.Title = request.Title
	memoes.MemoText = request.MemoText

	memoes = memoesService.memoesRepository.Update(ctx,tx,memoes)
	return helper.ToMemoesResponse(memoes)
}

func(memoesService *MemoesServiceImpl)Delete(ctx context.Context,memoesId int){
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	memoes,err := memoesService.memoesRepository.FindById(ctx,tx,memoesId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	memoesService.memoesRepository.Delete(ctx,tx,memoes)
}

func(memoesService *MemoesServiceImpl)FindById(ctx context.Context,memoesId int)web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	memoes,err := memoesService.memoesRepository.FindById(ctx,tx,memoesId)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMemoesResponse(memoes)
}

func(memoesService *MemoesServiceImpl)FindByTitle(ctx context.Context,memoesTitle string)web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	memoesTitle = "%" + memoesTitle + "%"
	memoes,err := memoesService.memoesRepository.FindByTitle(ctx,tx,memoesTitle)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToMemoesResponse(memoes)
}

func(memoesService *MemoesServiceImpl)OrderByTitleAsc(ctx context.Context)[]web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	memoes := memoesService.memoesRepository.OrderByTitleAsc(ctx,tx)

	return helper.ToMemoesResponses(memoes)
}

func(memoesService *MemoesServiceImpl)OrderByIdDesc(ctx context.Context)[]web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	memoes := memoesService.memoesRepository.OrderByIdDesc(ctx,tx)

	return helper.ToMemoesResponses(memoes)
}

func(memoesService *MemoesServiceImpl)FindAll(ctx context.Context)[]web.MemoesResponse{
	tx,err := memoesService.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	memoes := memoesService.memoesRepository.FindAll(ctx,tx)

	return helper.ToMemoesResponses(memoes)
}
