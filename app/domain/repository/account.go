package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	FindByID(ctx context.Context, accountID object.AccountID) (*object.Account, error)
	// TODO: Add Other APIs
	CreateUser(ctx context.Context, account *object.Account) error
}
