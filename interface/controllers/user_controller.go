package controllers

import (
	"gorm-initiation/usecase"
	"gorm-initiation/interface/database"
)

type UserController struct {
	UserInteractor usecase.UserInteractor
}

// database.SqlHandlerはinterfaceでありポインタ型
// database.UserRepositoryもinterfaceがポインタなのでポインタにしなければならない
func NewUserController(sqlHandler database.SqlHandler) UserController {
	return UserController {
		UserInteractor: usecase.UserInteractor {
			UserRepository: &database.UserRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

// *gin.Contextは、DIPで用意
func (controller *UserController) Index(c Context) {
	users := controller.UserInteractor.GetUsers()
	
	c.JSON(200, users)
}