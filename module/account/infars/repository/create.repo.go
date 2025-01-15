package accountrepository

import (
	"context"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

func (r *accountRepo) Create(ctx context.Context, entity *accountdomain.Account) error {
	accdto := ToDTO(entity)
	query := `insert into ` + table + ` (` + field + `) values (` + mapping + `)`
	if _, err := r.db.NamedExec(query, accdto); err != nil {
		return err
	}
	return nil
}
