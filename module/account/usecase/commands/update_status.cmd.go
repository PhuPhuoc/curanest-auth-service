package accountcommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

type updateAccountStatusHandler struct {
	commandrepo AccountCommandRepo
}

func NewUpdateAccountStatusHandler(cmdRepo AccountCommandRepo) *updateAccountStatusHandler {
	return &updateAccountStatusHandler{
		commandrepo: cmdRepo,
	}
}

type UpdateAccountStatusDTO struct {
	NewStatus string `json:"new-status"`
}

func (h *updateAccountStatusHandler) Handle(ctx context.Context, accId uuid.UUID, dto *UpdateAccountStatusDTO) error {
	var newStatus accountdomain.Status
	if accountdomain.Enum(dto.NewStatus) == accountdomain.StatusActivated {
		newStatus = accountdomain.StatusActivated
	} else if accountdomain.Enum(dto.NewStatus) == accountdomain.StatusBanned {
		newStatus = accountdomain.StatusBanned
	} else {
		return common.NewInternalServerError().
			WithReason("new status for this account is invalid")
	}
	if err := h.commandrepo.UpdateStatus(ctx, accId, newStatus); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update this account status to " + dto.NewStatus).
			WithInner(err.Error())
	}
	return nil
}
