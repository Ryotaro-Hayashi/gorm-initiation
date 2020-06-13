package usecase

import (
	"gorm-initiation/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) GetUsers() (users domain.Users){
	users = interactor.UserRepository.FindAll()
	
	return
}