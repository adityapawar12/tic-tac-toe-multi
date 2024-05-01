package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "sample.app/pkg/handlers"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/api/v1/sample", handlers.GetAllSamples).Methods(http.MethodGet)
    router.HandleFunc("/api/v1/sample", handlers.AddSample).Methods(http.MethodPost)
    router.HandleFunc("/api/v1/sample/{id}", handlers.GetSample).Methods(http.MethodGet)
    router.HandleFunc("/api/v1/sample/{id}", handlers.UpdateSample).Methods(http.MethodPut)
    router.HandleFunc("/api/v1/sample/{id}", handlers.DeleteSample).Methods(http.MethodDelete)

    log.Println("Server is running on http://localhost:3000")
    http.ListenAndServe(":3000", router)
}
