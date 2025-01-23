package accountqueries

import (
	"context"
	"time"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
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

type AccountIdsQuery struct {
	Role string      `json:"role"`
	Ids  []uuid.UUID `json:"ids"`
}

type AccountDTO struct {
	Id          uuid.UUID `json:"id"`
	RoleId      uuid.UUID `json:"-"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created-at"`
}

func toDTO(data *accountdomain.Account) AccountDTO {
	dto := AccountDTO{
		Id:          data.GetID(),
		RoleId:      data.GetRoleID(),
		FullName:    data.GetFullName(),
		Email:       data.GetEmail(),
		PhoneNumber: data.GetPhoneNumber(),
		Avatar:      data.GetAvatar(),
		CreatedAt:   data.GetCreatedAt(),
	}
	return dto
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
