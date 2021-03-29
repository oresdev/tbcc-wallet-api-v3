package router

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	h "github.com/oresdev/tbcc-wallet-api-v3/internal/controller"
	rsa "github.com/oresdev/tbcc-wallet-api-v3/internal/middleware"
)

// HTTP middleware setting a value on the request
func RSAmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if err := rsa.CheckRSA(r); err != nil {

			http.Error(w, "CheckRSA err", http.StatusUnauthorized)
		} else {
			// Assuming authentication passed, run the original handler
			next.ServeHTTP(w, r)
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
	})
}

func CreateHTTPHandler(db *sql.DB) (http.Handler, error) {
	mux := chi.NewMux()

	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(200)))
	})

	mux.Route("/api/v3", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Use(cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           30,
		}).Handler)
		// r.Use(RSAmiddleware) //RSAmiddleware

		r.Mount("/users", UserHandler(db))
		r.Mount("/inner-update", UpdateHandler(db))

		r.Mount("/vpn", VpnHandler(db))       //development routes
		r.Mount("/config", ConfigHandler(db)) //development routes

	})

	return mux, nil
}

func UserHandler(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/{uid}", h.GetUserHandler(db))
		r.Get("/", h.GetUsersHandler(db))
		r.Post("/", h.CreateUserHandler(db)) //development routes
		// todo POST users migration
	})

	return r
}

func UpdateHandler(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/", h.GetUpdateHandler(db))
		r.Post("/", h.CreateUpdateHandler(db))
	})

	return r
}

// development routes
func VpnHandler(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/", h.CreateVpnKeyHandler(db))
	})

	return r
}

// development routes
func ConfigHandler(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/", h.CreateConfigHandler(db))
	})

	return r
}
