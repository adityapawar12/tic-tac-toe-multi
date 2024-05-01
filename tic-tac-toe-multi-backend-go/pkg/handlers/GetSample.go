package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "sample.app/pkg/mocks"
)

func GetSample(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    for _, sample := range mocks.Samples {
        if sample.Id == id {
            // If ids are equal send book as a response
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)

            json.NewEncoder(w).Encode(sample)
            break
        }
    }
}
