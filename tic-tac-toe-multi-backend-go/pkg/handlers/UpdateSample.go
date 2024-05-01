package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "sample.app/pkg/mocks"
    "sample.app/pkg/models"
)

func UpdateSample(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var updatedSample models.Sample
    json.Unmarshal(body, &updatedSample)

    for index, sample := range mocks.Samples {
        if sample.Id == id {
            sample.Text = updatedSample.Text

            mocks.Samples[index] = sample 
            w.Header().Add("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)

            json.NewEncoder(w).Encode("Updated")
            break
        }
    }
}
