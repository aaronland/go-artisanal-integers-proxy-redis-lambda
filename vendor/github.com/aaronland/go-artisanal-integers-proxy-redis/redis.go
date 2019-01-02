package redis

import (
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aaronland/go-artisanal-integers-proxy"
	"github.com/whosonfirst/go-whosonfirst-log"
	pool "github.com/whosonfirst/go-whosonfirst-pool-redis"
)

type RedisProxyServiceArgs struct {
	RedisDSN         string         `json:"redis_dsn"`
	RedisKey         string         `json:"redis_key"`
	BrooklynIntegers bool           `json:"brooklyn_integers"`
	LondonIntegers   bool           `json:"london_integers"`
	MissionIntegers  bool           `json:"mission_integers"`
	MinCount         int            `json:"min_count"`
	Logger           *log.WOFLogger `json:",omitempty"`
}

func NewRedisProxyService(redis_args RedisProxyServiceArgs) (artisanalinteger.Service, error) {

	service_args := proxy.ProxyServiceArgs{
		BrooklynIntegers: redis_args.BrooklynIntegers,
		LondonIntegers:   redis_args.LondonIntegers,
		MissionIntegers:  redis_args.MissionIntegers,
		MinCount:         redis_args.MinCount,
		Logger:           redis_args.Logger,
	}

	p, err := pool.NewRedisLIFOIntPool(redis_args.RedisDSN, redis_args.RedisKey)

	if err != nil {
		return nil, err
	}

	return proxy.NewProxyServiceWithPool(p, service_args)
}
