package rolehttpservice

import (
	rolequeries "github.com/PhuPhuoc/curanest-auth-service/module/role/usecase/queries"
	"github.com/gin-gonic/gin"
)

type roleHttpService struct {
	// cmd   rolecommands.Commands
	query rolequeries.Queries
}

func NewCategoryHTTPService(query rolequeries.Queries) *roleHttpService {
	return &roleHttpService{
		query: query,
	}
}

func (s *roleHttpService) Routes(g *gin.RouterGroup) {
	role_route := g.Group("/roles")
	{
		role_route.GET("", s.handleGetRoles())
	}
}
