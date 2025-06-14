package handlers

import (
	"SGIMS/internal/services"
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
		key := r.FormValue("k")
		value := r.FormValue("v")

		if key == "" {
			http.Error(w, "\"k\" is not found", http.StatusBadRequest)
			return
		}
		if value == "" {
			http.Error(w, "\"v\" is not found", http.StatusBadRequest)
			return
		}

		services.Set(key, value)
		fmt.Fprint(w, key)
	} else {
		http.Error(w, "Use POST request", http.StatusNotFound)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Cannot read form", http.StatusBadRequest)
			return
		}
		key := r.FormValue("k")

		if key == "" {
			http.Error(w, "\"k\" is not found", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, services.Get(key))
	} else {
		http.Error(w, "Use POST request", http.StatusNotFound)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Cannot read form", http.StatusBadRequest)
			return
		}
		key := r.FormValue("k")

		if key == "" {
			http.Error(w, "\"k\" is not found", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, key)
	} else {
		http.Error(w, "Use POST request", http.StatusNotFound)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprint(w, services.GetAll())
	} else {
		http.Error(w, "Use POST request", http.StatusNotFound)
	}
}
