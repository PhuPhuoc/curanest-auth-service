package accountqueries

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

type validateAccountEmailHandler struct {
	queryRepo AccountQueryRepo
}

func NewValidateAccountEmailHandler(queryRepo AccountQueryRepo) *validateAccountEmailHandler {
	return &validateAccountEmailHandler{
		queryRepo: queryRepo,
	}
}

func (h *validateAccountEmailHandler) Handle(ctx context.Context, email string) error {
	entityFound, err := h.queryRepo.FindByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, common.ErrRecordNotFound) {
			return common.NewInternalServerError().
				WithReason("unable to perform the step of checking email exists in the system").
				WithInner(err.Error())
		}
	}
	if entityFound != nil {
		return common.NewBadRequestError().WithReason("Email already exists: " + email)
	}
	return nil
}
