package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"personal/cmd/internal/routing"
	"syscall"
	"time"
)

func main() {
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fs))
	router.HandleFunc("GET /", routing.IndexPage)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting server...")
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Printf("Server exited with error: %s\n", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("Server exited with code 1.")
	}
}
