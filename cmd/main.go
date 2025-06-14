package main

import (
	// "bytes"
	"context"
	// "github.com/yuin/goldmark"
	// "github.com/yuin/goldmark/renderer/html"
	"log"
	"net/http"
	"os"
	"os/signal"
	// "personal/cmd/internal"
	"personal/cmd/internal/routing"
	"syscall"
	"time"
)

func main() {
	// md := goldmark.New(
	// 	goldmark.WithRendererOptions(html.WithUnsafe()),
	// 	goldmark.WithExtensions(&internal.MDComponent{}))

	prod_env := os.Getenv("PRODUCTION_ENV")
	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if prod_env == "" {
			w.Header().Set("Cache-Control", "no-store")
		}
		fs.ServeHTTP(w, r)
	})))
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		routing.IndexPage().ServeHTTP(w, r)
	})

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
			log.Fatalf("Server exited with error: %s\n", err)
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
