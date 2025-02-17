package zones

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type ZoneHandler struct {
	service ZoneServiceInterface
}

func NewHandler(service ZoneServiceInterface) *ZoneHandler {
	return &ZoneHandler{service: service}
}

func (h *ZoneHandler) HandleGetAllZone(w http.ResponseWriter, r *http.Request) {
	zones, err := h.service.GetAllZone()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, zones)
}

func (h *ZoneHandler) HandleGetZoneByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	zone, err := h.service.GetZoneByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, zone)
}

func (h *ZoneHandler) HandleCreateZone(w http.ResponseWriter, r *http.Request) {
	var payload CreateZonePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	message, err := h.service.CreateZone(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, message)
}

func (h *ZoneHandler) HandleUpdateZone(w http.ResponseWriter, r *http.Request) {
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

	message, err := h.service.UpdateZone(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}

func (h *ZoneHandler) HandleDeleteZone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get user by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("zone id %d not found"))
		return
	}

	message, err := h.service.DeleteZone(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}
