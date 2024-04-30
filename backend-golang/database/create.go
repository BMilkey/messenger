package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/BMilkey/messenger/hlp"
	pgx "github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

func checkCreateDB(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	query := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname LIKE '%s';", cfg.DbName)
	var DbName string
	err := pool.QueryRow(context.Background(), query).Scan(&DbName)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			// Database doesn't exist, create it
			err := createDB(pool, cfg)
			if err != nil {
				return err
			}

			log.Info(fmt.Sprintf("Database %s created.\n", cfg.DbName))
			return nil
		}
		return err // Error other than "no rows in result set"
	}
	// Database already exists
	log.Info(fmt.Sprintf("Database %s already exists.\n", DbName))
	return nil
}

func createDB(pool *pgx.Pool, cfg hlp.DatabaseConfig) error {
	_, err := pool.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s;", cfg.DbName))
	if err != nil {
		return err
	}
	return nil
}
