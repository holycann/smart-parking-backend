package auth

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/holycann/smart-parking-backend/internal/middleware"
	utils "github.com/holycann/smart-parking-backend/pkg"
)

type AuthHandler struct {
	service AuthServiceInterface
}

func NewHandler(service AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	var payload LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		fmt.Printf("error parsing json: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		fmt.Printf("error validating payload: %v\n", err)
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", errors))
		return
	}

	token, err := h.service.UserLogin(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	middleware.SetJWTHttpOnlyCookie(w, r, token, false)

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
