package vehicles

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type Handler struct {
	store VehicleStore
}

func NewHandler(store VehicleStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) VehicleRoutes(router *mux.Router) {
	router.HandleFunc("/vehicle", h.HandleGet).Methods("GET")
	router.HandleFunc("/vehicle/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/vehicle", h.HandleCreate).Methods("POST")
	router.HandleFunc("/vehicle/{id}", h.HandleUpdate).Methods("PUT")
	router.HandleFunc("/vehicle/{id}", h.HandleDelete).Methods("DELETE")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	vehicles, err := h.store.GetAllVehicle()
	if err != nil {
		fmt.Printf("error getting all vehicle: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to retrieve vehicles"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, vehicles)
}

func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	vehicle, err := h.store.GetVehicleByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting vehicle by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Vehicle with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, vehicle)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload CreateVehiclePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := h.store.GetVehicleByPlate(payload.PlateNumber)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Vehicle Plate %s already exists", payload.PlateNumber))
		return
	}

	err = h.store.CreateVehicle(&CreateVehiclePayload{
		UserID:      payload.UserID,
		PlateNumber: payload.PlateNumber,
		Type:        payload.Type,
		Brand:       payload.Brand,
		Model:       payload.Model,
		Color:       payload.Color,
	})
	if err != nil {
		fmt.Printf("error create vehicle: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Create vehicle %s successfully", payload.PlateNumber))
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var payload UpdateVehiclePayload
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

	v, err := h.store.GetVehicleByID(payload.ID)
	if err != nil {
		fmt.Printf("error get vehicle by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("vehicle id %d not found"))
		return
	}

	if v == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Vehicle with ID %d does not exist", payload.ID))
		return
	}

	if payload.UserID == 0 && payload.PlateNumber == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Vehicle User And Plat Number Cannot Be Empty!"))
		return
	}

	err = h.store.UpdateVehicle(&UpdateVehiclePayload{
		ID:          payload.ID,
		UserID:      payload.UserID,
		PlateNumber: payload.PlateNumber,
		Type:        payload.Type,
		Brand:       payload.Brand,
		Model:       payload.Model,
		Color:       payload.Color,
	})
	if err != nil {
		fmt.Printf("error update vehicle: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update vehicle %s successfully", v.PlateNumber))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get vehicle by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("vehicle id %d not found", id))
		return
	}

	err = h.store.DeleteVehicle(id)
	if err != nil {
		fmt.Printf("error delete vehicle: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete vehicle successfully"))
}
