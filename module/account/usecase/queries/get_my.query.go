package accountqueries

import (
	"context"
	"time"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
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

type MyAccountDTO struct {
	Id          uuid.UUID `json:"id"`
	RoleId      uuid.UUID `json:"-"`
	Role        string    `json:"role"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"created-at"`
}

func toMyAccDTO(data *accountdomain.Account) *MyAccountDTO {
	dto := &MyAccountDTO{
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
