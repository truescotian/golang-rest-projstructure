package internal

import "context"

type Service interface {
	CreateUser(context.Context, CreateUserIn) (CreateUserOut, error)
}

type CreateUserIn struct {
	Name string
}

type CreateUserOut struct {
	Name string
}
