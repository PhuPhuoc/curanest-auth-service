package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	"github.com/gin-gonic/gin"
)

// @Summary		create account
// @Description	create account
// @Tags			rpc: account
// @Accept			json
// @Produce		json
// @Param			create	form		body					accountcommands.CreateAccountCmdDTO	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/external/rpc/accounts [post]
func (s *accountRPCService) handleCreateAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto accountcommands.CreateAccountCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		if err := s.query.VerifyEmail.Handle(ctx, dto.Email); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		if err := s.query.VerifyPhoneNumber.Handle(ctx, dto.PhoneNumber); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		resp, err := s.cmd.CreateAccount.Handle(ctx, &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, resp)
	}
}
