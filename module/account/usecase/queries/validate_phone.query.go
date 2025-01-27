package accountqueries

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

type validateAccountPhoneNumberHandler struct {
	queryRepo AccountQueryRepo
}

func NewVerifyPhoneHandler(queryRepo AccountQueryRepo) *validateAccountPhoneNumberHandler {
	return &validateAccountPhoneNumberHandler{
		queryRepo: queryRepo,
	}
}

func (h *validateAccountPhoneNumberHandler) Handle(ctx context.Context, phoneNumber string) error {
	entityFound, err := h.queryRepo.FindByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		if !errors.Is(err, common.ErrRecordNotFound) {
			return common.NewInternalServerError().
				WithReason("unable to perform the step of checking phone number exists in the system").
				WithInner(err.Error())
		}
	}
	if entityFound != nil {
		return common.NewBadRequestError().WithReason("Phone Number already exists: " + phoneNumber)
	}
	return nil
}
