package accountqueries

import (
	"context"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

type TokenProvider interface {
	IssueToken(ctx context.Context, id, sub, role string) (token string, err error)
	ParseToken(ctx context.Context, tokenString string) (claims map[string]interface{}, err error)
	TokenExpireInSeconds() int
}

type Queries struct {
	FindByEmail *findByEmailHandler
	FindByPhone *findByPhoneHandler

	LoginByPhone *loginByPhonePasswordHandler
}

type Builder interface {
	BuildAccountQueryRepo() AccountQueryRepo
	BuilderTokenProvider() TokenProvider
	BuildRoleFetcherRepoQuery() RoleFetcher
}

func NewAccountQueryWithBuilder(b Builder) Queries {
	return Queries{
		FindByEmail: NewFindByEmailHandler(b.BuildAccountQueryRepo()),
		FindByPhone: NewFindByPhoneHandler(b.BuildAccountQueryRepo()),

		LoginByPhone: NewLoginWithPhoneHandler(
			b.BuildAccountQueryRepo(),
			b.BuilderTokenProvider(),
			b.BuildRoleFetcherRepoQuery(),
		),
	}
}

type AccountQueryRepo interface {
	FindByEmail(ctx context.Context, email string) (*accountdomain.Account, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*accountdomain.Account, error)
}

type RoleFetcher interface {
	GetNameByRoleId(ctx context.Context, id uuid.UUID) (string, error)
}
