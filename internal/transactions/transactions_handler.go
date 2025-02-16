package transactions

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type Handler struct {
	store TransactionStore
}

func NewHandler(store TransactionStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) TransactionRoutes(router *mux.Router) {
	router.HandleFunc("/transaction", h.HandleGet).Methods("GET")
	router.HandleFunc("/transaction/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/transaction", h.HandleCreate).Methods("POST")
	router.HandleFunc("/transaction/{id}", h.HandleUpdate).Methods("PUT")
	router.HandleFunc("/transaction/{id}", h.HandleDelete).Methods("DELETE")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.store.GetAllTransaction()
	if err != nil {
		fmt.Printf("error getting all transaction: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to retrieve transactions"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, transactions)
}

func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	transaction, err := h.store.GetTransactionByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting transaction by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Transaction with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, transaction)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload CreateTransactionPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := h.store.GetTransactionByReservationID(payload.ReservationID)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Transaction Number %s already exists", payload.ReservationID))
		return
	}

	err = h.store.CreateTransaction(&CreateTransactionPayload{
		ReservationID:   payload.ReservationID,
		Amount:          payload.Amount,
		PaymentMethodID: payload.PaymentMethodID,
		Status:          payload.Status,
	})
	if err != nil {
		fmt.Printf("error create transaction: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Create transaction %s successfully", payload.ReservationID))
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
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

	t, err := h.store.GetTransactionByID(payload.ID)
	if err != nil {
		fmt.Printf("error get transaction by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("transaction id %d not found"))
		return
	}

	if t == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Transaction with ID %d does not exist", payload.ID))
		return
	}

	if payload.ReservationID == 0 && payload.PaymentMethodID == 0 {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Transaction Reservation And Payment Method Cannot Be Empty!"))
		return
	}

	err = h.store.UpdateTransaction(&UpdateTransactionPayload{
		ID:              payload.ID,
		ReservationID:   payload.ReservationID,
		Amount:          payload.Amount,
		PaymentMethodID: payload.PaymentMethodID,
		Status:          payload.Status,
	})
	if err != nil {
		fmt.Printf("error update transaction: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update transaction %s successfully", t.PaymentMethodID))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get transaction by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("transaction id %d not found", id))
		return
	}

	err = h.store.DeleteTransaction(id)
	if err != nil {
		fmt.Printf("error delete transaction: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete transaction successfully"))
}
