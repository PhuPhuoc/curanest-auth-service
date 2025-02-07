package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		delete account
// @Description	delete account
// @Tags			rpc: account
// @Accept			json
// @Produce		json
// @Param			account-id	path		string					true	"Account ID (UUID)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/external/rpc/accounts/{account-id} [delete]
// @Security		ApiKeyAuth
func (s *accountRPCService) handleHardDeleteAccount() gin.HandlerFunc {
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

		err = s.cmd.HardDeleteAccount.Handle(ctx.Request.Context(), &accountUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseDeleted(ctx)
	}
}
