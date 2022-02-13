package api

import (
	"encoding/json"
	store "go-case/internal/infra/store"
	"net/http"
)

type Handler struct {
	Store *store.Store
}

// Get godoc
// @Summary Get value with key
// @Tags keyvalue
// @Success 200
// @Param key query string true "Key name"
// @Router /api/get [get]
func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	queryKeyParam := queryParams["key"]
	k := queryKeyParam[0]
	v := h.Store.Get(k)

	Json(w, http.StatusOK, map[string]string{"value": v})
}

type setParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Set godoc
// @Summary Set key and value
// @Tags keyvalueset
// @Accept json
// @Produce json
// @Param request body setParam true "Key Value Data"
// @Router /api/set [post]
func (h *Handler) SetHandler(w http.ResponseWriter, r *http.Request) {
	var req setParam
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		JsonErr(w, http.StatusBadRequest, err)
		return
	}

	h.Store.Set(req.Key, req.Value)

	w.WriteHeader(http.StatusNoContent)
}

// Flush godoc
// @Summary Flush all key value pairs
// @Tags keyvalueflush
// @Router /api/flush [get]
func (h *Handler) FlushHandler(w http.ResponseWriter, r *http.Request) {
	h.Store.Flush()

	w.WriteHeader(http.StatusNoContent)
}
