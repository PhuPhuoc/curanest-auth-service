package accountdomain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	id          uuid.UUID
	roleId      uuid.UUID
	fullName    string
	phoneNumber string
	email       string
	password    string
	salt        string
	avatar      string
	status      Status
	createdAt   *time.Time
}

func (a *Account) GetID() uuid.UUID {
	return a.id
}

func (a *Account) GetRoleID() uuid.UUID {
	return a.roleId
}

func (a *Account) GetPhoneNumber() string {
	return a.phoneNumber
}

func (a *Account) GetFullName() string {
	return a.fullName
}

func (a *Account) GetEmail() string {
	return a.email
}

func (a *Account) GetPassword() string {
	return a.password
}

func (a *Account) GetSalt() string {
	return a.salt
}

func (a *Account) GetAvatar() string {
	return a.avatar
}

func (a *Account) GetStatus() Status {
	return a.status
}

func (a *Account) GetCreatedAt() time.Time {
	return *a.createdAt
}

func NewAccount(id, roleId uuid.UUID, fullName, phoneNumber, email, password, salt, avatar string, status Status, createdAt *time.Time) (*Account, error) {
	return &Account{
		id:          id,
		roleId:      roleId,
		fullName:    fullName,
		phoneNumber: phoneNumber,
		email:       email,
		password:    password,
		salt:        salt,
		avatar:      avatar,
		status:      status,
		createdAt:   createdAt,
	}, nil
}

type Status int

const (
	StatusActivated Status = iota
	StatusBanned
)

func (r Status) String() string {
	switch r {
	case StatusActivated:
		return "activated"
	case StatusBanned:
		return "banned"
	default:
		return "unknown"
	}
}

func Enum(s string) Status {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "activated":
		return StatusActivated
	default:
		return StatusBanned
	}
}
