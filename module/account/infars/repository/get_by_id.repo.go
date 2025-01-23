package accountrepository

import (
	"context"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

func (r *accountRepo) FindById(ctx context.Context, id uuid.UUID) (*accountdomain.Account, error) {
	var accdto AccountDTO
	query := `select ` + getField + ` from ` + table + ` where id=?`
	if err := r.db.Get(&accdto, query, id); err != nil {
		return nil, err
	}
	return accdto.ToEntity()
}
