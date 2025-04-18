package config

import (
	"database/sql"

	"github.com/gentcod/db-admin/internal/db"
	"github.com/gentcod/db-admin/internal/pkg/logger"
)

type AppConfig struct {
	*db.Queries
	conn   *sql.DB
	Logger *logger.Logger
}

var Cfg AppConfig

func NewConfig(conn *sql.DB) AppConfig {
	return AppConfig{
		conn:    conn,
		Queries: db.New(conn),
		Logger:  logger.NewLogger(),
	}
}
