package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-artisanal-integers-lambda/server"
	"github.com/aaronland/go-artisanal-integers-proxy-redis"
	"github.com/aaronland/go-artisanal-integers-proxy-redis-lambda/util"
	"github.com/whosonfirst/go-whosonfirst-log"
	"io"
	"net/url"
	"os"
)

func main() {

	var host = flag.String("host", "localhost", "Host to listen on.")
	var port = flag.Int("port", 8080, "Port to listen on.")
	var min = flag.Int("min", 5, "The minimum number of artisanal integers to keep on hand at all times.")
	var loglevel = flag.String("loglevel", "info", "Log level.")

	var brooklyn_integers = flag.Bool("brooklyn-integers", false, "Use Brooklyn Integers as an artisanal integer source.")
	var london_integers = flag.Bool("london-integers", false, "Use London Integers as an artisanal integer source.")
	var mission_integers = flag.Bool("mission-integers", false, "Use Mission Integers as an artisanal integer source.")

	var dsn = flag.String("dsn", "redis://localhost:6379", "A valid Redis DSN string.")
	var key = flag.String("key", "artisanalintegers", "A valid Redis list key.")

	flag.Parse()

	writer := io.MultiWriter(os.Stdout)

	logger := log.NewWOFLogger("[proxy-server-redis]")
	logger.AddLogger(writer, *loglevel)

	service_args := redis.RedisProxyServiceArgs{
		RedisDSN:         *dsn,
		RedisKey:         *key,
		BrooklynIntegers: *brooklyn_integers,
		MissionIntegers:  *mission_integers,
		LondonIntegers:   *london_integers,
		MinCount:         *min,
		Logger:           logger,
	}

	err := util.EnsureArgs(&service_args)

	if err != nil {
		logger.Fatal(err)
	}

	proxy_service, err := redis.NewRedisProxyService(service_args)

	if err != nil {
		logger.Fatal(err)
	}

	addr := fmt.Sprintf("http://%s:%d", *host, *port)
	u, err := url.Parse(addr)

	if err != nil {
		logger.Fatal(err)
	}

	svr, err := server.NewLambdaServer(u)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Status("Listening for requests on %s", svr.Address())

	err = svr.ListenAndServe(proxy_service)

	if err != nil {
		logger.Fatal(err)
	}

	os.Exit(0)
}
