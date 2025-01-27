package accountcommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

type updateAccountHandler struct {
	commandrepo AccountCommandRepo
}

func NewUpdateAccountHandler(cmdRepo AccountCommandRepo) *updateAccountHandler {
	return &updateAccountHandler{
		commandrepo: cmdRepo,
	}
}

type UpdateAccountCmdDTO struct {
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
}

func (h *updateAccountHandler) Handle(ctx context.Context, id *uuid.UUID, dto *UpdateAccountCmdDTO) error {
	entity, _ := accountdomain.NewAccount(
		*id,
		uuid.Nil,
		dto.FullName,
		dto.PhoneNumber,
		dto.Email,
		"",
		"",
		dto.Avatar,
		accountdomain.StatusActivated,
		nil,
	)

	if err := h.commandrepo.Update(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot get insert account into db").
			WithInner(err.Error())
	}
	return nil
}
