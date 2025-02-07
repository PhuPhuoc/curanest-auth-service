package accountrepository

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

func (repo *accountRepo) HardDelete(ctx context.Context, accountId *uuid.UUID) error {
	where := "id=?"
	query := common.GenerateSQLQueries(common.HARD_DELETE, TABLE, FIELD, &where)

	result, err := repo.db.ExecContext(ctx, query, accountId.String())
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return common.ErrRecordNotFound
	}
	return nil
}
