package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	"github.com/gin-gonic/gin"
)

// @Summary		get account by ids
// @Description	get account by ids
// @Tags			rpc: account
// @Accept			json
// @Produce		json
// @Param			create	form		body					accountqueries.AccountIdsQuery	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/external/rpc/accounts/by-ids [post]
// @Security		ApiKeyAuth
func (s *accountRPCService) handleGetAccountByIds() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto accountqueries.AccountIdsQuery
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		ids, err := s.query.GetByIds.Handle(ctx, &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, ids)
	}
}
