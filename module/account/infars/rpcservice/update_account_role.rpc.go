package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		update role account for nurse or staff (admin)
//	@Description	update role account for nurse or staff (admin)
//	@Tags			rpc: account
//	@Accept			json
//	@Produce		json
//	@Param			account-id	path		string					true	"Account ID (UUID)"
//	@Param			target-role	query		string					true	"role to tranfer"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/external/rpc/accounts/{account-id}/role [patch]
//	@Security		ApiKeyAuth
func (s *accountRPCService) handleUpdateRoleOfAccountNurseAndStaff() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		targetRole := ctx.Query("target-role")
		if targetRole == "" || targetRole != "staff" && targetRole != "nurse" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing target-role - role must be 'staff' or 'nurse'"))
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

		err = s.cmd.UpdateAccountRole.Handle(ctx.Request.Context(), &accountUUID, targetRole)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}
