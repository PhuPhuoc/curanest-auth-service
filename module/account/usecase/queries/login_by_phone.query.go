package accountqueries

import (
	"context"
	"errors"

	"github.com/PhuPhuoc/curanest-auth-service/common"
)

type loginByPhonePasswordHandler struct {
	queryRepo     AccountQueryRepo
	tokenProvider TokenProvider
	roleFetcher   RoleFetcher
}

func NewLoginWithPhoneHandler(queryRepo AccountQueryRepo, tokenProvider TokenProvider, roleFetcher RoleFetcher) *loginByPhonePasswordHandler {
	return &loginByPhonePasswordHandler{
		queryRepo:     queryRepo,
		tokenProvider: tokenProvider,
		roleFetcher:   roleFetcher,
	}
}

func (h *loginByPhonePasswordHandler) Handle(ctx context.Context, req *LoginByPhoneRequestDTO) (*LoginResponseDTO, error) {
	entityFound, err := h.queryRepo.FindByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			return nil, common.NewBadRequestError().WithReason("this phone number does not exist")
		} else {
			return nil, common.NewInternalServerError().
				WithReason("cannot get entity from db").
				WithInner(err.Error())
		}
	}

	if isTrue := common.CompareHashPassword(entityFound.GetPassword(), entityFound.GetSalt(), req.Password); !isTrue {
		return nil, common.NewBadRequestError().WithReason("wrong password")
	}

	var role string
	if role, err = h.roleFetcher.GetNameByRoleId(ctx, entityFound.GetRoleID()); err != nil {
		return nil, err
	}

	// gen token
	tokenId := common.GenUUID()
	accId := entityFound.GetID()
	accessToken, err := h.tokenProvider.IssueToken(ctx, tokenId.String(), accId.String(), role)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot create access token").
			WithInner(err.Error())
	}

	response := &LoginResponseDTO{
		AccountInfo: AccountLoginInfo{
			Id:          entityFound.GetID(),
			FullName:    entityFound.GetFullName(),
			Email:       entityFound.GetEmail(),
			PhoneNumber: entityFound.GetPhoneNumber(),
			Role:        role,
		},
		Token: TokenReponseDTO{
			AccessToken:      accessToken,
			AccessTokenExpIn: h.tokenProvider.TokenExpireInSeconds(),
		},
	}

	return response, nil
}
