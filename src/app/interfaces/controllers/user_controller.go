package controllers

import (
	"github.com/yosukei3108/LearnCleanArchitecture/src/app/domain"
	"github.com/yosukei3108/LearnCleanArchitecture/src/app/interfaces/database"
	"github.com/yosukei3108/LearnCleanArchitecture/src/app/usecase"
	"strconv"
)

type UserController struct {
	interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Ingeractor: usecase.UserIngeractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(201)
}

func (controller *UserController) Index(c Context) {
	users err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user,err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	c.JSON(299, user)
}
