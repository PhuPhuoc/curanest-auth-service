package middleware

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, ok := ctx.Request.Context().Value(common.KeyRequester).(common.Requester)
		if !ok {
			err := common.NewUnauthorizedError().WithReason("cannot found requester data")
			common.ResponseError(ctx, err)
			ctx.Abort()
			return
		}

		for _, role := range allowedRoles {
			if req.Role() == role {
				ctx.Next()
				return
			}
		}

		err := common.NewForbiddenError().WithReason("your role cannot call this api")
		common.ResponseError(ctx, err)
		ctx.Abort()
	}
}
