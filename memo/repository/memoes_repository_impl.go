package repository

import (
	"context"
	"database/sql"
	"errors"
	"memoAPI/helper"
	"memoAPI/model/domain"
)

type MemoesRepositoryImpl struct{

}

func(memoesRepository *MemoesRepositoryImpl)Create(ctx context.Context,tx *sql.Tx,memoes domain.Memoes)domain.Memoes{
	SQL := "insert into memoes (title,memoText) values (?,?)"
	result ,err := tx.ExecContext(ctx,SQL,memoes.Title,memoes.MemoText)
	helper.PanicIfError(err)

	id,err := result.LastInsertId()
	helper.PanicIfError(err)

	memoes.Id = int(id)

	return memoes
}

func(memoesRepository *MemoesRepositoryImpl)Update(ctx context.Context,tx *sql.Tx,memoes domain.Memoes)domain.Memoes{
	SQL := "update memoes set title = ? ,memoText = ? where id = ?"
	_ ,err := tx.ExecContext(ctx,SQL,memoes.Title,memoes.MemoText,memoes.Id)
	helper.PanicIfError(err)

	return memoes
}

func(memoesRepository *MemoesRepositoryImpl)Delete(ctx context.Context,tx *sql.Tx,memoes domain.Memoes){
	SQL := "delete from memoes where id = ?"
	_ ,err := tx.ExecContext(ctx,SQL,memoes.Id)
	helper.PanicIfError(err)
}

func(memoesRepository *MemoesRepositoryImpl)FindById(ctx context.Context,tx *sql.Tx,memoesId int)(domain.Memoes,error){
	SQL := "select id,title,memoText,created_at,updated_at from memoes where id = ?"
	rows,err := tx.QueryContext(ctx,SQL,memoesId)
	helper.PanicIfError(err)

	memoes := domain.Memoes{}
	if rows.Next(){
		err := rows.Scan(&memoes.Id,&memoes.Title,&memoes.MemoText,&memoes.Created_At,&memoes.Updated_At)
		helper.PanicIfError(err)
		return memoes,nil

	}else{
		return memoes,errors.New("Id tidak ditemukan")
	}
}

func(memoesRepository *MemoesRepositoryImpl)FindByTitle(ctx context.Context,tx *sql.Tx,memoesTitle string)(domain.Memoes,error){
	SQL := "select id,title,memoText,created_at,updated_at from memoes where title = ?"
	rows,err := tx.QueryContext(ctx,SQL,memoesTitle)
	helper.PanicIfError(err)

	memoes := domain.Memoes{}
	if rows.Next(){
		err := rows.Scan(&memoes.Id,&memoes.Title,&memoes.MemoText,&memoes.Created_At,&memoes.Updated_At)
		helper.PanicIfError(err)
		return memoes,nil

	}else{
		return memoes,errors.New("Title tidak ditemukan")
	}
}

func(memoesRepository *MemoesRepositoryImpl)OrderByTitleAsc(ctx context.Context,tx *sql.Tx)[]domain.Memoes{
	SQL := "select id,title,memoText,created_at,updated_at from memoes order by title asc"
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var memoes []domain.Memoes
	for rows.Next(){
		memo := domain.Memoes{}
		err := rows.Scan(&memo.Id,&memo.Title,&memo.MemoText,&memo.Created_At,&memo.Updated_At)
		helper.PanicIfError(err)
		memoes = append(memoes, memo)
	}

	return memoes
}

func(memoesRepository *MemoesRepositoryImpl)OrderByIdDesc(ctx context.Context,tx *sql.Tx)[]domain.Memoes{
	SQL := "select id,title,memoText,created_at,updated_at from memoes order by id desc"
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var memoes []domain.Memoes
	for rows.Next(){
		memo := domain.Memoes{}
		err := rows.Scan(&memo.Id,&memo.Title,&memo.MemoText,&memo.Created_At,&memo.Updated_At)
		helper.PanicIfError(err)
		memoes = append(memoes, memo)
	}

	return memoes
}

func(memoesRepository *MemoesRepositoryImpl)FindAll(ctx context.Context,tx *sql.Tx)[]domain.Memoes{
	SQL := "select id,title,memoText,created_at,updated_at from memoes "
	rows,err := tx.QueryContext(ctx,SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var memoes []domain.Memoes
	for rows.Next(){
		memo := domain.Memoes{}
		err := rows.Scan(&memo.Id,&memo.Title,&memo.MemoText,&memo.Created_At,&memo.Updated_At)
		helper.PanicIfError(err)
		memoes = append(memoes, memo)
	}

	return memoes
}
