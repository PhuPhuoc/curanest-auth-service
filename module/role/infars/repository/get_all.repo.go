package rolerepository

import (
	"context"

	roledomain "github.com/PhuPhuoc/curanest-auth-service/module/role/domain"
)

func (repo *roleRepo) GetRoles(ctx context.Context) ([]roledomain.Role, error) {
	query := "select " + field + " from " + table

	var dtos []RoleDTO
	if err := repo.db.SelectContext(ctx, &dtos, query); err != nil {
		return nil, err
	}

	entities := make([]roledomain.Role, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToEntity()
		entities[i] = *entity
	}

	return entities, nil
}
