package accountcommands

import (
	"context"

	"github.com/google/uuid"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

type Commands struct {
	CreateAccount       *createAccountHandler
	UpdateAccount       *updateAccountHandler
	UpdateAccountStatus *updateAccountStatusHandler
	UpdateAccountRole   *updateAccountRoleHandler
	HardDeleteAccount   *hardDeleteAccountHandler
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
		UpdateAccountStatus: NewUpdateAccountStatusHandler(
			b.BuildAccountCmdRepo(),
		),
		UpdateAccountRole: NewUpdateAccountRoleHandler(
			b.BuildAccountCmdRepo(),
			b.BuildRoleFetcherRepoCmd(),
		),
		HardDeleteAccount: NewHardDeleteAccountHandler(
			b.BuildAccountCmdRepo(),
		),
	}
}

type AccountCommandRepo interface {
	Create(ctx context.Context, entity *accountdomain.Account) error
	Update(ctx context.Context, entity *accountdomain.Account) error
	UpdateRoleForNurseAndStaff(ctx context.Context, nurseId *uuid.UUID, roleId *uuid.UUID) error
	UpdateStatus(ctx context.Context, accId uuid.UUID, status accountdomain.Status) error

	HardDelete(ctx context.Context, accountId *uuid.UUID) error
}

type RoleFetcher interface {
	GetRoleIdByName(ctx context.Context, roleName string) (*uuid.UUID, error)
}
