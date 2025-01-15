package accountqueries

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

type findByEmailHandler struct {
	queryRepo AccountQueryRepo
}

func NewFindByEmailHandler(queryRepo AccountQueryRepo) *findByEmailHandler {
	return &findByEmailHandler{
		queryRepo: queryRepo,
	}
}

func (h *findByEmailHandler) Handle(ctx context.Context, email string) (*accountdomain.Account, error) {
	entityFound, err := h.queryRepo.FindByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.NewInternalServerError().
				WithReason("cannot get entity from db").
				WithInner(err.Error())
		}
		return nil, err
	}
	return entityFound, nil
}
