package accountcommands

import (
	"context"

	"github.com/google/uuid"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

type Commands struct {
	CreateAccount     *createAccountHandler
	UpdateAccount     *updateAccountHandler
	HardDeleteAccount *hardDeleteAccountHandler
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
		HardDeleteAccount: NewHardDeleteAccountHandler(
			b.BuildAccountCmdRepo(),
		),
	}
}

type AccountCommandRepo interface {
	Create(ctx context.Context, entity *accountdomain.Account) error
	Update(ctx context.Context, entity *accountdomain.Account) error

	HardDelete(ctx context.Context, accountId *uuid.UUID) error
}

type RoleFetcher interface {
	GetRoleIdByName(ctx context.Context, roleName string) (*uuid.UUID, error)
}
