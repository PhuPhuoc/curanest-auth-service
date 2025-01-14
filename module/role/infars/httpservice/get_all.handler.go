package rolehttpservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/gin-gonic/gin"
)

// @BasePath		/api/v1
// @Summary		get appointment (card)
// @Description	get appointment (card)
// @Tags			roles
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"data"
// @Failure		400	{object}	error					"Bad request error"
// @Router			/api/v1/roles [get]
func (s *roleHttpService) handleGetRoles() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := s.query.GetAllRoles.Handle(c)
		if err != nil {
			common.ResponseError(c, err)
			return
		}

		common.ResponseSuccess(c, result)
	}
}
