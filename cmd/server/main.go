package main

import (
	"SGIMS/internal/config"
	"SGIMS/internal/handlers"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/set", handlers.Set)
	http.HandleFunc("/get", handlers.Get)
	http.HandleFunc("/delete", handlers.Delete)
	http.HandleFunc("/getall", handlers.GetAll)

	fmt.Println("Starting server at port", config.PORT)
	http.ListenAndServe(":"+strconv.Itoa(config.PORT), nil)
}
