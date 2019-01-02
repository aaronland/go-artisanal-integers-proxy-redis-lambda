package main

import (
	"context"
	_ "flag"
	"github.com/aaronland/go-artisanal-integers-proxy-redis"
	"github.com/aaronland/go-artisanal-integers-proxy-redis-lambda/util"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(next_int)
}

func next_int(ctx context.Context, redis_args redis.RedisProxyServiceArgs) (int64, error) {

	err := util.EnsureArgs(&redis_args)

	if err != nil {
		return -1, err
	}

	service, err := redis.NewRedisProxyService(redis_args)

	if err != nil {
		return -1, err
	}

	return service.NextInt()
}
