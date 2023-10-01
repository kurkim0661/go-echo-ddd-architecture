package command

import (
	"myapp/api/application"
	"myapp/api/domain"
	"myapp/api/domain/dto"
	"myapp/api/interfaces/client/command/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	application.UserService
}

func (UserController) NewUserController(service application.UserService) *UserController {
	return &UserController{service}
}

func (userController *UserController) Init(e *echo.Group) {
	e.POST("", userController.CreateUser)
	e.GET("", userController.GetUsers)
	e.DELETE("/:id", userController.DeleteUser)
	e.GET("/:id", userController.GetUser)
	e.PATCH("", userController.UpdateUser)
}

// CreateUser is create new user
// @Summary Create user
// @Description Create new user passing name parameter
// @Accept json
// @Produce json
// @Param user body UserDto true "body of the user"
// @Success 201 {object} ApiResult{result=model.User}
// @Failure 500 {object} ApiResult{result=model.User} "Internal Server Error"
// @Router /users [post]
func (userController *UserController) CreateUser(c echo.Context) error {
	userDto := &dto.UserDto{}
	bindErr := c.Bind(userDto)
	user := userDto.ToEntity()

	if bindErr != nil {

		c.Logger().Error(bindErr)
		return response.ReturnApiFail(c, http.StatusBadRequest, response.ApiParameterError, bindErr)
	}
	createUser, err := userController.UserService.CreateUser(user)
	if err != nil {
		c.Logger().Error(err)
		return response.ReturnApiFail(c, http.StatusInternalServerError, response.ApiQueryError, err)
	}
	c.Logger().Info(createUser)
	return response.ReturnApiSuccess(c, http.StatusCreated, createUser)
}

// GetUsers get all users' list
// @Summary Get all users
// @Description Get all user's info
// @Accept json
// @Produce json
// @Success 200 {object} ApiResult{result=model.User}
// @Router /users [get]
func (userController *UserController) GetUsers(c echo.Context) error {
	users, err := userController.UserService.GetUsers()
	if err != nil {
		c.Logger().Error(err)
		return response.ReturnApiFail(c, http.StatusInternalServerError, response.ApiQueryError, err)
	}
	c.Logger().Info(users)
	return response.ReturnApiSuccess(c, http.StatusOK, users)
}

// DeleteUser delete specific user's info
// @Summary Delete user
// @Description Delete existing user's info passing id parameter
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 204 {object} ApiResult{result=model.User}
// @Router /users/{id} [delete]
func (userController *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return response.ReturnApiFail(c, http.StatusBadRequest, response.ApiParameterError, err)
	}

	err = userController.UserService.DeleteUser(id)
	if err != nil {
		c.Logger().Error(err)
		return response.ReturnApiFail(c, http.StatusInternalServerError, response.ApiQueryError, err)
	}
	return response.ReturnApiSuccess(c, http.StatusNoContent, nil)
}

// GetUser get user's info using id
// @Summary Get user
// @Description Get user's info passing id parameter
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 200 {object} ApiResult{result=model.User}
// @Failure 500 {object} ApiResult{result=model.User} "Internal Server Error"
// @Router /users/{id} [get]
func (userController *UserController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return response.ReturnApiFail(c, http.StatusBadRequest, response.ApiParameterError, err)
	}

	user, err := userController.UserService.GetUser(id)
	if err != nil {
		c.Logger().Error(err)
		return response.ReturnApiFail(c, http.StatusInternalServerError, response.ApiQueryError, err)
	}
	c.Logger().Info(user)
	return response.ReturnApiSuccess(c, http.StatusOK, user)
}

// UpdateUser updates existing user's info
// @Summary Update user
// @Description Update existing user's information
// @Accept json
// @Produce json
// @Param name body model.User true "body of the user"
// @Success 201 {object} ApiResult{result=model.User}
// @Router /users [patch]
func (userController *UserController) UpdateUser(c echo.Context) error {
	user := &domain.User{}
	bindErr := c.Bind(user)
	if bindErr != nil {
		c.Logger().Error(bindErr)
		return response.ReturnApiFail(c, http.StatusBadRequest, response.ApiParameterError, bindErr)
	}

	createUser, err := userController.UserService.UpdateUser(user)
	if err != nil {
		c.Logger().Error(bindErr)
		return response.ReturnApiFail(c, http.StatusInternalServerError, response.ApiQueryError, err)
	}
	c.Logger().Info(createUser)
	return response.ReturnApiSuccess(c, http.StatusOK, user)
}
