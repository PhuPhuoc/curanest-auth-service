package accounthttpservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/middleware"
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	"github.com/gin-gonic/gin"
)

type accountHttpService struct {
	cmd   accountcommands.Commands
	query accountqueries.Queries
	auth  middleware.AuthClient
}

func NewAccountHTTPService(cmd accountcommands.Commands, query accountqueries.Queries) *accountHttpService {
	return &accountHttpService{
		cmd:   cmd,
		query: query,
	}
}

func (s *accountHttpService) AddAuth(auth middleware.AuthClient) *accountHttpService {
	s.auth = auth
	return s
}

func (s *accountHttpService) Routes(g *gin.RouterGroup) {
	acc_route := g.Group("/accounts")
	{
		acc_route.POST("/user-login", s.handleLoginByPhone())
		acc_route.POST("/admin-login", s.handleLoginByEmailForAdmin())
		acc_route.PATCH("/:account-id/status", middleware.RequireAuth(s.auth), s.handleUpdateAccountStatus())
	}
}
