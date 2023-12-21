package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/internal/app/event/entity"
	"github.com/trikrama/Depublic/internal/app/event/service"
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/validator"
)

type EventHandler struct {
	eventService service.EventServiceInterface
}

func NewEventHandler(cfg *config.Config, eventService service.EventServiceInterface) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}

func (h *EventHandler) GetAllEvents(c echo.Context) error {
	query := entity.QueryFilter{}
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	events, err := h.eventService.GetAllEvent(c.Request().Context(), query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"events": events,
	})
}

func (h *EventHandler) GetEventByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	event, err := h.eventService.GetEventByID(c.Request().Context(), int64(idInt))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"event": event,
	})
}

func (h *EventHandler) CreateEvent(c echo.Context) error {
	eventRequest := entity.EventRequest{}
	if err := c.Bind(&eventRequest); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	eventNewRequest := entity.NewEvent(eventRequest)
	if err := h.eventService.CreateEvent(c.Request().Context(), eventNewRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	eventResponse := entity.NewEventResponse(eventNewRequest)
	return c.JSON(http.StatusCreated, echo.Map{
		"event": eventResponse,
	})
}

func (h *EventHandler) UpdateEvent(c echo.Context) error {
	event := entity.EventRequestUpdate{}
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	eventRequest := entity.NewEventUpdate(event)
	if err := h.eventService.UpdateEvent(c.Request().Context(), eventRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	eventResponse := entity.NewEventResponse(eventRequest)
	return c.JSON(http.StatusOK, echo.Map{
		"event": eventResponse,
	})
}

func (h *EventHandler) DeleteEvent(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	if err := h.eventService.DeleteEvent(c.Request().Context(), idInt); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "event deleted",
	})
}
