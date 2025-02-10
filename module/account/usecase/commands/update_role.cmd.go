package accountcommands

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

type updateAccountRoleHandler struct {
	commandrepo AccountCommandRepo
	rolefetcher RoleFetcher
}

func NewUpdateAccountRoleHandler(cmdRepo AccountCommandRepo, roleFetch RoleFetcher) *updateAccountRoleHandler {
	return &updateAccountRoleHandler{
		commandrepo: cmdRepo,
		rolefetcher: roleFetch,
	}
}

type UpdateRoleRequest struct {
	Role string `json:"role"`
}

func (h *updateAccountRoleHandler) Handle(ctx context.Context, nurseId *uuid.UUID, payload *UpdateRoleRequest) error {
	if payload.Role == "admin" || payload.Role == "relatives" {
		return common.NewBadRequestError().
			WithReason("cannot change to this role '" + payload.Role + "'")
	}

	roleid, err := h.rolefetcher.GetRoleIdByName(ctx, payload.Role)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return common.NewBadRequestError().
				WithReason("no role with name '" + payload.Role + "'")
		}
	}

	if err := h.commandrepo.UpdateRoleForNurseAndStaff(ctx, nurseId, roleid); err != nil {
		if errors.Is(err, common.ErrNoRecordsAreChanged) {
			return common.NewInternalServerError().
				WithReason("can't update role because account can't be found").
				WithInner(err.Error())
		}
		return common.NewInternalServerError().
			WithReason("cannot update this role account to: " + payload.Role).
			WithInner(err.Error())
	}

	return nil
}
