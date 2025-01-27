package accountrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

func (repo *accountRepo) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*accountdomain.Account, error) {
	var accdto AccountDTO
	where := "phone_number=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE, FIELD, &where)
	if err := repo.db.Get(&accdto, query, phoneNumber); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return accdto.ToEntity()
}
