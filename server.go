package stalkerlauncher

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"hash"`
}

func Server() {
	hashClient := UpdateClient()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: hashClient}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe("localhost:8080", mux)
}
