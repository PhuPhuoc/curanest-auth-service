package accounthttpservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	"github.com/gin-gonic/gin"
)

// @Summary		login by email for admin
// @Description	login by email for admin
// @Tags			accounts
// @Accept			json
// @Produce		json
// @Param			create	form		body					accountqueries.LoginByEmailRequestDTO	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/accounts/admin-login [post]
func (s *accountHttpService) handleLoginByEmailForAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto accountqueries.LoginByEmailRequestDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}
		data, err := s.query.LoginByEmail.Handle(ctx, &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseSuccess(ctx, data)
	}
}
