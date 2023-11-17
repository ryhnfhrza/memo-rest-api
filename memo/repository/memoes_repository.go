package repository

import (
	"context"
	"database/sql"
	"memoAPI/model/domain"
)

type MemoesRepository interface {
	Create(ctx context.Context,tx *sql.Tx,memoes domain.Memoes)domain.Memoes
	Update(ctx context.Context,tx *sql.Tx,memoes domain.Memoes)domain.Memoes
	Delete(ctx context.Context,tx *sql.Tx,memoes domain.Memoes)
	FindById(ctx context.Context,tx *sql.Tx,memoesId int)(domain.Memoes,error)
	FindByTitle(ctx context.Context,tx *sql.Tx,memoesTitle string)(domain.Memoes,error)
	OrderByTitleAsc(ctx context.Context,tx *sql.Tx)[]domain.Memoes
	OrderByIdDesc(ctx context.Context,tx *sql.Tx)[]domain.Memoes
	FindAll(ctx context.Context,tx *sql.Tx)[]domain.Memoes
}