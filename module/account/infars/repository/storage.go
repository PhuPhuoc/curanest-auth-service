package accountrepository

import "github.com/jmoiron/sqlx"

type accountRepo struct {
	db *sqlx.DB
}

func NewAccountRepo(db *sqlx.DB) *accountRepo {
	return &accountRepo{
		db: db,
	}
}
