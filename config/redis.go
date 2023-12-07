package config

import (
	"context"
	"fmt"

	"github.com/redis/rueidis"
)

func NewRedisConnection(env *Env) *rueidis.Client {
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

	ping := cli.B().Ping().Message("deu bom").Build()
	if _, err = cli.Do(context.TODO(), ping).AsBool(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("✅ Redis client connected...")

	return &cli
}
