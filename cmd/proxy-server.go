package main

import (
	"context"
	"errors"
	_ "flag"
	"github.com/aaronland/go-artisanal-integers-proxy"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/whosonfirst/go-whosonfirst-pool-redis"
	"os"
	"strconv"
)

type RedisProxyServiceArgs struct {
	RedisDSN         string `json:"redis_dsn"`
	RedisKey         string `json:"redis_key"`
	BrooklynIntegers bool   `json:"brooklyn_integers"`
	LondonIntegers   bool   `json:"london_integers"`
	MissionIntegers  bool   `json:"mission_integers"`
	MinCount         int    `json:"min_count"`
}

func main() {

	lambda.Start(next_int)
}

func next_int(ctx context.Context, redis_args RedisProxyServiceArgs) (int64, error) {

	err := ensure_args(&redis_args)

	if err != nil {
		return -1, err
	}

	service_args := proxy.ProxyServiceArgs{
		BrooklynIntegers: redis_args.BrooklynIntegers,
		LondonIntegers:   redis_args.LondonIntegers,
		MissionIntegers:  redis_args.MissionIntegers,
		MinCount:         redis_args.MinCount,
	}

	pool, err := redis.NewRedisLIFOIntPool(redis_args.RedisDSN, redis_args.RedisKey)

	if err != nil {
		return -1, err
	}

	service, err := proxy.NewProxyServiceWithPool(pool, service_args)

	if err != nil {
		return -1, err
	}

	return service.NextInt()
}

func ensure_args(args *RedisProxyServiceArgs) error {

	if args.RedisDSN == "" {

		dsn, ok := os.LookupEnv("REDIS_DSN")

		if !ok {
			return errors.New("Missing REDIS_DSN")
		}

		args.RedisDSN = dsn
	}

	if args.RedisKey == "" {

		key, ok := os.LookupEnv("REDIS_KEY")

		if !ok {
			return errors.New("Missing REDIS_KEY")
		}

		args.RedisKey = key
	}

	if args.MinCount == 0 {

		str_min, ok := os.LookupEnv("MIN_COUNT")

		if !ok {
			return errors.New("Missing MIN_COUNT")
		}

		min, err := strconv.Atoi(str_min)

		if err != nil {
			return err
		}

		args.MinCount = min
	}

	/*

	if args.BrooklynIntegers == nil {

		_, ok := os.LookupEnv("BROOKLYN_INTEGERS")
		args.BrooklynIntegers = ok
	}

	if args.MissionIntegers == nil {

		_, ok := os.LookupEnv("MISSION_INTEGERS")
		args.MissionIntegers = ok
	}

	if args.LondonIntegers == nil {

		_, ok := os.LookupEnv("LONDON_INTEGERS")
		args.LondonIntegers = ok
	}

	*/

	return nil
}
