package main

import (
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/gorilla/mux"
    "io"
    "log"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    if _, err := io.WriteString(w, "Hello, world!\n"); err != nil {
        log.Printf("write error: %v", err)
    }
}

func router() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", handler)
    r.Handle("/metrics", promhttp.Handler())
    return r
}

func main() {
    router := router()
    srv := &http.Server{
        Handler:      router,
        Addr:         "0.0.0.0:8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}
