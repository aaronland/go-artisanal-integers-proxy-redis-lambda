# go-artisanal-integers-proxy-redis

Go Redis-backed proxy for artisanal integer services.

## Install

You will need to have both `Go` and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Tools

### proxy-server

```
./bin/proxy-server -h
Usage of ./bin/proxy-server:
  -brooklyn-integers
	Use Brooklyn Integers as an artisanal integer source.
  -dsn string
       A valid Redis DSN string. (default "redis://localhost:6379")
  -host string
    	Host to listen on. (default "localhost")
  -httptest.serve string
    		  if non-empty, httptest.NewServer serves on this address and blocks
  -key string
       A valid Redis list key. (default "artisanalintegers")
  -loglevel string
    	    Log level. (default "info")
  -london-integers
	Use London Integers as an artisanal integer source.
  -min int
       The minimum number of artisanal integers to keep on hand at all times. (default 5)
  -mission-integers
	Use Mission Integers as an artisanal integer source.
  -port int
    	Port to listen on. (default 8080)
  -protocol string
    	    The protocol to use for the proxy server. (default "http")
```

For example, start the proxy server using Brooklyn Integers and Mission Integers as data sources:

```
./bin/proxy-server -brooklyn-integers -mission-integers
07:06:43.897017 [proxy-server][[proxy-server-redis] ] STATUS Listening for requests on http://localhost:8080
07:06:43.902664 [proxy-server][[proxy-server-redis] ] INFO time to refill the pool with 0 integers (success: 0 failed: 0): 3.619492ms (pool length is now 5)
07:06:48.902439 [proxy-server][[proxy-server-redis] ] STATUS pool length: 5
07:06:53.909229 [proxy-server][[proxy-server-redis] ] STATUS pool length: 5
```

Note the way that the Redis backend already has a pool of integers so we don't need to fetch any right away. In a separate terminal we connect to the Redis server and manually check the length of the pool and pop one off the list:

```
127.0.0.1:6379> LLEN artisanalintegers
(integer) 5
127.0.0.1:6379> LPOP artisanalintegers
"1360715845"
```

Switching back to the terminal running the proxy we can see that we've triggered the "minimum number of integers" flag and so the proxy server fetches and stores a new integer:

```
07:07:13.757728 [proxy-server][[proxy-server-redis] ] STATUS pool length: 4
07:07:15.813444 [proxy-server][[proxy-server-redis] ] INFO time to refill the pool with 1 integers (success: 1 failed: 0): 2.059029732s (pool length is now 5)
07:07:18.762247 [proxy-server][[proxy-server-redis] ] STATUS pool length: 5
```

Now we fetch an integer from the proxy server itself:

```
> curl localhost:8080
1360715849
```

Which again triggers the "fetch a new integer" signal. And so on...

```
07:07:48.392514 [proxy-server][[proxy-server-redis] ] STATUS pool length: 5
07:07:53.394646 [proxy-server][[proxy-server-redis] ] STATUS pool length: 4
07:07:56.977211 [proxy-server][[proxy-server-redis] ] INFO time to refill the pool with 1 integers (success: 1 failed: 0): 3.586287557s (pool length is now 5)
07:07:58.398863 [proxy-server][[proxy-server-redis] ] STATUS pool length: 5
```

## See also:

* https://github.com/aaronland/go-artisanal-integers-proxy
* https://github.com/aaronland/go-artisanal-integers
* https://redis.io/
