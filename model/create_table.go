package model

import "member_service_frame/repo"

func CreateTable(usrRepo repo.UserRepoInterface) (bool, error) {
	return usrRepo.CreateTable()
}
