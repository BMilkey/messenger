package database

import (
	"fmt"
	"testing"
	
	"github.com/BMilkey/messenger/hlp"
	pgx "github.com/jackc/pgx/v5/pgxpool"
)

func TestSingleArgSelects(t *testing.T) {
	appConfig, err := hlp.GetConfig("test_config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	dbConfig := appConfig.Db

	err = Init(dbConfig)
	if err != nil {
		t.Fatal(err)
	}
	dbpool, err := hlp.GetDbPool(dbConfig)

	if err != nil {
		t.Fatal(err)		
	}

	tests := []func(*pgx.Pool) error{
		userById,
		usersByName,
		userByLink,
		chatById,
		chatByLink,
		chatsByTitle,
		msgById,
		msgsByChatId,
		msgsByText,
	}

	for _, test := range tests {
		if err := test(dbpool); err != nil {
			t.Fatal(err)
		}
	}
}

func usersByName(dbpool *pgx.Pool) error {
	ans, err := SelectUsersByName(dbpool, "test_name")
	fmt.Println(ans)
	return err
}

func userByLink(dbpool *pgx.Pool) error {
	ans, err := SelectUserByLink(dbpool, "@test_link")
	fmt.Println(ans)
	return err
}

func userById(dbpool *pgx.Pool) error {
	ans, err := SelectUserById(dbpool, "kHBrjINqoIRPuG3ACxf5XFtQdhj1")
	fmt.Println(ans)
	return err
}

func chatById(dbpool *pgx.Pool) error {
	ans, err := SelectChatById(dbpool, "test_id")
	fmt.Println(ans)
	return err
}

func chatByLink(pool *pgx.Pool) error {
	ans, err := SelectChatByLink(pool, "@test_chat_link")
	fmt.Println(ans)
	return err
}

func chatsByTitle(pool *pgx.Pool) error {
	ans, err := SelectChatsByTitle(pool, "test_chat_title")
	fmt.Println(ans)
	return err
}

func msgById(pool *pgx.Pool) error {
	ans, err := SelectMessageById(pool, "1")
	fmt.Println(ans)
	return err
}

func msgsByChatId(pool *pgx.Pool) error {
	ans, err := SelectMessagesByChatId(pool, "test_id")
	fmt.Println(ans)
	return err
}

func msgsByText(pool *pgx.Pool) error {
	ans, err := SelectMessagesByText(pool, "БУД")
	fmt.Println(ans)
	return err
}

func TestSelectMessagesByChatAndText(t *testing.T) {
	config, err := hlp.GetConfig("test_config.yaml")
	if err != nil {
		t.Fatal(err)
	}	
	dbConfig := config.Db

	err = Init(dbConfig)
	if err != nil {
		t.Fatal(err)
	}

	dbPool, err := hlp.GetDbPool(config.Db)
	if err != nil {
		t.Fatal(err)
	}
	defer dbPool.Close()

	messages, err := SelectMessagesByChatAndText(dbPool, "test_id", "БУД")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(messages)
}
