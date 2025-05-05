package accounthttpservice

import (
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	"github.com/gin-gonic/gin"
)

type accountHttpService struct {
	query accountqueries.Queries
}

func NewAccountHTTPService(query accountqueries.Queries) *accountHttpService {
	return &accountHttpService{
		query: query,
	}
}

func (s *accountHttpService) Routes(g *gin.RouterGroup) {
	acc_route := g.Group("/accounts")
	{
		acc_route.POST("/user-login", s.handleLoginByPhone())
		acc_route.POST("/admin-login", s.handleLoginByEmailForAdmin())
	}
}
