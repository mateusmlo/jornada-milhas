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

type refreshService struct {
	cache rueidis.Client
	env   *config.Env
}

// NewRefreshService creates new refreshService instance
func NewRefreshService(cache *rueidis.Client, env *config.Env) RefreshService {
	return &refreshService{
		cache: *cache,
		env:   env,
	}
}

// GetRefreshToken tries to get user refresh token from cache
func (rs *refreshService) GetRefreshToken(userID string) (string, error) {
	var refreshTkn string
	getCmd := rs.cache.B().Get().Key(userID).Build()

	tknBytes, err := rs.cache.Do(ctx, getCmd).AsBytes()
	if err != nil {
		return "", err
	}

	refreshTkn = string(tknBytes[:])

	return refreshTkn, nil
}

// SetRefreshToken saves user refresh token to cache
func (rs *refreshService) SetRefreshToken(tkn, userID string) error {
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

// DeleteRefreshToken deletes user refresh token from cache
func (rs *refreshService) DeleteRefreshToken(userID string) bool {
	delCmd := rs.cache.B().Del().Key(userID).Build()
	_, err := rs.cache.Do(ctx, delCmd).AsBool()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
