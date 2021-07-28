package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ThorstenHans/bridge-demo/pkg/handlers"
	"github.com/ThorstenHans/bridge-demo/pkg/middlewares"
	"github.com/ThorstenHans/bridge-demo/pkg/store"
)

const DefaultPort int = 8080

func main() {
	l := log.Default()
	l.SetOutput(os.Stdout)
	s := store.NewStore()

	productHandler := &handlers.Product{
		Log:   l,
		Store: s,
	}

	productsHandler := &handlers.Products{
		Log:   l,
		Store: s,
	}

	http.Handle("/api/products/", middlewares.Json(productHandler))
	http.Handle("/api/products", middlewares.Json(productsHandler))
	http.HandleFunc("/", middlewares.NotFound)

	port := getPort()
	l.Printf("HTTP Server will start at 0.0.0.0:%d", port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		log.Fatal("Error while starting HTTP Server")
	}
}

func getPort() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return DefaultPort
	}
	return port

}
