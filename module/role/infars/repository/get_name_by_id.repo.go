package rolerepository

import (
	"context"

	"github.com/google/uuid"
)

func (r *roleRepo) GetNameByRoleId(ctx context.Context, id uuid.UUID) (string, error) {
	var roleName string
	query := `select name from ` + table + ` where id = ?`
	if err := r.db.GetContext(ctx, &roleName, query, id); err != nil {
		return "", err
	}
	return roleName, nil
}
