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
// routerでNewUserControllerを引数に渡すことからメモリ削減の観点でUserControllerをポインタで返す
func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController {
		UserInteractor: usecase.UserInteractor {
			UserRepository: &database.UserRepository {
				SqlHandler: sqlHandler,
			},
		},
	}
}

// *gin.Contextは、DIPで用意
func (controller *UserController) Index(c Context) {
	users, err := controller.UserInteractor.GetUsers()
	if err != nil {
		c.JSON(500, err)
		return
	}
	
	c.JSON(200, users)
}