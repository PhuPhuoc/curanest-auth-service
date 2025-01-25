package accountqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

type getAccountByIdsHandler struct {
	queryRepo   AccountQueryRepo
	roleFetcher RoleFetcher
}

func NewGetAccountByIdsHandler(queryRepo AccountQueryRepo, roleFetcher RoleFetcher) *getAccountByIdsHandler {
	return &getAccountByIdsHandler{
		queryRepo:   queryRepo,
		roleFetcher: roleFetcher,
	}
}

func (h *getAccountByIdsHandler) Handle(ctx context.Context, listQuery *AccountIdsQuery) ([]AccountDTO, error) {
	roldis, err := h.roleFetcher.GetRoleIdByName(ctx, listQuery.Role)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get account with role: " + listQuery.Role).
			WithInner(err.Error())
	}

	entities, err := h.queryRepo.GetAccountByIds(ctx, listQuery.Ids)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get entities account from db").
			WithInner(err.Error())
	}

	list_dto := make([]AccountDTO, len(entities))
	for i := range entities {
		if roldis != nil && entities[i].GetRoleID() == *roldis {
			list_dto[i] = toDTO(&entities[i])
		}
	}
	return list_dto, nil
}
