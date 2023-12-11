package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/common"
	eventService "github.com/trikrama/Depublic/internal/app/event/service"
	notifEntity "github.com/trikrama/Depublic/internal/app/notification/entity"
	notifService "github.com/trikrama/Depublic/internal/app/notification/service"
	"github.com/trikrama/Depublic/internal/app/transaction/entity"
	"github.com/trikrama/Depublic/internal/app/transaction/service"
	"github.com/trikrama/Depublic/internal/config"
	"github.com/trikrama/Depublic/internal/http/validator"
)

type TransactionHandler struct {
	transactionService service.TransactionServiceInterface
	notifService       notifService.NotificationServiceInterface
	eventService       eventService.EventServiceInterface
	paymentService     service.PaymentServiceInterface
}

func NewTransactionHandler(cfg *config.Config,
	transactionService service.TransactionServiceInterface,
	notifService notifService.NotificationServiceInterface,
	eventService eventService.EventServiceInterface,
	paymentService service.PaymentServiceInterface) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		notifService:       notifService,
		eventService:       eventService,
		paymentService:     paymentService,
	}
}

func (h *TransactionHandler) WebHookTransaction(c echo.Context) error {
	var notficationPayload entity.TransactionPaymentMidtrans
	if err := c.Bind(&notficationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	tx, err := h.transactionService.GetTransactionByID(c.Request().Context(), notficationPayload.OrderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	event, err := h.eventService.GetEventByID(c.Request().Context(), tx.EventID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error(),})
	}

	tx.TransactionStatus = notficationPayload.TransactionStatus 
	if notficationPayload.TransactionStatus == "settlement" {
		tx.TransactionStatus = "success"
		event.Quantity -= int64(tx.Quantity)
		if event.Quantity <= 0 {
			event.Status = "sold out"
		}
		err := h.eventService.UpdateEvent(c.Request().Context(),  event)
		if err != nil { 
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
		}
	}

	err = h.transactionService.UpdateTransaction(c.Request().Context(), tx, event)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error(),})
	}

	newHistory := entity.NewHistoryTransaction(tx.ID, tx.UserID, "Success Succes Payment", tx.CreatedAt, event.Name, int64(tx.Quantity), int64(tx.Total))
	err = h.transactionService.CreateHistory(c.Request().Context(), &newHistory)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{"message": err.Error(),})
	}

	notif := &notifEntity.NotificationRequest{
		UserID: tx.UserID,
		Title:  "Transaction",
		Body:   fmt.Sprintf("User Name: %d, Has Successfully Made Payment", tx.UserID),
		Status: "Success",
	}
	newNotif := notifEntity.NewNotification(*notif)
	err = h.notifService.CreateNotification(c.Request().Context(), newNotif)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "success payment"})
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	transaction := entity.TransactionRequest{}
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	//Mendapatkan Data Event
	event, err := h.eventService.GetEventByID(c.Request().Context(), transaction.EventId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}

	//Mendapatkan Data User Dari Token
	dataUser, _ := c.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*common.JwtCustomClaims)

	transaction.UserId = claims.ID

	//Membuat Transaksi
	transactionNewRequest := entity.NewTransaction(transaction)
	res, err := h.transactionService.CreateTransaction(c.Request().Context(), transactionNewRequest, event)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}

	//Melakukan Pembayaran
	paymentRequest := entity.NewPaymentRequest(res.ID, res.EventID, event.Price, int64(res.Total), res.Quantity, event.Name, event.Status, claims.Name, claims.Email)
	url, err := h.paymentService.CreateTransactionMidtrans(c.Request().Context(), paymentRequest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	res.RedirectURL = url
	transactionResponse := entity.NewTransactionResponse(res)
	return c.JSON(http.StatusCreated, transactionResponse)
}

func (h *TransactionHandler) UpdateTransaction(c echo.Context) error {
	transaction := entity.TransactionRequestUpdate{}
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	//Mendapatkan Data Event
	event, err := h.eventService.GetEventByID(c.Request().Context(), transaction.EventId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}

	//Mengupdate Transaksi
	transactionRequest := entity.NewTransactionUpdate(transaction)
	if err := h.transactionService.UpdateTransaction(c.Request().Context(), transactionRequest, event); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}

	//Mencatat Transaksi
	newHistory := entity.NewHistoryTransaction(transactionRequest.ID, transactionRequest.UserID, "Update Transaction", transactionRequest.CreatedAt, event.Name, int64(transactionRequest.Quantity), int64(transactionRequest.Total))
	err = h.transactionService.CreateHistory(c.Request().Context(), &newHistory)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	transactionResponse := entity.NewTransactionResponse(transactionRequest)
	return c.JSON(http.StatusOK, echo.Map{"data": transactionResponse})
}

func (h *TransactionHandler) DeleteTransaction(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	if err := h.transactionService.DeleteTransaction(c.Request().Context(), idInt); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "transaction deleted",
	})
}

func (h *TransactionHandler) GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	transaction, err := h.transactionService.GetTransactionByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": transaction})
}

func (h *TransactionHandler) GetTransactions(c echo.Context) error {
	transactions, err := h.transactionService.GetAllTransaction(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	var transactionDetail []entity.TransactionResponse
	for _, transaction := range transactions {
		transactionResp := entity.NewTransactionResponse(transaction)
		transactionDetail = append(transactionDetail, *transactionResp)
	}
	return c.JSON(http.StatusOK, echo.Map{"data": transactionDetail})
}

func (h *TransactionHandler) GetTransactionsByUser(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*common.JwtCustomClaims)
	if claims.Role != "Admin" {
		if int64(idInt) != claims.ID {
			return c.JSON(http.StatusForbidden, echo.Map{
				"message": "You do not have access to this account",
			})
		}
	}
	transactions, err := h.transactionService.GetTransactionByUser(c.Request().Context(), int64(idInt))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var transactionDetail []entity.TransactionResponse
	for _, transaction := range transactions {
		transactionResp := entity.NewTransactionResponse(transaction)
		transactionDetail = append(transactionDetail, *transactionResp)
	}
	return c.JSON(http.StatusOK, echo.Map{"data": transactionDetail})
}

func (h *TransactionHandler) GetAllHistory(c echo.Context) error {
	histories, err := h.transactionService.GetAllHistory(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": histories})
}

func (h *TransactionHandler) GetHistoryByUser(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	histories, err := h.transactionService.GetHistoryByUser(c.Request().Context(), int64(idInt))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": histories})
}
