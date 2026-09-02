package v1

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/wascool0nce/blog-platform-api/intenal/handler/v1/request"
	"github.com/wascool0nce/blog-platform-api/intenal/handler/v1/response"
	"github.com/wascool0nce/blog-platform-api/intenal/model"
)

type PostService interface {
	CreatePost(title, content, category string) (*model.Post, error)
}

type V1 struct {
	postService PostService
}

func NewPostHandler(postService PostService) *V1 {
	return &V1{
		postService: postService,
	}
}

func (h *V1) PostCreateHandler(w http.ResponseWriter, r *http.Request) {
	var body request.CreatePostRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(body); err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{
			"error": "failed validate post request",
		})
		return
	}

	post, err := h.postService.CreatePost(body.Title, body.Content, body.Category)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "failed to create post",
		})
		return
	}

	response := response.PostCreatedResponce{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Category:  post.Category,
		CreatedAt: post.CreatedAt,
	}

	writeJSON(w, http.StatusCreated, response)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Contnet-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}
