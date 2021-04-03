package api

import (
	"fmt"
	"log"
	"net/http"
	"projects/go-rest-api/src/api/router"
)

func Run()  {
	fmt.Println("\n\tListening [::]:3000\n")

	r:= router.New()
	log.Fatal(http.ListenAndServe(":3000", r))
}