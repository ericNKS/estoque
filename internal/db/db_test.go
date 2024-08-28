package db

import (
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
	rdb, ctx := RedisConnection()
	defer rdb.Close()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		t.Fatal(err)
	}
}
