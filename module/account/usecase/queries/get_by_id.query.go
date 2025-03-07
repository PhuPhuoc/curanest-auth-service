package accountqueries

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

type getAccountByIdHandler struct {
	queryRepo   AccountQueryRepo
	roleFetcher RoleFetcher
}

func NewGetAccountByIdHandler(queryRepo AccountQueryRepo, roleFetcher RoleFetcher) *getAccountByIdHandler {
	return &getAccountByIdHandler{
		queryRepo:   queryRepo,
		roleFetcher: roleFetcher,
	}
}

func (h *getAccountByIdHandler) Handle(ctx context.Context, accId uuid.UUID) (*MyAccountDTO, error) {
	entity, err := h.queryRepo.FindById(ctx, accId)
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
