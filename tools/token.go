package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func GenerateJWT(userID uuid.UUID) (string, error) {
	ttl, err := strconv.Atoi(viper.GetString("TOKEN_TTL"))
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

	return tkn.SignedString([]byte(viper.GetString("JWT_SECRET")))
}

func ValidateToken(ctx *gin.Context) error {
	tkn := ExtractToken(ctx)

	_, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func ExtractToken(ctx *gin.Context) string {
	bearerTkn := ctx.Request.Header.Get("Authorization")
	_, tkn, hasTkn := strings.Cut(bearerTkn, " ")
	if !hasTkn {
		return ""
	}

	return tkn
}

func ExtractTokenSub(ctx *gin.Context) (uuid.UUID, error) {
	tkn := ExtractToken(ctx)

	token, err := jwt.Parse(tkn, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tkn.Header["alg"])
		}

		return []byte(viper.GetString("JWT_SECRET")), nil
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
