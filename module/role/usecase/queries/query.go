package rolequeries

import (
	"context"

	roledomain "github.com/PhuPhuoc/curanest-auth-service/module/role/domain"
)

type Queries struct {
	GetAllRoles *getRolesHandler
}

type Builder interface {
	BuildRoleQueryRepo() RoleQueryRepo
}

func NewRoleQueryWithBuilder(b Builder) Queries {
	return Queries{
		GetAllRoles: NewGetRolesHandler(
			b.BuildRoleQueryRepo(),
		),
	}
}

type RoleQueryRepo interface {
	GetRoles(ctx context.Context) ([]roledomain.Role, error)
}
