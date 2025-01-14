package rolerepository

import "github.com/jmoiron/sqlx"

type roleRepo struct {
	db *sqlx.DB
}

func NewRoleRepo(db *sqlx.DB) *roleRepo {
	return &roleRepo{
		db: db,
	}
}
