package accounthttpservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		login by email for admin
//	@Description	login by email for admin
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			account-id	path		string					true									"Account ID (UUID)"
//	@Param			create		form		body					accountcommands.UpdateAccountStatusDTO	true	"account update data"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/api/v1/accounts/{account-id}/status [patch]
//	@Security		ApiKeyAuth
func (s *accountHttpService) handleUpdateAccountStatus() gin.HandlerFunc {
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

		var dto accountcommands.UpdateAccountStatusDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		err = s.cmd.UpdateAccountStatus.Handle(ctx, accountUUID, &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}
