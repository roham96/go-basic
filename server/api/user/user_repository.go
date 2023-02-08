package user

import (
	"github.com/roham96/go-basic/server/database"
)

type Repository interface {
	Get() string
}

type repository struct {
	*database.DB
}

func NewRepository(db *database.DB) Repository {
	return &repository{db}
}

func (r *repository) Get() string {
	return "This is users list route ..!"
}
