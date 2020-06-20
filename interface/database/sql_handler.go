package database

import (
	"gorm-initiation/domain"
)

type SqlHandler interface {
	Find() (users domain.Users, err error)
}