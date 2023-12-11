package handler

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/common"
	notifService "github.com/trikrama/Depublic/internal/app/notification/service"
	"github.com/trikrama/Depublic/internal/app/user/entity"
	"github.com/trikrama/Depublic/internal/app/user/service"
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/validator"
)

type UserHandler struct {
	userService  service.UserServiceInterface
	notifService notifService.NotificationServiceInterface
}

func NewUserHandler(cfg *config.Config, userService service.UserServiceInterface, notifService notifService.NotificationServiceInterface) *UserHandler {
	return &UserHandler{
		userService:  userService,
		notifService: notifService,
	}
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.userService.GetAllUser(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": users,
	})
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*common.JwtCustomClaims)
	if claims.Role != "Admin" {
		if int64(idInt) != claims.ID {
			return c.JSON(http.StatusForbidden, echo.Map{
				"message": "You don't have access to this user",
			})
		}
	}
	user, err := h.userService.GetUserByID(c.Request().Context(), idInt)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	user := entity.UserRequestUpdate{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*common.JwtCustomClaims)
	if claims.Role != "Admin" {
		if user.ID != claims.ID {
			return c.JSON(http.StatusForbidden, echo.Map{
				"message": "You don't have access to this user",
			})
		}
	}
	userRequest := entity.NewUserUpdate(user)
	userUpdate, err := h.userService.UpdateUser(c.Request().Context(), userRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	userUpdate.Role = "Buyer"
	userResponse := entity.NewUserResponse(userUpdate)
	return c.JSON(http.StatusOK, echo.Map{"data": userResponse})
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	if err := h.userService.DeleteUser(c.Request().Context(), idInt); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success delete user",
	})
}
