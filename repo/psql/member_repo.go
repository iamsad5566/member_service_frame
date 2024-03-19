package psql

import "member_service_frame/repo"

type PsqlUserRepository struct {
	client PsqlClientInterface
}

func NewUserRepository(client PsqlClientInterface) *PsqlUserRepository {
	return &PsqlUserRepository{client: client}
}

var _ repo.UserRepository = (*PsqlUserRepository)(nil)
