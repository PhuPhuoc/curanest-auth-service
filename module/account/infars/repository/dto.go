package accountrepository

import (
	"time"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

var (
	TABLE        = `accounts`
	FIELD        = []string{"id", "role_id", "full_name", "phone_number", "email", "password", "salt", "avatar", "status"}
	GET_FIELD    = []string{"id", "role_id", "full_name", "phone_number", "email", "status", "avatar", "created_at"}
	UPDATE_FIELD = []string{"full_name", "phone_number", "email", "avatar"}
)

type AccountDTO struct {
	Id          uuid.UUID  `db:"id"`
	RoleId      uuid.UUID  `db:"role_id"`
	FullName    string     `db:"full_name"`
	PhoneNumber string     `db:"phone_number"`
	Email       string     `db:"email"`
	Password    string     `db:"password"`
	Salt        string     `db:"salt"`
	Avatar      string     `db:"avatar"`
	Status      string     `db:"status"`
	CreatedAt   *time.Time `db:"created_at"`
}

func (dto *AccountDTO) ToEntity() (*accountdomain.Account, error) {
	return accountdomain.NewAccount(
		dto.Id,
		dto.RoleId,
		dto.FullName,
		dto.PhoneNumber,
		dto.Email,
		dto.Password,
		dto.Salt,
		dto.Avatar,
		accountdomain.Enum(dto.Status),
		dto.CreatedAt,
	)
}

func ToDTO(data *accountdomain.Account) *AccountDTO {
	dto := &AccountDTO{
		Id:          data.GetID(),
		RoleId:      data.GetRoleID(),
		FullName:    data.GetFullName(),
		PhoneNumber: data.GetPhoneNumber(),
		Email:       data.GetEmail(),
		Password:    data.GetPassword(),
		Salt:        data.GetSalt(),
		Avatar:      data.GetAvatar(),
		Status:      data.GetStatus().String(),
	}
	return dto
}
