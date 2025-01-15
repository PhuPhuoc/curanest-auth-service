package accountcommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

type CreateAccountCmdDTO struct {
	RoleName    string `json:"role-name"`
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type createAccountHandler struct {
	commandrepo AccountCommandRepo
	rolefetcher RoleFetcher
}

func NewCreateAccountHandler(cmdRepo AccountCommandRepo, roleFetch RoleFetcher) *createAccountHandler {
	return &createAccountHandler{
		commandrepo: cmdRepo,
		rolefetcher: roleFetch,
	}
}

func (h *createAccountHandler) Handle(ctx context.Context, dto *CreateAccountCmdDTO) error {
	//  generate salt
	var salt, hashedPassword string
	var err error
	if salt, err = common.RandomStr(30); err != nil {
		return err
	}

	//  hash password + salt
	if hashedPassword, err = common.HashPassword(salt, dto.Password); err != nil {
		return err
	}

	// get roleid
	var roleid *uuid.UUID
	if roleid, err = h.rolefetcher.GetRoleIdByName(ctx, dto.RoleName); err != nil {
		return err
	}

	accid := common.GenUUID()
	entity, _ := accountdomain.NewAccount(
		accid,
		*roleid,
		dto.FullName,
		dto.PhoneNumber,
		dto.Email,
		hashedPassword,
		salt,
		accountdomain.StatusActivated,
		nil,
		nil,
	)

	if err = h.commandrepo.Create(ctx, entity); err != nil {
		return err
	}
	return nil
}
