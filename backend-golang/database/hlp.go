package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/BMilkey/messenger/hlp"
	pgx "github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

func wrapRawExecQuerry(pool *pgx.Pool, querry string) error {
	_, err := pool.Exec(context.Background(), querry)

	if err != nil {
		return err
	}
	return nil
}

func isDbExist(pool *pgx.Pool, cfg hlp.DatabaseConfig) (bool, error) {
	query := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname LIKE '%s';", cfg.DbName)
	var DbName string
	err := pool.QueryRow(context.Background(), query).Scan(&DbName)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			log.Info(fmt.Sprintf("Database %s is not existing.\n", cfg.DbName))
			return false, nil
		}
		return false, err // Error other than "no rows in result set"
	}
	// Database already exists
	log.Info(fmt.Sprintf("Database %s already exists.\n", DbName))
	return true, nil
}
