package spots

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type Handler struct {
	store SpotStore
}

func NewHandler(store SpotStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) SpotRoutes(router *mux.Router) {
	router.HandleFunc("/spot", h.HandleGet).Methods("GET")
	router.HandleFunc("/spot/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/spot", h.HandleCreate).Methods("POST")
	router.HandleFunc("/spot/{id}", h.HandleUpdate).Methods("PUT")
	router.HandleFunc("/spot/{id}", h.HandleDelete).Methods("DELETE")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	spots, err := h.store.GetAllSpot()
	if err != nil {
		fmt.Printf("error getting all spot: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to retrieve spots"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, spots)
}

func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	spot, err := h.store.GetSpotByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting spot by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Spot with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, spot)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload CreateSpotPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := h.store.GetSpotByNumber(payload.SpotNumber)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Spot Number %s already exists", payload.SpotNumber))
		return
	}

	err = h.store.CreateSpot(&CreateSpotPayload{
		ZoneID:     payload.ZoneID,
		SpotNumber: payload.SpotNumber,
		Status:     payload.Status,
	})
	if err != nil {
		fmt.Printf("error create spot: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Create spot %s successfully", payload.SpotNumber))
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var payload UpdateSpotPayload
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

	v, err := h.store.GetSpotByID(payload.ID)
	if err != nil {
		fmt.Printf("error get spot by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("spot id %d not found"))
		return
	}

	if v == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Spot with ID %d does not exist", payload.ID))
		return
	}

	if payload.ZoneID == 0 && payload.SpotNumber == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Spot Zone And Number Cannot Be Empty!"))
		return
	}

	err = h.store.UpdateSpot(&UpdateSpotPayload{
		ID:         payload.ID,
		ZoneID:     payload.ZoneID,
		SpotNumber: payload.SpotNumber,
		Status:     payload.Status,
	})
	if err != nil {
		fmt.Printf("error update spot: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update spot %s successfully", v.SpotNumber))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get spot by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("spot id %d not found", id))
		return
	}

	err = h.store.DeleteSpot(id)
	if err != nil {
		fmt.Printf("error delete spot: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete spot successfully"))
}
