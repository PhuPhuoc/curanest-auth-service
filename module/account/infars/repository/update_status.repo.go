package accountrepository

import (
	"context"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

func (r *accountRepo) UpdateStatus(ctx context.Context, accId uuid.UUID, status accountdomain.Status) error {
	query := `UPDATE accounts SET status = :status WHERE id = :id`

	params := map[string]interface{}{
		"id":     accId,
		"status": status,
	}

	if _, err := r.db.NamedExecContext(ctx, query, params); err != nil {
		return err
	}
	return nil
}
