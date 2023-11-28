package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/internal/app/user"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func NewUserHandler(userService user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.userService.GetAllUser(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	user, err := h.userService.GetUserByID(c.Request().Context(), idInt)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	userRequest := new(user.UserRequest)
	userNewRequest := user.NewUser(*userRequest)
	if err := c.Bind(&userRequest); err != nil {
		return err
	}
	if err := h.userService.CreateUser(c.Request().Context(), userNewRequest); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, userNewRequest)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {

	userRequest := new(user.UserRequest)
	if err := c.Bind(&userRequest); err != nil {
		return err
	}
	if err := h.userService.UpdateUser(c.Request().Context(), userRequest); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, userRequest)
}
