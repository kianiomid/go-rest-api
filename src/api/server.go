package api

import (
	"api/router"
	"auto"
	"config"
	"fmt"
	"log"
	"net/http"
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
