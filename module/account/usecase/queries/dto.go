package accountqueries

import (
	"time"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

type LoginByPhoneRequestDTO struct {
	PhoneNumber string `json:"phone-number"`
	Password    string `json:"password"`
	PushToken   string `json:"push-token"`
}

type LoginByEmailRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	AccountInfo AccountLoginInfo `json:"account-info"`
	Token       TokenReponseDTO  `json:"token"`
}

type AccountLoginInfo struct {
	Id          uuid.UUID `json:"id"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	Role        string    `json:"role"`
}
type TokenReponseDTO struct {
	AccessToken      string `json:"access_token"`
	AccessTokenExpIn int    `json:"access_token_exp_in"`
	// RefreshToken      string `json:"refresh_token"`
	// RefreshTokenExpIn int    `json:"refresh_token_exp_in"`
}

type MyAccountDTO struct {
	Id          uuid.UUID `json:"id"`
	RoleId      uuid.UUID `json:"-"`
	Role        string    `json:"role"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	CreatedAt   time.Time `json:"created-at"`
}

func toMyAccDTO(data *accountdomain.Account) *MyAccountDTO {
	dto := &MyAccountDTO{
		Id:          data.GetID(),
		RoleId:      data.GetRoleID(),
		FullName:    data.GetFullName(),
		Email:       data.GetEmail(),
		PhoneNumber: data.GetPhoneNumber(),
		CreatedAt:   data.GetCreatedAt(),
	}
	return dto
}

type AccountIdsQuery struct {
	Role string      `json:"role"`
	Ids  []uuid.UUID `json:"ids"`
}

type FilterAccountQuery struct {
	Paging common.Paging      `json:"paging"`
	Filter FieldFilterAccount `json:"filter"`
}

type FieldFilterAccount struct {
	RoleId      string `json:"-"`
	Role        string `form:"role" json:"role"`
	FullName    string `form:"full-name" json:"full-name"`
	Email       string `form:"email" json:"email"`
	PhoneNumber string `form:"phone-number" json:"phone-number"`
}

type AccountDTO struct {
	Id          uuid.UUID `json:"id"`
	RoleId      uuid.UUID `json:"-"`
	FullName    string    `json:"full-name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone-number"`
	CreatedAt   time.Time `json:"created-at"`
}

func toDTO(data *accountdomain.Account) AccountDTO {
	dto := AccountDTO{
		Id:          data.GetID(),
		RoleId:      data.GetRoleID(),
		FullName:    data.GetFullName(),
		Email:       data.GetEmail(),
		PhoneNumber: data.GetPhoneNumber(),
		CreatedAt:   data.GetCreatedAt(),
	}
	return dto
}

type RegisPushToken struct {
	AccountId uuid.UUID `json:"account-id"`
	PushToken string    `json:"push-token"`
}
