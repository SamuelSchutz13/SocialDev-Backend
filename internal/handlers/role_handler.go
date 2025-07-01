package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/SamuelSchutz13/SocialDev/internal/db"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
	"github.com/google/uuid"
)

type RoleHandler struct {
	roleService *services.RoleService
}
type RoleWithUserRequest struct {
	RoleID uuid.UUID `json:"role_id"`
	UserID uuid.UUID `json:"user_id"`
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

func (h *RoleHandler) GetRoleHandler(w http.ResponseWriter, r *http.Request) {
	role_id := r.URL.Path
	roleID := strings.TrimPrefix(role_id, "/role/")

	uuid, err := uuid.Parse(roleID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid role ID format: %v", err), http.StatusBadRequest)
		return
	}

	role, err := h.roleService.GetRole(uuid)

	if err != nil {
		http.Error(w, "Failed to get role: "+err.Error(), http.StatusInternalServerError)
	}

	res, err := json.Marshal(role)

	if err != nil {
		http.Error(w, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (h *RoleHandler) GetAllRolesHandler(w http.ResponseWriter, r *http.Request) {
	roles, err := h.roleService.GetAllRoles()

	if err != nil {
		http.Error(w, "Failed to get roles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(roles)

	if err != nil {
		http.Error(w, "Failed to marshal response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (h *RoleHandler) GetRoleWithNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	role, err := h.roleService.GetRoleWithName(name)

	if err != nil {
		http.Error(w, "Failed to get role with name: "+err.Error(), http.StatusInternalServerError)
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

func (h *RoleHandler) UpdateRoleHandler(w http.ResponseWriter, r *http.Request) {
	role_id := r.URL.Path
	roleID := strings.TrimPrefix(role_id, "/role/")

	uuid, err := uuid.Parse(roleID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid role ID format: %v", err), http.StatusBadRequest)
		return
	}

	role, err := h.roleService.GetRole(uuid)

	if err != nil {
		http.Error(w, "Role not found: "+err.Error(), http.StatusNotFound)
		return
	}

	rb := r.Body
	reader, err := io.ReadAll(rb)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	rb.Close()

	var roleUpdate *db.UpdateRoleParams

	err = json.Unmarshal(reader, &roleUpdate)

	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	role, err = h.roleService.UpdateRole(role.RoleID, roleUpdate.Name)

	if err != nil {
		http.Error(w, "Failed to update role: "+err.Error(), http.StatusInternalServerError)
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

func (h *RoleHandler) DeleteRoleHandler(w http.ResponseWriter, r *http.Request) {
	role_id := r.URL.Path
	roleID := strings.TrimPrefix(role_id, "/role/")

	uuid, err := uuid.Parse(roleID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid role ID format: %v", err), http.StatusBadRequest)
		return
	}

	err = h.roleService.DeleteRole(uuid)

	if err != nil {
		http.Error(w, "Failed to delete role: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Role deleted successfully"))
}

func (h *RoleHandler) CreateUserWithRoleHandler(w http.ResponseWriter, r *http.Request) {
	rb := r.Body
	reader, err := io.ReadAll(rb)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	rb.Close()

	var userWithRole *RoleWithUserRequest

	err = json.Unmarshal(reader, &userWithRole)

	if err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
		return
	}

	user, err := h.roleService.CreateUserWithRole(userWithRole.UserID, userWithRole.RoleID)

	if err != nil {
		http.Error(w, "Failed to create user with role: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Failed to marshal user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func (h *RoleHandler) DeleteUserWithRoleHandler(w http.ResponseWriter, r *http.Request) {
	rb := r.Body
	reader, err := io.ReadAll(rb)

	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	rb.Close()

	var userWithRole *db.DeleteUserWithRoleParams

	err = json.Unmarshal(reader, &userWithRole)

	if err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
		return
	}

	err = h.roleService.DeleteUserWithRole(userWithRole.UserID, userWithRole.RoleID)

	if err != nil {
		http.Error(w, "Failed to delete user with role: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User with role deleted successfully"))
}
