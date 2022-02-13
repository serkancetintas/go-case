package api

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"go-case/internal/infra/store"
	"go-case/pkg"
	"log"
	"net/http"
	"os"
)


func NewServer() {
	handler := &Handler{ Store: store.NewStore() }

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	router := http.NewServeMux()

	basePath := "/api"
	router.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	router.HandleFunc(basePath+ "/set", handler.SetHandler)
	router.HandleFunc(basePath+ "/get", handler.GetHandler)
	router.HandleFunc(basePath+"/flush", handler.FlushHandler)

	listenAddr := ":8080"
	server := &http.Server{
		Addr:         listenAddr,
		Handler:      pkg.Logging(logger)(router),
		ErrorLog:     logger,
	}

	logger.Println("Server is ready to handle requests at", listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}
}

