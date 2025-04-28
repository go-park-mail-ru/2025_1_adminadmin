package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/survey/delivery/grpc/gen"
	jwtUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/jwt"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	utils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/send_error"
	"github.com/golang-jwt/jwt"
	"github.com/mailru/easyjson"
	"github.com/samber/lo"
	"github.com/satori/uuid"
)

type SurveyHandler struct {
	client gen.StatClient
	secret string
}

func CreateSurveyHandler(client gen.StatClient) *SurveyHandler {
	return &SurveyHandler{
		client: client,
		secret: os.Getenv("JWT_SECRET"),
	}
}

func (h *SurveyHandler) Vote(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	req := models.Vote{}
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	_, err := h.client.Vote(r.Context(), &gen.VoteRequest{
		QuestionId: req.QuestionId.String(),
		Vote:       int32(req.Voice),
	})
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusBadRequest)
		utils.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func getQuestion(question *gen.Question) models.Question {
	return models.Question{
		Id:           uuid.FromStringOrNil(question.Id),
		Title:        question.Title,
		MinMark:      int(question.MinMark),
		Skip:         int(question.Skip),
		QuestionType: question.QuestionType,
		Number:       int(question.Number),
		SurveyId:     uuid.FromStringOrNil(question.SurveyId),
	}
}

func (h *SurveyHandler) CreateSurvey(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	cookie, err := r.Cookie("AdminJWT")
	if err != nil {
		if err == http.ErrNoCookie {
			log.LogHandlerError(logger, fmt.Errorf("токен отсутствует: %w", err), http.StatusUnauthorized)
			utils.SendError(w, "токен отсутствует", http.StatusUnauthorized)
			return
		}
		log.LogHandlerError(logger, fmt.Errorf("ошибка при чтении куки: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка при чтении куки", http.StatusBadRequest)
		return
	}
	JWTStr := cookie.Value

	claims := jwt.MapClaims{}

	login, ok := jwtUtils.GetLoginFromJWT(JWTStr, claims, h.secret)
	if !ok || login == "" {
		log.LogHandlerError(logger, errors.New("недействительный токен: login отсутствует"), http.StatusUnauthorized)
		utils.SendError(w, "недействительный токен: login отсутствует", http.StatusUnauthorized)
		return
	}

	req := models.CreateSurveyRequest{}
	if err := easyjson.UnmarshalFromReader(r.Body, &req); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка парсинга JSON: %w", err), http.StatusBadRequest)
		utils.SendError(w, "ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	questions := make([]*gen.CreateQuestionRequest, len(req.Questions))
	for i, question := range req.Questions {
		questions[i] = &gen.CreateQuestionRequest{
			Title:        question.Title,
			MinMark:      int64(question.MinMark),
			Skip:         int64(question.Skip),
			QuestionType: question.QuestionType,
		}
	}

	if _, err := h.client.CreateSurvey(r.Context(), &gen.CreateSurveyRequest{
		Questions: questions,
	}); err != nil {
		log.LogHandlerError(logger, err, http.StatusBadRequest)
		utils.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.LogHandlerInfo(logger, "Success", http.StatusOK)
	w.WriteHeader(http.StatusNoContent)
}

func getStat(stat *gen.StatModel) models.Stat {
	val, _ := strconv.ParseInt(stat.Label, 10, 32)
	return models.Stat{
		QuestionId:   uuid.FromStringOrNil(stat.QuestionId),
		Title:        stat.QuestionTitle,
		QuestionType: stat.QuestionType,
		Voice:        int(val),
		Count:        int(stat.Value),
	}
}
func (h *SurveyHandler) GetSurvey(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	protoData, err := h.client.GetSurvey(r.Context(), &gen.GetSurveyRequest{})
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusBadRequest)
		utils.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := make([]models.Question, len(protoData.Questions))

	for i, question := range protoData.Questions {
		data[i] = getQuestion(question)
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка формирования JSON: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "ошибка формирования JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}

func getCSAT(data []models.Stat) (map[int]int, float64) {
	result := make(map[int]int)
	var summary int
	var total int

	data = lo.Filter(data, func(item models.Stat, i int) bool {
		return item.Voice > 0 && item.Voice <= 5
	})

	for _, item := range data {
		result[item.Voice] = item.Count
		summary += item.Voice * item.Count
		total += item.Count
	}
	return result, float64(summary) / float64(total)
}
func getNPS(data []models.Stat) (map[string]int, float64) {
	result := make(map[string]int)
	// for _, item := range data {
	// 	if item.Voice < 6{
	// 		result["d"] = item.Count
	// 	}
	// }

	lo.ForEach(data, func(d models.Stat, i int) {
		if d.Voice <= 6 {
			data[i].Type = "detractor"
		} else if d.Voice >= 6 && d.Voice <= 8 {
			data[i].Type = "passive"
		} else {
			data[i].Type = "promouter"
		}
	})
	var total int
	for _, v := range data {
		result[v.Type] += v.Count
		total += v.Count
	}

	d := result["detractor"]
	p := result["promouter"]

	var r float64
	if total != 0 {
		r = float64(p-d) / float64(total)
	}
	return result, r
}
func (h *SurveyHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	protoData, err := h.client.GetStats(r.Context(), &gen.GetStatsRequest{})
	if err != nil {
		log.LogHandlerError(logger, err, http.StatusBadRequest)
		utils.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := make([]models.Stat, len(protoData.Stats))

	for i, s := range protoData.Stats {
		data[i] = getStat(s)
	}

	var groupedData map[string][]models.Stat

	groupedData = lo.GroupBy(data, func(item models.Stat) string {
		return string(item.QuestionId.String())

	})

	var resp []models.StatResponse

	for _, item := range groupedData {
		if len(item) == 0 {
			continue
		}

		respItem := models.StatResponse{
			QuestionId:   item[0].QuestionId,
			Title:        item[0].Title,
			QuestionType: item[0].QuestionType,
		}
		switch item[0].QuestionType {
		case "CSAT":
			respItem.Stats, respItem.Value = getCSAT(item)
		case "NPS":
			respItem.Stats, respItem.Value = getNPS(item)

		}
		resp = append(resp, respItem)
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.LogHandlerError(logger, fmt.Errorf("ошибка формирования JSON: %w", err), http.StatusInternalServerError)
		utils.SendError(w, "ошибка формирования JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	log.LogHandlerInfo(logger, "Success", http.StatusOK)
}
