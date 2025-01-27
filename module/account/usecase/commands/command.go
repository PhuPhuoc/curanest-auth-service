package accountcommands

import (
	"context"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

type Commands struct {
	CreateAccount *createAccountHandler
	UpdateAccount *updateAccountHandler
}

type Builder interface {
	BuildAccountCmdRepo() AccountCommandRepo
	BuildRoleFetcherRepoCmd() RoleFetcher
}

func NewAccountCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateAccount: NewCreateAccountHandler(
			b.BuildAccountCmdRepo(),
			b.BuildRoleFetcherRepoCmd(),
		),
		UpdateAccount: NewUpdateAccountHandler(
			b.BuildAccountCmdRepo(),
		),
	}
}

type AccountCommandRepo interface {
	Create(ctx context.Context, entity *accountdomain.Account) error
	Update(ctx context.Context, entity *accountdomain.Account) error
}

type RoleFetcher interface {
	GetRoleIdByName(ctx context.Context, roleName string) (*uuid.UUID, error)
}
