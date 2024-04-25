package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	internal "go-gpt-api/internal/openAI"
)

type MergeRequest struct {
	New      string `json:"new"`
	Old      string `json:"old"`
	Rejected string `json:"rejected,omitempty"`
}

type MergeSession struct {
	History []MergeRequest
	mutex   sync.Mutex
}

var sessions = make(map[string]*MergeSession)

func NewMergeSession() *MergeSession {
	return &MergeSession{
		History: make([]MergeRequest, 0),
	}
}

func (s *MergeSession) AddToHistory(msg MergeRequest) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.History = append(s.History, msg)
}

type Handler struct {
	OpenAI *internal.OpenAI
}

func NewHandler(openAI *internal.OpenAI) *Handler {
	return &Handler{
		OpenAI: openAI,
	}
}

func (h *Handler) MergeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read the body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Printf("Received request: %s\n", string(body))

	var request MergeRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Failed to unmarshal the body", http.StatusInternalServerError)
		return
	}

	response, err := h.OpenAI.Merge(internal.MergeRequest(request))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
