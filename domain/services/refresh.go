package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/redis/rueidis"
)

var ctx = context.Background()

type RefreshService struct {
	cache rueidis.Client
	env   *config.Env
}

func NewRefreshService(cache *rueidis.Client, env *config.Env) *RefreshService {
	return &RefreshService{
		cache: *cache,
		env:   env,
	}
}

func (rs *RefreshService) GetRefreshToken(userID string) (string, error) {
	var refreshTkn string
	getCmd := rs.cache.B().Get().Key(userID).Build()

	tknBytes, err := rs.cache.Do(ctx, getCmd).AsBytes()
	if err != nil {
		return "", err
	}

	refreshTkn = string(tknBytes[:])

	return refreshTkn, nil
}

func (rs *RefreshService) SetRefreshToken(tkn, userID string) error {
	refreshTTL, err := strconv.Atoi(rs.env.RefreshTokenTTL)
	if err != nil {
		fmt.Println(err)
		return err
	}

	refCmd := rs.cache.B().Set().Key(userID).Value(tkn).Ex(time.Hour * 24 * time.Duration(refreshTTL)).Build()
	_, err = rs.cache.Do(ctx, refCmd).AsBytes()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (rs *RefreshService) DeleteRefreshToken(userID string) bool {
	delCmd := rs.cache.B().Del().Key(userID).Build()
	_, err := rs.cache.Do(ctx, delCmd).AsBool()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
