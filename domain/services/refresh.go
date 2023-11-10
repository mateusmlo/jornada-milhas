package service

import (
	"context"

	"github.com/redis/rueidis"
)

var ctx = context.Background()

type RefreshService struct {
	cache rueidis.Client
}

func (rs *RefreshService) GetRefreshToken(tkn string) (string, error) {
	var refreshTkn string
	getCmd := rs.cache.B().Get().Key("refreshTkn:" + tkn).Build()

	tknBytes, err := rs.cache.Do(ctx, getCmd).AsBytes()
	if err != nil {
		return "", err
	}

	refreshTkn = string(tknBytes[:])

	return refreshTkn, nil
}

/* func (rs *RefreshService) SetRefreshToken(tkn string) error {

}
*/
