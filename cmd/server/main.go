package main

import (
	"log"
	"net/http"
	"os"

	"alquimia-alimentar/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Static assets (CSS, JS, images, video).
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Routes.
	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("GET /en/{$}", handlers.Home)
	mux.HandleFunc("GET /nl/{$}", handlers.Home)

	// Redirect language roots without trailing slash.
	mux.HandleFunc("GET /en", redirect("/en/"))
	mux.HandleFunc("GET /nl", redirect("/nl/"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started at http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func redirect(to string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, to, http.StatusMovedPermanently)
	}
}
