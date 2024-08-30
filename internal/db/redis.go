package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ericNKS/estoque/internal/entities/types"
	"github.com/redis/go-redis/v9"
)

type DbRedis struct {
	Db  *redis.Client
	Ctx context.Context
}

func RedisConnection() *DbRedis {
	dbNum, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		dbNum = 0
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNum,
	})

	return &DbRedis{
		Db: rdb,
	}
}

func (r *DbRedis) Get(key string, tipo *types.Fornecedor) error {
	b, err := r.Db.Get(context.Background(), key).Bytes()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = r.unmarshalBinary(tipo, b)
	if err != nil {
		return err
	}

	return nil
}
func (r *DbRedis) Set(key string, tipo any, expiration time.Duration) error {
	b, err := r.marshalBinary(tipo)
	if err != nil {
		return err
	}

	r.Db.Set(context.Background(), key, b, expiration)
	if err != nil {
		return err
	}

	return nil
}

func (r *DbRedis) marshalBinary(f any) ([]byte, error) {
	return json.Marshal(f)
}

func (r *DbRedis) unmarshalBinary(f any, data []byte) error {
	return json.Unmarshal(data, f)
}
