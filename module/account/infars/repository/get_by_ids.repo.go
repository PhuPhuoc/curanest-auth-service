package accountrepository

import (
	"context"
	"strings"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (r *accountRepo) GetAccountByIds(ctx context.Context, ids []uuid.UUID) ([]accountdomain.Account, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	dtos := []AccountDTO{}
	fields := strings.Join(GET_FIELD, ", ")
	query := "select " + fields + " from " + TABLE + " where id in (?)"

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		return nil, err
	}

	if err := r.db.Select(&dtos, query, args...); err != nil {
		return nil, err
	}

	entities := make([]accountdomain.Account, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToEntity()
		entities[i] = *entity
	}
	return entities, nil
}
