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
	VerifyEmail       *validateAccountEmailHandler
	VerifyPhoneNumber *validateAccountPhoneNumberHandler

	LoginByPhone *loginByPhonePasswordHandler

	GetById              *getAccountByIdHandler
	GetByIds             *getAccountByIdsHandler
	GetMyAccount         *getMyAccountHandler
	GetAccountWithFilter *getAccountWithFilterHandler
}

type Builder interface {
	BuildAccountQueryRepo() AccountQueryRepo
	BuilderTokenProvider() TokenProvider
	BuildRoleFetcherRepoQuery() RoleFetcher
}

func NewAccountQueryWithBuilder(b Builder) Queries {
	return Queries{
		VerifyEmail:       NewValidateAccountEmailHandler(b.BuildAccountQueryRepo()),
		VerifyPhoneNumber: NewVerifyPhoneHandler(b.BuildAccountQueryRepo()),

		LoginByPhone: NewLoginWithPhoneHandler(
			b.BuildAccountQueryRepo(),
			b.BuilderTokenProvider(),
			b.BuildRoleFetcherRepoQuery(),
		),

		GetByIds: NewGetAccountByIdsHandler(
			b.BuildAccountQueryRepo(),
			b.BuildRoleFetcherRepoQuery(),
		),
		GetMyAccount: NewGetMyAccountHandler(
			b.BuildAccountQueryRepo(),
			b.BuildRoleFetcherRepoQuery(),
		),
		GetById: NewGetAccountByIdHandler(
			b.BuildAccountQueryRepo(),
			b.BuildRoleFetcherRepoQuery(),
		),
		GetAccountWithFilter: NewGetAccountWithFilterHandler(
			b.BuildAccountQueryRepo(),
			b.BuildRoleFetcherRepoQuery(),
		),
	}
}

type AccountQueryRepo interface {
	FindByEmail(ctx context.Context, email string) (*accountdomain.Account, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*accountdomain.Account, error)
	FindById(ctx context.Context, id uuid.UUID) (*accountdomain.Account, error)

	GetAccountByIds(ctx context.Context, ids []uuid.UUID) ([]accountdomain.Account, error)
	GetAccountWithFilter(ctx context.Context, filter *FilterAccountQuery) ([]accountdomain.Account, error)
}

type RoleFetcher interface {
	GetNameByRoleId(ctx context.Context, id uuid.UUID) (string, error)
	GetRoleIdByName(ctx context.Context, roleName string) (*uuid.UUID, error)
}
