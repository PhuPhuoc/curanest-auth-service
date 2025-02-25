package accountrpcservice

import (
	"github.com/PhuPhuoc/curanest-auth-service/middleware"
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	"github.com/gin-gonic/gin"
)

type accountRPCService struct {
	cmd   accountcommands.Commands
	query accountqueries.Queries
	auth  middleware.AuthClient
}

func NewAccountRPCService(cmd accountcommands.Commands, query accountqueries.Queries) *accountRPCService {
	return &accountRPCService{
		cmd:   cmd,
		query: query,
	}
}

func (s *accountRPCService) AddAuth(auth middleware.AuthClient) *accountRPCService {
	s.auth = auth
	return s
}

func (s *accountRPCService) Routes(g *gin.RouterGroup) {
	account_route := g.Group("/accounts")
	{
		account_route.POST("", s.handleCreateAccount())
		account_route.POST(
			"/by-ids",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin", "staff"),
			s.handleGetAccountByIds(),
		)
		account_route.POST(
			"/filter",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin", "staff"),
			s.handleGetAccountsWithFilter(),
		)
		account_route.GET(
			"/me",
			middleware.RequireAuth(s.auth),
			s.handleGetMyAccount(),
		)
		account_route.PUT(
			"/:account-id",
			middleware.RequireAuth(s.auth),
			s.handleUpdateAccount(),
		)
		account_route.PATCH(
			"/:account-id/role",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleUpdateRoleOfAccountNurseAndStaff(),
		)
		account_route.DELETE(
			"/:account-id",
			middleware.RequireAuth(s.auth),
			middleware.RequireRole("admin"),
			s.handleHardDeleteAccount(),
		)
	}
}
