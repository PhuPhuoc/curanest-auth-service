package accountrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

func (r *accountRepo) Update(ctx context.Context, entity *accountdomain.Account) error {
	dto := ToDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE, UPDATE_FIELD, &where)
	if _, err := r.db.NamedExec(query, dto); err != nil {
		return err
	}
	return nil
}
