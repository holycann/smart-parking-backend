package payment_methods

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type Handler struct {
	store PaymentMethodStore
}

func NewHandler(store PaymentMethodStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) PaymentMethodRoutes(router *mux.Router) {
	router.HandleFunc("/payment_method", h.HandleGet).Methods("GET")
	router.HandleFunc("/payment_method/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/payment_method", h.HandleCreate).Methods("POST")
	router.HandleFunc("/payment_method/{id}", h.HandleUpdate).Methods("PUT")
	router.HandleFunc("/payment_method/{id}", h.HandleDelete).Methods("DELETE")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	payments, err := h.store.GetAllPaymentMethod()
	if err != nil {
		fmt.Printf("error getting all payment_method: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to retrieve payment_methods"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, payments)
}

func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	payment, err := h.store.GetPaymentMethodByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting payment_method by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Payment_method with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, payment)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload CreatePaymentMethodPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := h.store.GetPaymentMethodByMethodName(payload.MethodName)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Payment Method With Method Name %s already exists", payload.MethodName))
		return
	}

	err = h.store.CreatePaymentMethod(&CreatePaymentMethodPayload{
		MethodName: payload.MethodName,
		Details:    payload.Details,
		Status:     payload.Status,
	})
	if err != nil {
		fmt.Printf("error create payment_method: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Create payment_method %s successfully", payload.MethodName))
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var payload UpdatePaymentMethodPayload
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

	pm, err := h.store.GetPaymentMethodByID(payload.ID)
	if err != nil {
		fmt.Printf("error get payment_method by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("payment_method id %d not found"))
		return
	}

	if pm == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Payment_method with ID %d does not exist", payload.ID))
		return
	}

	if payload.MethodName == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Payment Method Name Cannot Be Empty!!!"))
		return
	}

	err = h.store.UpdatePaymentMethod(&UpdatePaymentMethodPayload{
		ID:         payload.ID,
		MethodName: payload.MethodName,
		Details:    payload.Details,
		Status:     payload.Status,
	})
	if err != nil {
		fmt.Printf("error update payment: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update payment method %v successfully", pm.MethodName))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get payment_method by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("payment_method id %d not found", id))
		return
	}

	err = h.store.DeletePaymentMethod(id)
	if err != nil {
		fmt.Printf("error delete payment_method: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete payment_method successfully"))
}
