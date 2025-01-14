package rolequeries

import (
	"context"

	"github.com/PhuPhuoc/curanest-auth-service/common"
	roledomain "github.com/PhuPhuoc/curanest-auth-service/module/role/domain"
	"github.com/google/uuid"
)

type getRolesHandler struct {
	queryRepo RoleQueryRepo
}

func NewGetRolesHandler(queryRepo RoleQueryRepo) *getRolesHandler {
	return &getRolesHandler{
		queryRepo: queryRepo,
	}
}

type RoleDTO struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func toDTO(data *roledomain.Role) RoleDTO {
	dto := RoleDTO{
		Id:   data.GetID(),
		Name: data.GetName(),
	}
	return dto
}

func (h *getRolesHandler) Handle(ctx context.Context) ([]RoleDTO, error) {
	entities, err := h.queryRepo.GetRoles(ctx)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithMessage("cannot get list role").
			WithReason("error at GetRoles-repo").
			WithInner(err.Error())
	}

	list_dto := make([]RoleDTO, len(entities))
	for i := range entities {
		list_dto[i] = toDTO(&entities[i])
	}

	return list_dto, nil
}
