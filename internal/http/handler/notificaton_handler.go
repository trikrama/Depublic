package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/internal/app/notification/entity"
	"github.com/trikrama/Depublic/internal/app/notification/service"
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/validator"
)

type NotificationHandler struct {
	notificationService service.NotificationServiceInterface
}

func NewNotificationHandler(cfg *config.Config, notificationService service.NotificationServiceInterface) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

func (h *NotificationHandler) GetUserNotifications(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	notifications, err := h.notificationService.GetByUser(c.Request().Context(), int64(idInt))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	var responseNotif []entity.NotificationResponse
	for _, notif := range notifications {
		responseNotif = append(responseNotif, *entity.NewNotificationResponse(notif))
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": responseNotif,
	})
}

func (h *NotificationHandler) GetAllNotifications(c echo.Context) error {
	notifications, err := h.notificationService.GetAllNotifications(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	var responseNotif []entity.NotificationResponse
	for _, notif := range notifications {
		responseNotif = append(responseNotif, *entity.NewNotificationResponse(notif))
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": responseNotif,
	})
}

func (h *NotificationHandler) UpdateNotification(c echo.Context) error {
	notifUpdate := entity.NotificationRequestUpdate{}
	if err := c.Bind(&notifUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	newNotif := entity.NewNotificationUpdate(notifUpdate)
	err := h.notificationService.UpdateNotification(c.Request().Context(), newNotif)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Notification Updated",
	})
}

func (h *NotificationHandler) DeleteNotification(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	err := h.notificationService.DeleteNotification(c.Request().Context(), idInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Notification Deleted",
	})
}
