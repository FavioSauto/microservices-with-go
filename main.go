package main

import (
	"log"
	"net/http"
	"os"
	"github.com/FavioSauto/microservices-with-go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := NewServerMux()
	sm.Handle('/', hh)

	

	http.ListenAndServe(":3000", sm)
}