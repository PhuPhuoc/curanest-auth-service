package accountrpcservice

import (
	"errors"
	"fmt"

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

		fmt.Println("dto: ", dto)
		dtoFindByEmail, err := s.query.FindByEmail.Handle(ctx, dto.Email)
		if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
			common.ResponseError(ctx, err)
			return
		}
		if dtoFindByEmail != nil {
			common.ResponseError(ctx, fmt.Errorf("email: %s has been existed", dto.Email))
			return
		}

		dtoFindByPhone, err := s.query.FindByPhone.Handle(ctx, dto.PhoneNumber)
		if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
			common.ResponseError(ctx, err)
			return
		}
		if dtoFindByPhone != nil {
			common.ResponseError(ctx, fmt.Errorf("phone number: %s has been existed", dto.PhoneNumber))
			return
		}

		id, err := s.cmd.CreateAccount.Handle(ctx, &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, id)
	}
}
