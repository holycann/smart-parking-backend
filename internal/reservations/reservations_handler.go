package reservations

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type ReservationHandler struct {
	service ReservationServiceInterface
}

func NewHandler(service ReservationServiceInterface) *ReservationHandler {
	return &ReservationHandler{service: service}
}

func (h *ReservationHandler) HandleGetAllReservation(w http.ResponseWriter, r *http.Request) {
	reservations, err := h.service.GetAllReservation()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, reservations)
}

func (h *ReservationHandler) HandleGetReservationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	reservation, err := h.service.GetReservationByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, reservation)
}

func (h *ReservationHandler) HandleCreateReservation(w http.ResponseWriter, r *http.Request) {
	var payload CreateReservationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	message, err := h.service.CreateReservation(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, message)
}

func (h *ReservationHandler) HandleUpdateReservation(w http.ResponseWriter, r *http.Request) {
	var payload UpdateReservationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v", err))
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	payload.ID = id

	if err := utils.Validate.Struct(payload); err != nil {
		fmt.Printf("error validating payload: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	re, err := h.service.GetReservationByID(payload.ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if re == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Reservation with ID %d does not exist", payload.ID))
		return
	}

	if payload.StartTime <= time.Now().Unix() && payload.EndTime >= payload.StartTime {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Reservation Start Time Must Be <= Time Now And End Time Must Be >= Start Time"))
		return
	}

	message, err := h.service.UpdateReservation(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}

func (h *ReservationHandler) HandleDeleteReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get reservation by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("reservation id %d not found", id))
		return
	}

	message, err := h.service.DeleteReservation(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}
