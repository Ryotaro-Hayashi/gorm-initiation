package database

// usecaseで定義したinterfaceの実装はここではないので、usecaseをimportする必要はない（？）
import (
	"gorm-initiation/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repository *UserRepository) FindAll() (users domain.Users, err error){
	users, err = repository.SqlHandler.Find()
	return
}