package builder

import (
	rolerepository "github.com/PhuPhuoc/curanest-auth-service/module/role/infars/repository"
	rolequeries "github.com/PhuPhuoc/curanest-auth-service/module/role/usecase/queries"
	"github.com/jmoiron/sqlx"
)

type builderForRole struct {
	db *sqlx.DB
}

func NewRoleBuilder(db *sqlx.DB) builderForRole {
	return builderForRole{db: db}
}

// func (s builderForRole) BuildCategoryCmdRepo() categorycommands.CategoryCommandRepo {
// 	return categoryrepository.NewCategoryRepo(s.db)
// }

func (s builderForRole) BuildRoleQueryRepo() rolequeries.RoleQueryRepo {
	return rolerepository.NewRoleRepo(s.db)
}
