package model

import "github.com/iamsad5566/member_service_frame/repo"

func CreateTable(usrRepo repo.UserRepoInterface) (bool, error) {
	return usrRepo.CreateTable()
}
