package router

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	h "github.com/oresdev/tbcc-wallet-api-v3/internal/server/controller"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/middleware/rsa"
)

// CreateHTTPHandler ...
func CreateHTTPHandler(db *sql.DB) (http.Handler, error) {
	mux := chi.NewMux()

	mux.Get("/api/v3/ping", func(w http.ResponseWriter, r *http.Request) {
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

		r.Use(rsa.CheckRSASignature) // CheckRSASignature

		r.Mount("/users", UserHandler(db))
		r.Mount("/vpn", VpnHandler(db))
		r.Mount("/app", AppHandler(db))

	})

	return mux, nil
}

// UserHandler ...
func UserHandler(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/", h.GetUsersHandler(db))
		r.Get("/{uuid}", h.GetUserHandler(db))
		r.Get("/ext/{uuid}", h.GetExtendedUserHandler(db))
		r.Post("/{uuid}/update", h.UpdateUserHandler(db))
		r.Post("/", h.CreateUserHandler(db)) // TODO remove development routes

		// Migrate user data from depricated database (public scheme)
		// Returns extended user data
		r.Post("/migrate", h.MigrateUserHandler(db))
	})

	return r
}

// TODO remove development handler
// VpnHandler ...
func VpnHandler(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/", h.CreateVpnKeyHandler(db))
	})

	return r
}

// AppHandler ...
func AppHandler(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/config", h.CreateConfigHandler(db)) // TODO remove development routes
		r.Get("/config", h.GetConfigHandler(db))
		r.Post("/update", h.CreateUpdateHandler(db)) // TODO remove development routes
		r.Get("/update", h.GetUpdateHandler(db))
	})

	return r
}
