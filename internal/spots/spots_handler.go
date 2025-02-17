package spots

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type SpotHandler struct {
	service SpotServiceInterface
}

func NewHandler(service SpotServiceInterface) *SpotHandler {
	return &SpotHandler{service: service}
}

func (h *SpotHandler) HandleGetAllSpot(w http.ResponseWriter, r *http.Request) {
	spots, err := h.service.GetAllSpot()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, spots)
}

func (h *SpotHandler) HandleGetSpotByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	spot, err := h.service.GetSpotByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, spot)
}

func (h *SpotHandler) HandleCreateSpot(w http.ResponseWriter, r *http.Request) {
	var payload CreateSpotPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	message, err := h.service.CreateSpot(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, message)
}

func (h *SpotHandler) HandleUpdateSpot(w http.ResponseWriter, r *http.Request) {
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

	message, err := h.service.UpdateSpot(&payload)
	if err != nil {
		fmt.Printf("error update spot: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}

func (h *SpotHandler) HandleDeleteSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get spot by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("spot id %d not found", id))
		return
	}

	message, err := h.service.DeleteSpot(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}
