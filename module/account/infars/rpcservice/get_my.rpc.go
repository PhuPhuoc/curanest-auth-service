package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/gin-gonic/gin"
)

//	@Summary		get my account by token
//	@Description	get my account by token
//	@Tags			rpc: account
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"data"
//	@Failure		400	{object}	error					"Bad request error"
//	@Router			/external/rpc/accounts/me [get]
//	@Security		ApiKeyAuth
func (s *accountRPCService) handleGetMyAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		info, err := s.query.GetMyAccount.Handle(ctx.Request.Context())
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, info)
	}
}
