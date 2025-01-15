package accountrepository

import (
	"time"

	accountdomain "github.com/PhuPhuoc/curanest-auth-service/module/account/domain"
	"github.com/google/uuid"
)

const (
	table   = `accounts`
	field   = `id, role_id, full_name, phone_number, email, password, salt, status`
	mapping = `:id, :role_id, :full_name, :phone_number, :email, :password, :salt, :status`

	getField   = `id, role_id, full_name, phone_number, email, password, salt, status, created_at, updated_at`
	getMapping = `:id, :role_id, :full_name, :phone_number, :email, :password, :salt, :status, :created_at, :updated_at`
)

type AccountDTO struct {
	Id          uuid.UUID  `db:"id"`
	RoleId      uuid.UUID  `db:"role_id"`
	FullName    string     `db:"full_name"`
	PhoneNumber string     `db:"phone_number"`
	Email       string     `db:"email"`
	Password    string     `db:"password"`
	Salt        string     `db:"salt"`
	Status      string     `db:"status"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
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
		accountdomain.Enum(dto.Status),
		dto.CreatedAt,
		dto.UpdatedAt,
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
		Status:      data.GetStatus().String(),
	}
	return dto
}
