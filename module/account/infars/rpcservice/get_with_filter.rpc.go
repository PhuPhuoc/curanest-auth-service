package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	"github.com/gin-gonic/gin"
)

// @Summary		get account with filter
// @Description	get account with filter
// @Tags			rpc: account
// @Accept			json
// @Produce		json
// @Param			create	form		body					accountqueries.FilterAccountQuery	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/external/rpc/accounts/filter [post]
// @Security		ApiKeyAuth
func (s *accountRPCService) handleGetAccountsWithFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var param accountqueries.FilterAccountQuery

		if err := ctx.Bind(&param); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		result, err := s.query.GetAccountWithFilter.Handle(ctx, &param)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseGetWithPagination(ctx, result, param.Paging, param.Filter)
	}
}
