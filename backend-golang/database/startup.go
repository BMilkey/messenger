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
	
	{
		var isDb bool
		dbUrl := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/"
		dbpool, err := pgx.New(context.Background(), dbUrl)
		if err != nil {
			log.Fatal(fmt.Sprintf("Unable to create connection pool: %v\n", err))
		}
		defer dbpool.Close()
		
		isDb, err = isDbExist(dbpool, cfg)
		if err != nil {
			return fmt.Errorf("isDbExist failed: %v", err)
		}

		if isDb {
			return nil;
		}

		err = createDB(dbpool, cfg)
		if err != nil {
			return fmt.Errorf("createDB failed: %v", err)
		}
		log.Info(fmt.Sprintf("Database %s created.\n", cfg.DbName))
		dbpool.Close()
	}

	{
		dbUrl := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DbName
		dbpool, err := pgx.New(context.Background(), dbUrl)
		if err != nil {
			log.Fatal(fmt.Sprintf("Unable to create connection pool: %v\n", err))
		}
		defer dbpool.Close()

		err = createTables(dbpool, cfg)
		if err != nil {
			return fmt.Errorf("CreateTables failed: %v", err)
		}
		log.Info(fmt.Sprintf("All tables of database %s created.\n", cfg.DbName))
	}

	return nil
}
