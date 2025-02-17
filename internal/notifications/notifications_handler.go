package notifications

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type NotificationHandler struct {
	service NotificationServiceInterface
}

func NewHandler(service NotificationServiceInterface) *NotificationHandler {
	return &NotificationHandler{service: service}
}

func (h *NotificationHandler) HandleGetAllNotifications(w http.ResponseWriter, r *http.Request) {
	notification, err := h.service.GetAllNotification()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, notification)
}

func (h *NotificationHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	notification, err := h.service.GetNotificationByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting notification by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Notification with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, notification)
}

func (h *NotificationHandler) HandleCreateNotification(w http.ResponseWriter, r *http.Request) {
	var payload CreateNotificationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	message, err := h.service.CreateNotification(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, message)
}

func (h *NotificationHandler) HandleUpdateNotification(w http.ResponseWriter, r *http.Request) {
	var payload UpdateNotificationPayload
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

	n, err := h.service.GetNotificationByID(payload.ID)
	if err != nil {
		fmt.Printf("error get notification by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("notification id %d not found"))
		return
	}

	if n == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Notification with ID %d does not exist", payload.ID))
		return
	}

	if payload.UserID == 0 && payload.Message == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Notification User ID And Message Cannot Be Empty!"))
		return
	}

	message, err := h.service.UpdateNotification(&payload)
	if err != nil {
		fmt.Printf("error update notification: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}

func (h *NotificationHandler) HandleDeleteNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get notification by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("notification id %d not found", id))
		return
	}

	message, err := h.service.DeleteNotification(id)
	if err != nil {
		fmt.Printf("error delete notification: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, message)
}
