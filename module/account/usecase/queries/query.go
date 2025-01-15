package accountqueries

import (
	"context"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

type Queries struct {
	FindByEmail *findByEmailHandler
	FindByPhone *findByPhoneHandler
}

type Builder interface {
	BuildAccountQueryRepo() AccountQueryRepo
}

func NewAccountQueryWithBuilder(b Builder) Queries {
	return Queries{
		FindByEmail: NewFindByEmailHandler(b.BuildAccountQueryRepo()),
		FindByPhone: NewFindByPhoneHandler(b.BuildAccountQueryRepo()),
	}
}

type AccountQueryRepo interface {
	FindByEmail(ctx context.Context, email string) (*accountdomain.Account, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*accountdomain.Account, error)
}
