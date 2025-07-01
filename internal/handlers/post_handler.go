package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	//configs "github.com/SamuelSchutz13/SocialDev/config"
	"github.com/SamuelSchutz13/SocialDev/internal/entity"
	"github.com/SamuelSchutz13/SocialDev/internal/services"
	"github.com/go-playground/validator/v10"
)

//var cfg = configs.LoadConfig()

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (ph *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	reader, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error to read request body", http.StatusInternalServerError)
		return
	}

	var postCreate *entity.PostResponse

	err = json.Unmarshal(reader, &postCreate)

	if err != nil {
		http.Error(w, "Error to unmarshal request body", http.StatusBadRequest)
		return
	}

	validate = validator.New()

	err = validate.Struct(postCreate)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		var validationErrors []string

		for _, e := range errors {
			validationErrors = append(validationErrors, e.Error())
		}

		response, _ := json.Marshal(map[string]interface{}{"errors": validationErrors})
		http.Error(w, string(response), http.StatusBadRequest)
		return
	}

	createPost, err := ph.postService.CreatePost(postCreate.UserID, postCreate.Title, postCreate.Content)

	if err != nil {
		http.Error(w, "Error to create post", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Post created successfully",
		"post":    createPost,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (ph *PostHandler) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	reader, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error to read request body", http.StatusInternalServerError)
		return
	}

	var postCreate *entity.PostResponse

	err = json.Unmarshal(reader, &postCreate)

	if err != nil {
		http.Error(w, "Error to unmarshal request body", http.StatusBadRequest)
		return
	}

	validate = validator.New()

	err = validate.Struct(postCreate)

	if err != nil {
		errors := err.(validator.ValidationErrors)
		var validationErrors []string

		for _, e := range errors {
			validationErrors = append(validationErrors, e.Error())
		}

		response, _ := json.Marshal(map[string]interface{}{"errors": validationErrors})
		http.Error(w, string(response), http.StatusBadRequest)
		return
	}

	createPost, err := ph.postService.CreatePost(postCreate.UserID, postCreate.Title, postCreate.Content)

	if err != nil {
		http.Error(w, "Error to create post", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Post created successfully",
		"post":    createPost,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
