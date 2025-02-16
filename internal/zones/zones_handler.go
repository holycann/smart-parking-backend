package zones

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type Handler struct {
	store ZoneStore
}

func NewHandler(store ZoneStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) ZoneRoutes(router *mux.Router) {
	router.HandleFunc("/zone", h.HandleGet).Methods("GET")
	router.HandleFunc("/zone/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/zone", h.HandleCreate).Methods("POST")
	router.HandleFunc("/zone/{id}", h.HandleUpdate).Methods("PUT")
	router.HandleFunc("/zone/{id}", h.HandleDelete).Methods("DELETE")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	zones, err := h.store.GetAllZone()
	if err != nil {
		fmt.Printf("error getting all zone: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to retrieve zones"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, zones)
}

func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	zone, err := h.store.GetZoneByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting zone by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("ZOne with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, zone)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload CreateZonePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := h.store.GetZoneByName(payload.Name)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Zone Name %s already exists", payload.Name))
		return
	}

	err = h.store.CreateZone(&CreateZonePayload{
		Name:       payload.Name,
		Location:   payload.Location,
		TotalSpots: payload.TotalSpots,
	})
	if err != nil {
		fmt.Printf("error create user: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Create zone %s successfully", payload.Name))
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var payload UpdateZonePayload
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

	z, err := h.store.GetZoneByID(payload.ID)
	if err != nil {
		fmt.Printf("error get zone by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("zone id %d not found"))
		return
	}

	if z == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Zone with ID %d does not exist", payload.ID))
		return
	}

	if payload.Name == "" && payload.Location == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Zone Name And Location Cannot Be Empty!"))
		return
	}

	err = h.store.UpdateZone(&UpdateZonePayload{
		ID:         payload.ID,
		Name:       payload.Name,
		Location:   payload.Location,
		TotalSpots: payload.TotalSpots,
	})
	if err != nil {
		fmt.Printf("error update zone: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update zone %s successfully", z.Name))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get user by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("zone id %d not found"))
		return
	}

	err = h.store.DeleteZone(id)
	if err != nil {
		fmt.Printf("error delete zone: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete zone successfully"))
}
