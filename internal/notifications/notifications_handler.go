package notifications

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	utils "github.com/holycann/smart-parking-backend/pkg"
)

type Handler struct {
	store NotificationStore
}

func NewHandler(store NotificationStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) NotificationRoutes(router *mux.Router) {
	router.HandleFunc("/notification", h.HandleGet).Methods("GET")
	router.HandleFunc("/notification/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/notification", h.HandleCreate).Methods("POST")
	router.HandleFunc("/notification/{id}", h.HandleUpdate).Methods("PUT")
	router.HandleFunc("/notification/{id}", h.HandleDelete).Methods("DELETE")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	notification, err := h.store.GetAllNotification()
	if err != nil {
		fmt.Printf("error getting all notification: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to retrieve notifications"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, notification)
}

func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	notification, err := h.store.GetNotificationByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting notification by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Notification with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, notification)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload CreateNotificationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := h.store.GetNotificationByMessage(payload.Message)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Notification Message %s already exists", payload.Message))
		return
	}

	err = h.store.CreateNotification(&CreateNotificationPayload{
		UserID:  payload.UserID,
		Message: payload.Message,
		Status:  payload.Status,
	})
	if err != nil {
		fmt.Printf("error create notification: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Create notification %s successfully", payload.Message))
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
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

	n, err := h.store.GetNotificationByID(payload.ID)
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

	err = h.store.UpdateNotification(&UpdateNotificationPayload{
		ID:      payload.ID,
		UserID:  payload.UserID,
		Message: payload.Message,
		Status:  payload.Status,
	})
	if err != nil {
		fmt.Printf("error update notification: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update notification %s successfully", n.Message))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get notification by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("notification id %d not found", id))
		return
	}

	err = h.store.DeleteNotification(id)
	if err != nil {
		fmt.Printf("error delete notification: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete notification successfully"))
}
