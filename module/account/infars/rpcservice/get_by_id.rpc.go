package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		get account by id
//	@Description	get account by id
//	@Tags			rpc: account
//	@Accept			json
//	@Produce		json
//	@Param			account-id	path		string					true	"Account ID (UUID)"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/external/rpc/accounts/{account-id} [get]
//	@Security		ApiKeyAuth
func (s *accountRPCService) handleGetAccountById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accountId := ctx.Param("account-id")
		if accountId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing patient-id"))
			return
		}

		accountUUID, err := uuid.Parse(accountId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("id invalid (not a uuid)"))
			return
		}
		info, err := s.query.GetById.Handle(ctx, accountUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, info)
	}
}
