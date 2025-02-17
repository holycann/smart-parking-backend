package transactions

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type TransactionHandler struct {
	service TransactionServiceInterface
}

func NewHandler(service TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleGetAllTransaction(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.service.GetAllTransaction()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, transactions)
}

func (h *TransactionHandler) HandleGetTransactionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	transaction, err := h.service.GetTransactionByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, transaction)
}

func (h *TransactionHandler) HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	var payload CreateTransactionPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := h.service.GetTransactionByID(payload.ReservationID)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	message, err := h.service.CreateTransaction(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, message)
}

func (h *TransactionHandler) HandleUpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var payload UpdateTransactionPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v", err))
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	payload.ID = id

	if err := utils.Validate.Struct(payload); err != nil {
		fmt.Printf("error validating payload: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	message, err := h.service.UpdateTransaction(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}

func (h *TransactionHandler) HandleDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get transaction by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("transaction id %d not found", id))
		return
	}

	message, err := h.service.DeleteTransaction(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}
