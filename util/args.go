package util

import (
	"errors"
	"github.com/aaronland/go-artisanal-integers-proxy-redis"
	"os"
	"strconv"
)

func EnsureArgs(args *redis.RedisProxyServiceArgs) error {

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

	if args.BrooklynIntegers == false {
		_, ok := os.LookupEnv("BROOKLYN_INTEGERS")
		args.BrooklynIntegers = ok
	}

	if args.MissionIntegers == false {
		_, ok := os.LookupEnv("MISSION_INTEGERS")
		args.MissionIntegers = ok
	}

	if args.LondonIntegers == false {
		_, ok := os.LookupEnv("LONDON_INTEGERS")
		args.LondonIntegers = ok
	}

	return nil
}
