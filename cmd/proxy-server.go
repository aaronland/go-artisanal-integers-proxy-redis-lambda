package main

import (
       "github.com/aaronland/go-artisanal-integers-proxy"
       "github.com/whosonfirst/go-whosonfirst-pool-redis"
)

type RedisProxyServiceArgs struct {
	DSN              string `json:"dsn"`
	Key              string `json:"key"`
	BrooklynIntegers bool   `json:"brooklyn_integers"`
	LondonIntegers   bool   `json:"london_integers"`
	MissionIntegers  bool   `json:"mission_integers"`
	MinCount         int    `json:"min_count"`
}

func next_int(cxt context.Context, args RedisProxyServiceArgs) (*proxy.ProxyServiceResponse, error) {

	pl, err := redis.NewRedisLIFOPool(args.DSN, args.Key)

	if err != nil {
		return nil, err
	}

	scv_func, err := proxy.NewProxyServiceLambdaFunc(pl)

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

func main() {

	flag.Parse()

	lambda.Start(next_int)
}
