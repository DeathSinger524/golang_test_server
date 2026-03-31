package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/fast", fastHandler)
	http.HandleFunc("/slow", slowHandler)
	http.HandleFunc("/random", randomHandler)
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/health", healthHandler)

	port := ":8080"
	log.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func fastHandler(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "Fast response")
}

func slowHandler(w http.ResponseWriter, r *http.Request) {
	delay := rand.Intn(2000) + 500 // 500–2500 ms
	time.Sleep(time.Duration(delay) * time.Millisecond)
	respondJSON(w, http.StatusOK, "Slow response")
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(100)

	if n < 20 {
		http.Error(w, "Random error occurred", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, "Random success")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		msg = "empty"
	}
	respondJSON(w, http.StatusOK, msg)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func respondJSON(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Status:  http.StatusText(status),
		Message: message,
		Time:    time.Now().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(resp)
}
