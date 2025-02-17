package payment_methods

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type PaymentMethodHandler struct {
	service PaymentMethodServiceInterface
}

func NewHandler(service PaymentMethodServiceInterface) *PaymentMethodHandler {
	return &PaymentMethodHandler{service: service}
}

func (h *PaymentMethodHandler) HandleGetAllPaymentMethod(w http.ResponseWriter, r *http.Request) {
	payments, err := h.service.GetAllPaymentMethod()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, payments)
}

func (h *PaymentMethodHandler) HandleGetPaymentMethodByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	payment, err := h.service.GetPaymentMethodByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, payment)
}

func (h *PaymentMethodHandler) HandleCreatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	var payload CreatePaymentMethodPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	message, err := h.service.CreatePaymentMethod(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, message)
}

func (h *PaymentMethodHandler) HandleUpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {
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

	pm, err := h.service.GetPaymentMethodByID(payload.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
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

	message, err := h.service.UpdatePaymentMethod(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}

func (h *PaymentMethodHandler) HandleDeletePaymentMethod(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	message, err := h.service.DeletePaymentMethod(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}
