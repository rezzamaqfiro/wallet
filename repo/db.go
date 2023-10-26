package repo

import (
	"database/sql"

	repo "github.com/rezzamaqfiro/wallet/repo/generated"
)

type Queries struct {
	*repo.Queries
	db repo.DBTX
}

func New(db *sql.DB) *Queries {
	return &Queries{repo.New(db), db}
}
