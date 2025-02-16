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

type Handler struct {
	store ReservationStore
}

func NewHandler(store ReservationStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) ReservationRoutes(router *mux.Router) {
	router.HandleFunc("/reservation", h.HandleGet).Methods("GET")
	router.HandleFunc("/reservation/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/reservation", h.HandleCreate).Methods("POST")
	router.HandleFunc("/reservation/{id}", h.HandleUpdate).Methods("PUT")
	router.HandleFunc("/reservation/{id}", h.HandleDelete).Methods("DELETE")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	reservations, err := h.store.GetAllReservation()
	if err != nil {
		fmt.Printf("error getting all reservation: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to retrieve reservations"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, reservations)
}

func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	reservation, err := h.store.GetReservationByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting reservation by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Reservation with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, reservation)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload CreateReservationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := h.store.GetReservationByStartTime(payload.StartTime)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Reservation Start Time %s already exists", payload.StartTime))
		return
	}

	err = h.store.CreateReservation(&CreateReservationPayload{
		UserID:    payload.UserID,
		SpotID:    payload.SpotID,
		VehicleID: payload.VehicleID,
		StartTime: payload.StartTime,
		EndTime:   payload.EndTime,
		Status:    payload.Status,
	})
	if err != nil {
		fmt.Printf("error create reservation: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Create reservation %s successfully", payload.StartTime))
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var payload UpdateReservationPayload
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

	re, err := h.store.GetReservationByID(payload.ID)
	if err != nil {
		fmt.Printf("error get reservation by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("reservation id %d not found"))
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

	err = h.store.UpdateReservation(&UpdateReservationPayload{
		ID:        payload.ID,
		UserID:    payload.UserID,
		SpotID:    payload.SpotID,
		VehicleID: payload.VehicleID,
		StartTime: payload.StartTime,
		EndTime:   payload.EndTime,
		Status:    payload.Status,
	})
	if err != nil {
		fmt.Printf("error update reservation: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update reservation %v successfully", re.StartTime))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get reservation by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("reservation id %d not found", id))
		return
	}

	err = h.store.DeleteReservation(id)
	if err != nil {
		fmt.Printf("error delete reservation: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete reservation successfully"))
}
