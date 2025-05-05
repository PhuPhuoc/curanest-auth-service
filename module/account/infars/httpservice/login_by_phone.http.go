package accounthttpservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	"github.com/gin-gonic/gin"
)

// @Summary		login by phone number
// @Description	login by phone number
// @Tags			accounts
// @Accept			json
// @Produce		json
// @Param			create	form		body					accountqueries.LoginByPhoneRequestDTO	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/accounts/user-login [post]
func (s *accountHttpService) handleLoginByPhone() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto accountqueries.LoginByPhoneRequestDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}
		data, err := s.query.LoginByPhone.Handle(ctx, &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseSuccess(ctx, data)
	}
}
