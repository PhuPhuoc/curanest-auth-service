package rolerepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/google/uuid"
)

func (r *roleRepo) GetRoleIdByName(ctx context.Context, roleName string) (*uuid.UUID, error) {
	var roleId *uuid.UUID
	query := `select id from ` + table + ` where name = ?`
	if err := r.db.GetContext(ctx, &roleId, query, roleName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return roleId, nil
}
