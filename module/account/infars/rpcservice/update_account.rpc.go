package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		update account
// @Description	update account
// @Tags			rpc: account
// @Accept			json
// @Produce		json
// @Param			account-id	path		string					true								"Account ID (UUID)"
// @Param			update		form		body					accountcommands.UpdateAccountCmdDTO	true	"account data to update"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/external/rpc/accounts/{account-id} [put]
// @Security		ApiKeyAuth
func (s *accountRPCService) handleUpdateAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto accountcommands.UpdateAccountCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

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

		err = s.cmd.UpdateAccount.Handle(ctx.Request.Context(), &accountUUID, &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}
