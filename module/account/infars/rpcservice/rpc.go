package accountrpcservice

import (
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	"github.com/gin-gonic/gin"
)

type accountRPCService struct {
	cmd   accountcommands.Commands
	query accountqueries.Queries
}

func NewAccountRPCService(cmd accountcommands.Commands, query accountqueries.Queries) *accountRPCService {
	return &accountRPCService{
		cmd:   cmd,
		query: query,
	}
}

func (s *accountRPCService) Routes(g *gin.RouterGroup) {
	account_route := g.Group("/accounts")
	{
		account_route.POST("", s.handleCreateAccount())
	}
}
