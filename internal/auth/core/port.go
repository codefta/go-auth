package core

import (
	"context"

	"github.com/codefta/go-auth/internal/user/core"
)

type AuthStorage interface {
	Login(ctx context.Context, email string, password string) (core.User, error)
}

type UserStorage interface {
	Insert(ctx context.Context, user core.User) (core.User, error)
	Update(ctx context.Context, id string, user core.User) (core.User, error)
	GetById(ctx context.Context, id string) (core.User, error)
	GetByEmail(ctx context.Context, email string) (core.User, error)
}
