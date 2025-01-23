package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthClient interface {
	ParseToken(ctx context.Context, tokenString string) (map[string]interface{}, error)
}

func RequireAuth(ac AuthClient) func(*gin.Context) {
	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))
		if err != nil {
			common.ResponseError(ctx, err)
			ctx.Abort()
			return
		}

		requester, err := introspectToken(ctx.Request.Context(), ac, token)
		if err != nil {
			common.ResponseError(ctx, err)
			ctx.Abort()
			return
		}

		key := common.KeyRequester
		newCtx := context.WithValue(ctx.Request.Context(), key, requester)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}

func introspectToken(ctx context.Context, ac AuthClient, accessToken string) (common.Requester, error) {
	claim, err := ac.ParseToken(ctx, accessToken)
	if err != nil {
		return nil, errors.New("cannot parse token")
	}

	var id, sub, role string
	var ok bool
	if id, ok = claim["id"].(string); !ok {
		return nil, fmt.Errorf("token invalid - missing field id")
	}
	if sub, ok = claim["sub"].(string); !ok {
		return nil, fmt.Errorf("token invalid - missing field sub")
	}
	if role, ok = claim["role"].(string); !ok {
		return nil, fmt.Errorf("token invalid - missing field role")
	}

	iduuid := uuid.MustParse(id)
	subuuid := uuid.MustParse(sub)

	requester := common.NewRequester(subuuid, iduuid, role)
	return requester, nil
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ") //"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("missing access token")
	}

	return parts[1], nil
}
