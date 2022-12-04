package repository

import (
	"database/sql"
)

//HandlerManager constructs a new ReplaceManager
type ReplaceManager interface {
}

func NewReplaceManager(repository Type) ReplaceManager {
	switch repository {
	case PostgresSQL:
		return &replaceManager{DB: NewSQLConnection()}
	}

	return nil
}

type replaceManager struct {
	*sql.DB
}
