package psql

import (
	"database/sql"

	"github.com/iamsad5566/member_service_frame/repo"
)

type PsqlUserRepository struct {
	client PsqlClientInterface
}

type PsqlClientInterface interface {
	Begin() (*sql.Tx, error)
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func NewUserRepository(client PsqlClientInterface) *PsqlUserRepository {
	return &PsqlUserRepository{client: client}
}

var _ repo.UserRepoInterface = (*PsqlUserRepository)(nil)
