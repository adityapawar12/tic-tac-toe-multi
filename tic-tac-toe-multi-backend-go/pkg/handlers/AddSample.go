package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"

    "sample.app/pkg/mocks"
    "sample.app/pkg/models"
)

func AddSample(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    log.Printf("Request Body: %s\n", body)

    var sample models.Sample
    json.Unmarshal(body, &sample)

    log.Printf("Request sample: %s\n", sample)

    sample.Id = rand.Intn(100)
    mocks.Samples = append(mocks.Samples, sample)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode("Created")
}
