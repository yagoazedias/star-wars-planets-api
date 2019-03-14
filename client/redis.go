package client

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/yagoazedias/star-wars-planets-api/environment"
	"os"
	"strconv"
)

type Redis struct {
	db *redis.Client
}

func NewRedisClient() Redis {

	var c Redis

	host,     envHost     := os.LookupEnv("REDIS_HOST")
	port,     envPort     := os.LookupEnv("REDIS_PORT")
	database, envPassword := os.LookupEnv("REDIS_NAME")
	password, envDatabase := os.LookupEnv("REDIS_PASSWORD")


	nDatabase, err := strconv.Atoi(database)

	if !envDatabase || !envHost || !envPassword || !envPort || err != nil {
		c.db = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", environment.REDIS_HOST, environment.REDIS_PORT),
			Password: environment.REDIS_PASSWORD, // no password set
			DB:       environment.REDIS_DATABASE,  // use default DB
		})
	} else {
		c.db = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password, // no password set
			DB: nDatabase,
		})
	}

	return c
}

var Cache = NewRedisClient()