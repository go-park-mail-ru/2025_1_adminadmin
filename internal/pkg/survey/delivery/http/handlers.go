package http

import (
	"encoding/json"
	"net/http"

	generatedSurvey "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/survey/delivery/grpc/gen/proto"
)

type SurveyHandler struct {
	client generatedSurvey.StatClient
}

func NewSurveyHandler(client generatedSurvey.StatClient) *SurveyHandler {
	return &SurveyHandler{client: client}
}
func (h *SurveyHandler) GetSurvey(w http.ResponseWriter, r *http.Request) {
	sur, err := h.client.GetSurvey(r.Context(),&generatedSurvey.GetSurveyRequest{})
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sur)

}
