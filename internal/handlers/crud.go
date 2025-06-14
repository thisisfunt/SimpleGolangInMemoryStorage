package handlers

import (
	"SGIMS/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func Set(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Cannot read form", http.StatusBadRequest)
			return
		}
		key := r.PostFormValue("k")
		value := r.PostFormValue("v")

		if key == "" {
			http.Error(w, "\"k\" is not found", http.StatusNotFound)
			return
		}
		if value == "" {
			http.Error(w, "\"v\" is not found", http.StatusBadRequest)
			return
		}

		services.Set(key, value)
		fmt.Fprint(w, key)
	} else {
		http.Error(w, "Use POST request", http.StatusMethodNotAllowed)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Cannot read form", http.StatusBadRequest)
			return
		}
		key := r.PostFormValue("k")

		if key == "" {
			http.Error(w, "\"k\" is not found", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, services.Get(key))
	} else {
		http.Error(w, "Use POST request", http.StatusMethodNotAllowed)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Cannot read form", http.StatusBadRequest)
			return
		}
		key := r.PostFormValue("k")

		if key == "" {
			http.Error(w, "\"k\" is not found", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, key)
	} else {
		http.Error(w, "Use POST request", http.StatusMethodNotAllowed)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data := services.GetAll()
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "Use POST request", http.StatusMethodNotAllowed)
	}
}
