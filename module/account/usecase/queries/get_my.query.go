package accountqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

type getMyAccountHandler struct {
	queryRepo   AccountQueryRepo
	roleFetcher RoleFetcher
}

func NewGetMyAccountHandler(queryRepo AccountQueryRepo, roleFetcher RoleFetcher) *getMyAccountHandler {
	return &getMyAccountHandler{
		queryRepo:   queryRepo,
		roleFetcher: roleFetcher,
	}
}

func (h *getMyAccountHandler) Handle(ctx context.Context) (*MyAccountDTO, error) {
	requester, ok := ctx.Value(common.KeyRequester).(common.Requester)
	if !ok {
		return nil, common.NewUnauthorizedError()
	}
	sub := requester.UserId()
	entity, err := h.queryRepo.FindById(ctx, sub)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get entities account from db").
			WithInner(err.Error())
	}

	dto := toMyAccDTO(entity)
	role_name, _ := h.roleFetcher.GetNameByRoleId(ctx, dto.RoleId)
	dto.Role = role_name

	return dto, nil
}
