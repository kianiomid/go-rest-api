package api

import (
	"fmt"
	"log"
	"net/http"
	"projects/go-rest-api/src/api/router"
	"projects/go-rest-api/src/auto"
	"projects/go-rest-api/src/config"
)

func Run() {
	config.Load()
	auto.Load()

	fmt.Println("\n\tListening [::]:", config.PORT)

	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
