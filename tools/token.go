package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/config"
)

type TokenUtils struct {
	env *config.Env
}

func NewTokenUtils(env *config.Env) *TokenUtils {
	return &TokenUtils{
		env: env,
	}
}

func (tu *TokenUtils) GenerateAccessToken(userID uuid.UUID) (string, error) {
	ttl, err := strconv.Atoi(tu.env.AccessTokenTTL)
	if err != nil {
		fmt.Println(err)

		return "", err
	}

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * time.Duration(ttl)).Unix(),
		"iss": "ALTIMIT Corp.",
		"iat": time.Now().Unix(),
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tkn.SignedString([]byte(tu.env.AccessTokenSecret))
}

func (tu *TokenUtils) GenerateRefreshToken(userID uuid.UUID) (string, error) {
	ttl, err := strconv.Atoi(tu.env.RefreshTokenTTL)
	if err != nil {
		fmt.Println(err)

		return "", err
	}

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24 * time.Duration(ttl)).Unix(),
		"iss": "ALTIMIT Corp.",
		"iat": time.Now().Unix(),
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tkn.SignedString([]byte(tu.env.RefreshTokenSecret))
}

func (tu *TokenUtils) ValidateAccessToken(ctx *gin.Context) error {
	tkn := tu.ExtractToken(ctx)

	_, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tu.env.AccessTokenSecret), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (tu *TokenUtils) ValidateRefreshToken(ctx *gin.Context) error {
	tkn := tu.ExtractToken(ctx)

	_, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tu.env.RefreshTokenSecret), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (tu *TokenUtils) ExtractToken(ctx *gin.Context) string {
	bearerTkn := ctx.Request.Header.Get("Authorization")
	_, tkn, hasTkn := strings.Cut(bearerTkn, " ")
	if !hasTkn {
		return ""
	}

	return tkn
}

func (tu *TokenUtils) ExtractTokenSub(ctx *gin.Context, isRefresh bool) (uuid.UUID, error) {
	tkn := tu.ExtractToken(ctx)

	token, err := jwt.Parse(tkn, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tkn.Header["alg"])
		}

		if isRefresh {
			return []byte(tu.env.RefreshTokenSecret), nil
		}

		return []byte(tu.env.AccessTokenSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := uuid.Parse(claims["sub"].(string))
		if err != nil {
			return uuid.UUID{}, err
		}

		return userID, nil
	}

	return uuid.UUID{}, err
}
