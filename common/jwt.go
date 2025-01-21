package common

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	defaultExpireTokenInSeconds   = 60 * 60 * 24 * 7  // 7d
	defaultExpireRefreshInSeconds = 60 * 60 * 24 * 14 // 14d
)

type JWTx struct {
	secret               string
	expireTokenInSeconds int
	// expireRefreshInSeconds int
}

func NewJWTx(secret string, expireTokenInSeconds int) *JWTx {
	return &JWTx{
		secret:               secret,
		expireTokenInSeconds: expireTokenInSeconds,
		// expireRefreshInSeconds: expireRefreshInSeconds,
	}
}

func (j *JWTx) TokenExpireInSeconds() int { return j.expireTokenInSeconds }

func (j *JWTx) IssueToken(ctx context.Context, id, sub, role string) (string, error) {
	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"sub":  sub,
		"role": role,
		"exp":  now.Add(time.Second * time.Duration(j.expireTokenInSeconds)).Unix(),
		"iat":  now.Unix(),
		"id":   id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func (j *JWTx) ParseToken(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return claims, nil
}
