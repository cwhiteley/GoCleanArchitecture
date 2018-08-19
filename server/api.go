package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"recipes/repository"
)

func StartAPIServer() {
	router := Router(&Dependencies{repository.NewRecipeRepository()})

	server := &http.Server{Addr: fmt.Sprintf(":%s", os.Getenv("PORT")), Handler: router}
	go listenServer(server)
	waitForShutdown(server)
}

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		fmt.Println(fmt.Sprintf("Error in starting server - %s", err))
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	fmt.Println("API server shutting down")
	// Finish all apis being served and shutdown gracefully
	apiServer.Shutdown(context.Background())
	fmt.Println("API server shutdown complete")
}
