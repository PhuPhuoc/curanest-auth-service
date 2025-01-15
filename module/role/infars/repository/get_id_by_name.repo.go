package rolerepository

import (
	"context"

	"github.com/google/uuid"
)

func (r *roleRepo) GetRoleIdByName(ctx context.Context, roleName string) (*uuid.UUID, error) {
	var roleId *uuid.UUID
	query := `select id from ` + table + ` where name = ?`
	if err := r.db.GetContext(ctx, &roleId, query, roleName); err != nil {
		return nil, err
	}
	return roleId, nil
}
