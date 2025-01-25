package accountqueries

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

type getAccountWithFilterHandler struct {
	queryRepo   AccountQueryRepo
	roleFetcher RoleFetcher
}

func NewGetAccountWithFilterHandler(queryRepo AccountQueryRepo, roleFetcher RoleFetcher) *getAccountWithFilterHandler {
	return &getAccountWithFilterHandler{
		queryRepo:   queryRepo,
		roleFetcher: roleFetcher,
	}
}

func (h *getAccountWithFilterHandler) Handle(ctx context.Context, filter *FilterAccountQuery) ([]AccountDTO, error) {
	filter.Paging.Process()
	if filter.Filter.Role != "" {
		roleid, err := h.roleFetcher.GetRoleIdByName(ctx, filter.Filter.Role)
		if err != nil {
			if !errors.Is(err, common.ErrRecordNotFound) {
				return nil, common.NewInternalServerError().
					WithReason("unable to perform the step of checking email exists in the system").
					WithInner(err.Error())
			}
			return nil, common.NewBadRequestError().
				WithReason("invalid role: " + filter.Filter.Role).
				WithInner(err.Error())
		}
		filter.Filter.RoleId = roleid.String()
	}

	entities, err := h.queryRepo.GetAccountWithFilter(ctx, filter)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get entities account from db").
			WithInner(err.Error())
	}

	list_dto := make([]AccountDTO, len(entities))
	for i := range entities {
		list_dto[i] = toDTO(&entities[i])
	}
	return list_dto, nil
}
