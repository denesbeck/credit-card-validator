package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/validator"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Info("Starting ValidateHandler function...")
		vars := mux.Vars(r)
		id := vars["id"]
		log.Infof("Id: %s", id)
		checkDigit := vars["checkDigit"]
		log.Infof("Check Digit: %s", checkDigit)
		checkDigitInt, err := strconv.Atoi(checkDigit)
		if err != nil {
			log.Error(err)
		}
		result := validator.Validate(id, checkDigitInt)
		log.Infof("Result: %t", result)
		fmt.Fprintf(w, "%t", result)
		log.Info("Ending ValidateHandler function...")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/validate/{id}/{checkDigit}", ValidateHandler)

	http.Handle("/", r)

	fmt.Println("Server is listening at http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
