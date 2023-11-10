package config

import (
	"fmt"

	"github.com/redis/rueidis"
)

func NewRedisConnection(env Env) *rueidis.Client {
	rdsAddr := fmt.Sprintf("%s:%s", env.RedisHost, env.RedisPort)

	cli, err := rueidis.NewClient(
		rueidis.ClientOption{
			InitAddress: []string{rdsAddr},
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer cli.Close()

	return &cli
}
