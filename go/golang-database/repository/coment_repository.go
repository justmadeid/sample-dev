package repository

import (
	"context"
	"golang-database/entity"
)

type ComentRepository interface {
	Insert(ctx context.Context, coment entity.Coment) (entity.Coment, error)
	FindById(ctx context.Context, id int32) (entity.Coment, error)
	FindAll(ctx context.Context) ([]entity.Coment, error)
}
