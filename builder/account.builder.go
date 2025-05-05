package builder

import (
	accountnotirpc "github.com/PhuPhuoc/curanest-auth-service/module/account/infars/externalapi/noti"
	accountrepository "github.com/PhuPhuoc/curanest-auth-service/module/account/infars/repository"
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	rolerepository "github.com/PhuPhuoc/curanest-auth-service/module/role/infars/repository"
	"github.com/jmoiron/sqlx"
)

type accountBuilder struct {
	db *sqlx.DB
	tp accountqueries.TokenProvider

	urlPathNotiService string
}

func (s accountBuilder) AddTokenProvider(tp accountqueries.TokenProvider) accountBuilder {
	s.tp = tp
	return s
}

func (s accountBuilder) AddUrlPathNotiService(url string) accountBuilder {
	s.urlPathNotiService = url
	return s
}

func NewAccountBuilder(db *sqlx.DB) accountBuilder {
	return accountBuilder{db: db}
}

func (s accountBuilder) BuildAccountCmdRepo() accountcommands.AccountCommandRepo {
	return accountrepository.NewAccountRepo(s.db)
}

func (s accountBuilder) BuildAccountQueryRepo() accountqueries.AccountQueryRepo {
	return accountrepository.NewAccountRepo(s.db)
}

func (s accountBuilder) BuildRoleFetcherRepoCmd() accountcommands.RoleFetcher {
	return rolerepository.NewRoleRepo(s.db)
}

func (s accountBuilder) BuildRoleFetcherRepoQuery() accountqueries.RoleFetcher {
	return rolerepository.NewRoleRepo(s.db)
}

func (s accountBuilder) BuilderTokenProvider() accountqueries.TokenProvider {
	return s.tp
}

func (s accountBuilder) BuildExternalNotiServiceQuery() accountqueries.ExternalNotiService {
	return accountnotirpc.NewNotiRPC(s.urlPathNotiService)
}
