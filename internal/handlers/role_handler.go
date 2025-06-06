package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
	"github.com/google/uuid"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

func (h *RoleHandler) CreateRoleHandler(w http.ResponseWriter, r *http.Request) {
	rb := r.Body
	reader, err := io.ReadAll(rb)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	rb.Close()

	var roleCreate *db.CreateRoleParams

	err = json.Unmarshal(reader, &roleCreate)

	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	role_id := uuid.New()

	role, err := h.roleService.CreateRole(role_id, roleCreate.Name)

	if err != nil {
		http.Error(w, "Failed to create role: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(role)

	if err != nil {
		http.Error(w, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
