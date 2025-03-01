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

func (h *updateAccountRoleHandler) Handle(ctx context.Context, nurseId *uuid.UUID, targetRole string) error {
	if targetRole == "admin" || targetRole == "relatives" {
		return common.NewBadRequestError().
			WithReason("cannot change to this role '" + targetRole + "'")
	}

	roleid, err := h.rolefetcher.GetRoleIdByName(ctx, targetRole)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return common.NewBadRequestError().
				WithReason("no role with name '" + targetRole + "'")
		}
	}

	if err := h.commandrepo.UpdateRoleForNurseAndStaff(ctx, nurseId, roleid); err != nil {
		if errors.Is(err, common.ErrNoRecordsAreChanged) {
			return common.NewInternalServerError().
				WithReason("can't update role because account can't be found").
				WithInner(err.Error())
		}
		return common.NewInternalServerError().
			WithReason("cannot update this role account to: " + targetRole).
			WithInner(err.Error())
	}

	return nil
}
