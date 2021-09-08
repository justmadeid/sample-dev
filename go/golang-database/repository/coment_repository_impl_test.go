package repository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestComentInsert(t *testing.T) {
	comentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	coment := entity.Coment{
		Email:  "made@repository.com",
		Coment: "oke berhasil",
	}

	r, err := comentRepository.Insert(ctx, coment)
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}

func TestFindById(t *testing.T) {
	comentRepository := NewCommentRepository(golang_database.GetConnection())

	coment, err := comentRepository.FindById(context.Background(), 14)
	if err != nil {
		panic(err)
	}
	fmt.Println(coment)
}

func TestFindAll(t *testing.T) {
	comentRepository := NewCommentRepository(golang_database.GetConnection())

	coments, err := comentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, coment := range coments {
		fmt.Println(coment)
	}

}
