package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type comentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) ComentRepository {
	return &comentRepositoryImpl{DB: db}
}

func (repository *comentRepositoryImpl) Insert(ctx context.Context, coment entity.Coment) (entity.Coment, error) {
	script := "INSERT INTO coments(email, coment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, coment.Email, coment.Coment)
	if err != nil {
		return coment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return coment, err
	}
	coment.Id = int32(id)
	return coment, nil
}

func (repository *comentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Coment, error) {
	script := "SELECT id, email, coment FROM coments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	coment := entity.Coment{}
	if err != nil {
		return coment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&coment.Id, &coment.Email, &coment.Coment)
		return coment, nil
	} else {
		// tidak ada
		return coment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *comentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Coment, error) {
	script := "SELECT id, email, coment FROM coments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var coments []entity.Coment
	for rows.Next() {
		coment := entity.Coment{}
		rows.Scan(&coment.Id, &coment.Email, &coment.Coment)
		coments = append(coments, coment)
	}
	return coments, nil
}
