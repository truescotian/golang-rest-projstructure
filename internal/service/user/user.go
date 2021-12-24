package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/truescotian/golang-rest-projstructure/internal"
)

type service struct {
	db *sql.DB
}

// func New(db *sql.DB) internal.Service {
func New() internal.Service {
	// return &service{db}
	return &service{}
}

func (s *service) CreateUser(ctx context.Context, in internal.CreateUserIn) (internal.CreateUserOut, error) {
	return internal.CreateUserOut{}, errors.New("unimplemented")
}
