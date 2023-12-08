package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/common"
	"github.com/trikrama/Depublic/internal/app/user/entity"
	"github.com/trikrama/Depublic/internal/app/user/service"
	notifService "github.com/trikrama/Depublic/internal/app/notification/service"
	notifEntity "github.com/trikrama/Depublic/internal/app/notification/entity"
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/validator"
)

type UserHandler struct {
	userService service.UserServiceInterface
	notifService notifService.NotificationServiceInterface
}

func NewUserHandler(cfg *config.Config, userService service.UserServiceInterface, notifService notifService.NotificationServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
		notifService: notifService,
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
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*common.JwtCustomClaims)
	if claims.Role != "Admin" {
		if int64(idInt) != claims.ID {
			return c.JSON(http.StatusForbidden, echo.Map{
				"message": "Anda tidak punya akses untuk akun ini",
			})
		}
	}
	user, err := h.userService.GetUserByID(c.Request().Context(), idInt)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
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
		Title: "Buat Akun",
		Body: fmt.Sprintf("User Dengan Nama: %s, Telah Berhasil Membuat Akun", userNewRequest.Name,),
		Status: "Berhasil",
	}
	newNotif := notifEntity.NewNotification(*notif)
	err := h.notifService.CreateNotification(c.Request().Context(), newNotif)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, userNewRequest)
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
				"message": "Anda tidak punya akses untuk akun ini",
			})
		}
	}
	userRequest := entity.NewUserUpdate(user)
	userUpdate, err := h.userService.UpdateUser(c.Request().Context(), userRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	userResponse := entity.NewUserResponse(userUpdate)
	return c.JSON(http.StatusOK, userResponse)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	if err := h.userService.DeleteUser(c.Request().Context(), idInt); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "user deleted",
	})
}

func (h *UserHandler) LoginUser(c echo.Context) error {
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
