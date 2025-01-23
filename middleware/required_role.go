package middleware

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, ok := ctx.Request.Context().Value(common.KeyRequester).(common.Requester)
		if !ok {
			common.ResponseUnauthorizedError(ctx, "cannot found requester info")
			ctx.Abort()
			return
		}

		for _, role := range allowedRoles {
			if req.Role() == role {
				ctx.Next()
				return
			}
		}

		common.ResponseFobiddenError(ctx, "your role cannot use this api")
		ctx.Abort()
	}
}
