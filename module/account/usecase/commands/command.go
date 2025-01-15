package accountcommands

import (
	"context"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

type Commands struct {
	CreateAccount *createAccountHandler
}

type Builder interface {
	BuildAccountCmdRepo() AccountCommandRepo
	BuildRoleFetcherRepo() RoleFetcher
}

func NewAccountCmdWithBuilder(b Builder) Commands {
	return Commands{
		CreateAccount: NewCreateAccountHandler(
			b.BuildAccountCmdRepo(),
			b.BuildRoleFetcherRepo(),
		),
	}
}

type AccountCommandRepo interface {
	Create(ctx context.Context, entity *accountdomain.Account) error
}

type RoleFetcher interface {
	GetRoleIdByName(ctx context.Context, roleName string) (*uuid.UUID, error)
}
