package accountrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

func (repo *accountRepo) FindByEmail(ctx context.Context, email string) (*accountdomain.Account, error) {
	var accdto AccountDTO
	where := "email=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE, FIELD, &where)
	if err := repo.db.Get(&accdto, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return accdto.ToEntity()
}
