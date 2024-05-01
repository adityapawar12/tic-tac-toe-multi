package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "sample.app/pkg/mocks"
)

func DeleteSample(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    for index, sample := range mocks.Samples {
        if sample.Id == id {
            mocks.Samples = append(mocks.Samples[:index], mocks.Samples[index+1:]...)

            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode("Deleted")
            break
        }
    }
}
