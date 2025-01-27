package accountrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

func (r *accountRepo) Create(ctx context.Context, entity *accountdomain.Account) error {
	accdto := ToDTO(entity)
	query := common.GenerateSQLQueries(common.INSERT, TABLE, FIELD, nil)
	if _, err := r.db.NamedExec(query, accdto); err != nil {
		return err
	}
	return nil
}
