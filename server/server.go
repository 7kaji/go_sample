package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	user "go_sample/controller"
)

// Init is initialize server
func Init() {
	router := chi.NewRouter()
	ctrl := user.Controller{}

	// Set output for logging
	middleware.DefaultLogger = middleware.RequestLogger(
		&middleware.DefaultLogFormatter{
			Logger: newLogger(),
		},
	)
	router.Use(middleware.Logger)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	// router.HandleFunc("/*", Index) // WildCard.

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/users", ctrl.Index)
		r.Get("/users/{userID}", ctrl.Show)
		r.Post("/users", ctrl.Create)
		r.Patch("/users/{userID}", ctrl.Update)
		r.Delete("/users/{userID}", ctrl.Delete)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "chi-log: ", log.Lshortfile)
}
