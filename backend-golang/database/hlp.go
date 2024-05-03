package database

import (
	"context"
	"fmt"
	"strings"
    "crypto/sha256"
    "encoding/hex"
	"github.com/BMilkey/messenger/hlp"
	pgx "github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"github.com/google/uuid"
)

func GetDbPool(dbConfig hlp.DatabaseConfig) (*pgx.Pool, error) {
	dbUrl := "postgres://" + dbConfig.User + ":" + dbConfig.Password + "@" + dbConfig.Host + ":" + dbConfig.Port + "/" + dbConfig.DbName
	dbpool, err := pgx.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}

func GenerateUUID() string {
    uuid := uuid.New().String()
    return uuid
}

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


func Sha256Hash(input string) string {
	inputBytes := []byte(input)    
    // Calculate the SHA-256 hash
    hash := sha256.Sum256(inputBytes)    
    // Convert the hash to a hexadecimal string
    hashString := hex.EncodeToString(hash[:])
    
	return hashString
}
