package roledomain

import "github.com/google/uuid"

type Role struct {
	id   uuid.UUID
	name string
}

func NewRole(id uuid.UUID, name string) (*Role, error) {
	return &Role{
		id:   id,
		name: name,
	}, nil
}

func (r *Role) GetID() uuid.UUID {
	return r.id
}

func (r *Role) GetName() string {
	return r.name
}
