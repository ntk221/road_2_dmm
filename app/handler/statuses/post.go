package status

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"yatter-backend-go/app/domain/object"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// PostRequest
// Request body for `POST /v1/statuses`
type PostRequest struct {
	Status   string
	MediaIDs []int
}

// PostResponse
// Response body for `POST /v1/statuses`
type PostResponse struct {
	ID              object.StatusID
	Account         object.Account
	Content         string
	CreateAt        object.DateTime
	MediaAttachment []string // 仮実装
}

// Post
// Handle request for `POST /v1/statuses`
func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req PostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		httperror.BadRequest(w, err)
		return
	}

	// log.Printf("req: %+v", req)

	account := auth.AccountOf(r)
	accountID := account.ID
	content := req.Status
	log.Printf("content: %v", content)
	status, err := object.CreateStatus(accountID, content)
	if err != nil {
		log.Println(err)
		httperror.InternalServerError(w, err)
	}
	statusRepo := h.app.Dao.Status()

	posted, err := statusRepo.PostStatus(ctx, status)
	if err != nil {
		log.Println(err)
		httperror.InternalServerError(w, err)
	}

	response := PostResponse{
		ID:              posted.ID,
		Account:         *account,
		Content:         posted.Content,
		CreateAt:        posted.CreateAt,
		MediaAttachment: []string{}, // 仮実装
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		httperror.InternalServerError(w, errors.New("サーバー側でなんらかの問題が発生しました"))
		return
	}
}
