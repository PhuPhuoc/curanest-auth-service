package common

import "github.com/google/uuid"

type contextKey string

const (
	KeyRequester contextKey = "requester"
	KeyToken     contextKey = "token"
)

type Requester interface {
	UserId() uuid.UUID
	TokenId() uuid.UUID
	Role() string
}
type requesterData struct {
	sub  uuid.UUID
	id   uuid.UUID
	role string
}

func (r *requesterData) UserId() uuid.UUID {
	return r.sub
}

func (r *requesterData) TokenId() uuid.UUID {
	return r.id
}

func (r *requesterData) Role() string {
	return r.role
}

func NewRequester(sub, tid uuid.UUID, role string) Requester {
	return &requesterData{
		sub:  sub,
		id:   tid,
		role: role,
	}
}
