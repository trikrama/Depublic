package handler

import (
	"net/http"
	"strconv"
	"time"

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
	events, err := h.eventService.GetAllEvent(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, events)
}

func (h *EventHandler) GetEventByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	event, err := h.eventService.GetEventByID(c.Request().Context(), idInt)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, event)
}

func (h *EventHandler) CreateEvent(c echo.Context) error {
	eventRequest := entity.EventRequest{}
	if err := c.Bind(&eventRequest); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	eventNewRequest := entity.NewEvent(eventRequest)
	if err := h.eventService.CreateEvent(c.Request().Context(), eventNewRequest); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, eventNewRequest)
}

func (h *EventHandler) UpdateEvent(c echo.Context) error {
	event := entity.EventRequestUpdate{}
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	eventRequest := entity.NewEventUpdate(event)
	if err := h.eventService.UpdateEvent(c.Request().Context(), eventRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, eventRequest)
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

func (h *EventHandler) SearchEvent(c echo.Context) error {
	var inputSearch struct {
		KeyWord string `param:"keyword" validate:"required"`
	}

	if err := c.Bind(&inputSearch); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	events, err := h.eventService.SearchEvent(c.Request().Context(), inputSearch.KeyWord)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, events)
}

func (h *EventHandler) FilterEventByPrice(ctx echo.Context) error {
	var inputFilter struct {
		Min string `param:"min"`
		Max string `param:"max"`
	}

	if err := ctx.Bind(&inputFilter); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	events, err := h.eventService.FilterEventByPrice(ctx.Request().Context(), inputFilter.Min, inputFilter.Max)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, events)
}

func (h *EventHandler) FilterEventByDate(ctx echo.Context) error {
	var inputFilter struct {
		StartDate string `param:"start_date"`
		EndDate   string `param:"end_date"`
	}

	if err := ctx.Bind(&inputFilter); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	startDate, err := time.Parse("2006-01-02", inputFilter.StartDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	endDate, err := time.Parse("2006-01-02", inputFilter.EndDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	events, err := h.eventService.FilterEventByDate(ctx.Request().Context(), startDate, endDate)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, events)
}

func (h *EventHandler) FilterEventByLocation(ctx echo.Context) error {
	var input struct {
		Location string `param:"location"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	event, err := h.eventService.FilterEventByLocation(ctx.Request().Context(), input.Location)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) FilterEventByStatus(ctx echo.Context) error {
	var input struct {
		Status string `param:"status"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	event, err := h.eventService.SortEventByStatus(ctx.Request().Context(), input.Status)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) SortEventByExpensive(ctx echo.Context) error {
	sort := ctx.QueryParam("sort")

	if sort != "termahal" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "sort must be termahal",
		})
	}

	event, err := h.eventService.SortEventByExpensive(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) SortEventByCheapest(ctx echo.Context) error {
	sort := ctx.QueryParam("sort")

	if sort != "termurah" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "sort must be termurah",
		})
	}

	event, err := h.eventService.SortEventByCheapest(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) SortEventByNewest(ctx echo.Context) error {
	sort := ctx.QueryParam("sort")

	if sort != "terbaru" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "sort must be terbaru",
		})
	}

	event, err := h.eventService.SortEventByNewest(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, event)
}
