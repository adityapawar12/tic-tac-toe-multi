package handlers

import (
    "encoding/json"
    "net/http"

    "sample.app/pkg/mocks"
)

func GetAllSamples(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(mocks.Samples)
}
