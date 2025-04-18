package db

import "context"

type Querier interface {
	CreateUserWithPassword(ctx context.Context, arg CreateUserWithPasswordParams) error
}

var _ Querier = (*Queries)(nil)
