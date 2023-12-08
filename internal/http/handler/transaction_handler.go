package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/trikrama/Depublic/common"
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
}

func NewTransactionHandler(cfg *config.Config, transactionService service.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) WebHookTransaction(c echo.Context) error {
	fmt.Println("ini webhook")
	var notficationPayload entity.TransactionPaymentMidtrans
	if err := c.Bind(&notficationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	fmt.Println("ini webhook : ", notficationPayload)
	tx, err := h.transactionService.GetTransactionByID(c.Request().Context(), notficationPayload.OrderID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	if notficationPayload.TransactionStatus == "settlement" {
		tx.TransactionStatus = "paid"
	}

	err = h.transactionService.UpdateTransaction(c.Request().Context(), tx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	// err = h.transactionService.UpdateStatus(c.Request().Context(), tx.ID, tx.TransactionStatus)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, echo.Map{
	// 		"message": err.Error(),
	// 	})
	// }
	notif := &notifEntity.NotificationRequest{
		UserID: tx.UserID,
		Title:  "Transaksi Berhasil",
		Body:   fmt.Sprintf("User Dengan Nama: %d, Telah Berhasil Membuat Transaksi", tx.UserID),
		Status: "Berhasil",
	}
	newNotif := notifEntity.NewNotification(*notif)
	err = h.notifService.CreateNotification(c.Request().Context(), newNotif)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update transaction status",
	})
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	transaction := entity.TransactionRequest{}
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	dataUser, _ := c.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*common.JwtCustomClaims)

	transaction.UserId = claims.ID
	transactionNewRequest := entity.NewTransaction(transaction)
	res, err := h.transactionService.CreateTransaction(c.Request().Context(), transactionNewRequest)
	if err != nil {
		fmt.Println("error 8")
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	event, err := h.transactionService.GetEvent(c.Request().Context(), res.EventID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	newHistory := entity.NewHistoryTransaction(res.ID, res.UserID, "create", res.CreatedAt, event.Name, int64(res.Quantity), int64(res.Total))
	err = h.transactionService.CreateHistory(c.Request().Context(), &newHistory)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	transactionResponse := entity.NewTransactionResponse(res)
	return c.JSON(http.StatusCreated, transactionResponse)
}

func (h *TransactionHandler) UpdateTransaction(c echo.Context) error {
	transaction := entity.TransactionRequestUpdate{}
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	transactionRequest := entity.NewTransactionUpdate(transaction)
	if err := h.transactionService.UpdateTransaction(c.Request().Context(), transactionRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	event, err := h.transactionService.GetEvent(c.Request().Context(), transactionRequest.EventID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	newHistory := entity.NewHistoryTransaction(transactionRequest.ID, transactionRequest.UserID, "update", transactionRequest.CreatedAt, event.Name, int64(transactionRequest.Quantity), int64(transactionRequest.Total))
	err = h.transactionService.CreateHistory(c.Request().Context(), &newHistory)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}
	transactionResponse := entity.NewTransactionResponse(transactionRequest)
	return c.JSON(http.StatusOK, transactionResponse)
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
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*common.JwtCustomClaims)
	transaction, err := h.transactionService.GetTransactionByID(c.Request().Context(), id)
	if err != nil {
		return err
	}
	if claims.Role != "Admin" {
		if transaction.UserID != claims.ID {
			return c.JSON(http.StatusForbidden, echo.Map{
				"message": "Anda tidak punya akses untuk akun ini",
			})
		}
	}

	return c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) GetTransactions(c echo.Context) error {
	transactions, err := h.transactionService.GetAllTransaction(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var transactionDetail []entity.TransactionResponse
	for _, transaction := range transactions {
		transactionResp := entity.NewTransactionResponse(transaction)
		transactionDetail = append(transactionDetail, *transactionResp)
	}
	return c.JSON(http.StatusOK, transactionDetail)
}

func (h *TransactionHandler) GetTransactionsByUser(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	transactions, err := h.transactionService.GetTransactionByUser(c.Request().Context(), int64(idInt))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var transactionDetail []entity.TransactionResponse
	for _, transaction := range transactions {
		transactionResp := entity.NewTransactionResponse(transaction)
		transactionDetail = append(transactionDetail, *transactionResp)
	}
	return c.JSON(http.StatusOK, transactionDetail)
}

func (h *TransactionHandler) GetAllHistory(c echo.Context) error {
	histories, err := h.transactionService.GetAllHistory(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, histories)
}

func (h *TransactionHandler) GetHistoryByUser(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	histories, err := h.transactionService.GetHistoryByUser(c.Request().Context(), int64(idInt))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, histories)
}

