package accountqueries

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
)

type findByPhoneHandler struct {
	queryRepo AccountQueryRepo
}

func NewFindByPhoneHandler(queryRepo AccountQueryRepo) *findByPhoneHandler {
	return &findByPhoneHandler{
		queryRepo: queryRepo,
	}
}

func (h *findByPhoneHandler) Handle(ctx context.Context, phoneNumber string) (*accountdomain.Account, error) {
	entityFound, err := h.queryRepo.FindByPhoneNumber(ctx, phoneNumber)
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
