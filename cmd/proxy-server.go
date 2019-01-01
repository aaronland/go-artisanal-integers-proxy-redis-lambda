package main

import (
	"flag"
	"github.com/aaronland/go-artisanal-integers-proxy"
	"github.com/whosonfirst/go-whosonfirst-pool-redis"
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

	flag.Parse()

	// this is the place to abstract out the run-as-lambda or invoke-as-lambda stuff...
	// (20190101/thisisaaronland)

	lambda.Start(next_int)
}

func next_int(cxt context.Context, redis_args RedisProxyServiceArgs) (*proxy.ProxyServiceResponse, error) {

	err := ensure_args(redis_args)

	if err != nil {
		return nil, err
	}

	pool, err := redis.NewRedisLIFOPool(redis_args.RedisDSN, redis_args.RedisKey)

	if err != nil {
		return nil, err
	}

	scv_func, err := proxy.NewProxyServiceLambdaFunc(pool)

	if err != nil {
		return nil, err
	}

	svc_args := proxy.ProxyServiceArgs{
		BrooklynIntegers: args.BrooklynIntegers,
		LondonIntegers:   args.LondonIntegers,
		MissionIntegers:  args.MissionIntegers,
		MinCount:         args.MinCount,
	}

	return svc_func(ctx, svc_args)
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

	// Get other args here... (20190101/thisisaaronland)

	return nil
}
