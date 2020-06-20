package usecase

import (
	"gorm-initiation/domain"
)

type UserRepository interface {
	FindAll() (users domain.Users, err error)
}