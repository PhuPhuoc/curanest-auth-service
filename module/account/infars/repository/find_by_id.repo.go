package accountrepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

func (r *accountRepo) FindById(ctx context.Context, id uuid.UUID) (*accountdomain.Account, error) {
	var accdto AccountDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE, GET_FIELD, &where)
	if err := r.db.Get(&accdto, query, id); err != nil {
		return nil, err
	}
	return accdto.ToEntity()
}
