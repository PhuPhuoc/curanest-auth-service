package rolerepository

import (
	roledomain "github.com/PhuPhuoc/curanest-auth-service/module/role/domain"
	"github.com/google/uuid"
)

const (
	table   = `roles`
	field   = `id, name`
	mapping = `:id, :name`
)

type RoleDTO struct {
	Id   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (dto *RoleDTO) ToEntity() (*roledomain.Role, error) {
	return roledomain.NewRole(
		dto.Id,
		dto.Name,
	)
}

func ToDTO(data *roledomain.Role) *RoleDTO {
	dto := &RoleDTO{
		Id:   data.GetID(),
		Name: data.GetName(),
	}
	return dto
}
