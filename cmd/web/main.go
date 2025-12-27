package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"
	
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	
	"github.com/Tauhid-UAP/golang-sample-web-app/core/handlers"
	"github.com/Tauhid-UAP/golang-sample-web-app/core/middleware"
	"github.com/Tauhid-UAP/golang-sample-web-app/core/store"

	"github.com/Tauhid-UAP/golang-sample-web-app/core/redisclient"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}
	
	DATABASE_URL := os.Getenv("DATABASE_URL")
	log.Printf("DATABASE_URL: %s", DATABASE_URL)

	db, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	store.DB = db

	redisclient.Init()
	if err := redisclient.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/register", handlers.Register)
	mux.HandleFunc("/login", handlers.Login)

	// Static files
	mux.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))
	
	// Protected routes
	protected := http.NewServeMux()

	protected.HandleFunc("/logout", handlers.Logout)
	protected.HandleFunc("/profile", handlers.Profile)
	
	protectedHandler := middleware.AuthMiddleware(middleware.CSRFMiddleware(protected))

	mux.Handle("/", protectedHandler)
	
	addr := ":8000"
	server := &http.Server{
		Addr: addr,
		Handler: loggingMiddleware(mux),
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		IdleTimeout: 60*time.Second,
	}

	log.Println("Server running on ", addr)
	log.Fatal(server.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
