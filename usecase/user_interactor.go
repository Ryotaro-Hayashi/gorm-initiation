package usecase

import (
	"gorm-initiation/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) GetUsers() (users domain.Users, err error){
	users, err = interactor.UserRepository.FindAll()
	
	return
}