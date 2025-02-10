package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		update role account for nurse or staff (admin)
// @Description	update role account for nurse or staff (admin)
// @Tags			rpc: account
// @Accept			json
// @Produce		json
// @Param			account-id	path		string					true								"Account ID (UUID)"
// @Param			update		form		body					accountcommands.UpdateRoleRequest	true	"account data to update"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/external/rpc/accounts/{account-id}/role [patch]
// @Security		ApiKeyAuth
func (s *accountRPCService) handleUpdateRoleOfAccountNurseAndStaff() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto accountcommands.UpdateRoleRequest
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

		err = s.cmd.UpdateAccountRole.Handle(ctx.Request.Context(), &accountUUID, &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}
