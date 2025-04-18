package db

import (
	"context"
	"fmt"
)

const createUser = "CREATE USER %s WITH PASSWORD '%s';"

type CreateUserWithPasswordParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateUserWithPassword(ctx context.Context, arg CreateUserWithPasswordParams) error {
	query := fmt.Sprintf(createUser, arg.Username, arg.Password)
	_, err := q.db.ExecContext(ctx, query)
	return err
}
