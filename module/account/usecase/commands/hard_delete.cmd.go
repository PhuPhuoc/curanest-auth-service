package accountcommands

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

type hardDeleteAccountHandler struct {
	commandrepo AccountCommandRepo
}

func NewHardDeleteAccountHandler(cmdRepo AccountCommandRepo) *hardDeleteAccountHandler {
	return &hardDeleteAccountHandler{
		commandrepo: cmdRepo,
	}
}

func (h *hardDeleteAccountHandler) Handle(ctx context.Context, accountId *uuid.UUID) error {
	if err := h.commandrepo.HardDelete(ctx, accountId); err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return common.NewBadRequestError().
				WithReason("cannot found account with id: " + accountId.String()).
				WithInner(err.Error())
		}
		return common.NewInternalServerError().
			WithReason("cannot delete account with id: " + accountId.String()).
			WithInner(err.Error())
	}
	return nil
}
