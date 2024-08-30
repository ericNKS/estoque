package db

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	godotenv.Load("../../.env")
	_, err := Connection()
	if err != nil {
		t.Fatal(err)
	}

}

func TestRedisConnection(t *testing.T) {
	godotenv.Load("../../.env")
	rdb := RedisConnection()
	defer rdb.Db.Close()

	_, err := rdb.Db.Ping(context.Background()).Result()
	if err != nil {
		t.Fatal(err)
	}
}
