package accountrepository

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

func (repo *accountRepo) UpdateRoleForNurseAndStaff(ctx context.Context, nurseId *uuid.UUID, roleId *uuid.UUID) error {
	query := `update ` + TABLE + ` set role_id=? where id=?`
	result, err := repo.db.ExecContext(ctx, query, roleId, nurseId)
	if err != nil {
		return err
	}

	numIsChanged, _ := result.RowsAffected()
	if numIsChanged == 0 {
		return common.ErrNoRecordsAreChanged
	}
	return nil
}
