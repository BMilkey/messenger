package database

import (
	"context"
	"fmt"

	"github.com/BMilkey/messenger/hlp"
	pgx "github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

// import the postgres driver

func Init(cfg hlp.DatabaseConfig) error {
	dbUrl := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/"
	dbpool, err := pgx.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to create connection pool: %v\n", err))
	}
	defer dbpool.Close()

	err = checkCreateDB(dbpool, cfg)
	if err != nil {
		return fmt.Errorf("CheckAndCreateDB failed: %v", err)

	}
	/*
		var greeting string
		err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
		if err != nil {
			return errors.New(fmt.Sprintf("QueryRow failed: %v\n", err))

		}

		fmt.Println(greeting)
	*/
	return nil
}
