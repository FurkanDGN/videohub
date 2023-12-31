package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"videohub/util"
)

type VideosHandler struct {
	MongoDb *util.MongoDB
}

func (vh *VideosHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page <= 0 || pageSize <= 0 {
		http.Error(w, "Invalid page or pageSize parameter", http.StatusBadRequest)
		return
	}

	videos, err := vh.MongoDb.GetAllVideosPaged(page, pageSize)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(map[string]interface{}{"videos": videos})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
