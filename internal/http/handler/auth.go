package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	notifEntity "github.com/trikrama/Depublic/internal/app/notification/entity"
	notifService "github.com/trikrama/Depublic/internal/app/notification/service"
	"github.com/trikrama/Depublic/internal/app/user/entity"
	"github.com/trikrama/Depublic/internal/app/user/service"
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/validator"
)

type AuthHandler struct {
	userService service.UserServiceInterface
	notifService notifService.NotificationServiceInterface
}

func NewAuthHandler(cfg *config.Config, userService service.UserServiceInterface, notifService notifService.NotificationServiceInterface) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		notifService: notifService,
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	userRequest := entity.UserRequest{}
	if err := c.Bind(&userRequest); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	userNewRequest := entity.NewUser(userRequest)
	if err := h.userService.CreateUser(c.Request().Context(), userNewRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	notif := &notifEntity.NotificationRequest{
		UserID: userNewRequest.ID,
		Title:  "Buat Akun",
		Body:   fmt.Sprintf("User Dengan Nama: %s, Telah Berhasil Membuat Akun", userNewRequest.Name),
		Status: "Berhasil",
	}
	newNotif := notifEntity.NewNotification(*notif)
	err := h.notifService.CreateNotification(c.Request().Context(), newNotif)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	userResponse := entity.NewUserResponse(userNewRequest)
	return c.JSON(http.StatusCreated, echo.Map{
		"data": userResponse,
	})
}

func (h *AuthHandler)Login(c echo.Context) error {
	login := entity.UserRequestLogin{}
	if err := c.Bind(&login); err != nil {
		fmt.Println("salah di handler bind")
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	user, token, err := h.userService.LoginUser(c.Request().Context(), login.Email, login.Password)
	if err != nil {
		fmt.Println("salah di handler panggil login user")
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user":  user,
		"token": token,
	})
}
