package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"github.com/rezzamaqfiro/wallet/api/order"
	"github.com/rezzamaqfiro/wallet/middleware"

	repo "github.com/rezzamaqfiro/wallet/repo/generated"
	"github.com/rezzamaqfiro/wallet/util"
)

type Handler struct {
	router *chi.Mux
}

type Middleware func(http.Handler) http.Handler

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func New(db *sql.DB, rdb *redis.Client, cors Middleware) *Handler {
	r := chi.NewMux()
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.BirthTime)
	// r.Use(chiMiddleware.RequestID)
	r.Use(cors)

	h := &Handler{
		router: r,
	}

	// init
	validate = validator.New(validator.WithRequiredStructEnabled())
	dbGenerated := repo.New(db)

	r.Get("/ping", h.Ping)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		util.NewResponse(http.StatusNotFound, http.StatusNotFound, "404 Not found!", nil).WriteResponse(w, r)
	})

	orderHandler := order.NewHandler(dbGenerated)
	r.Route("/v1/order", func(r chi.Router) {
		r.Post("/disbursement", orderHandler.OrderDisbursement)
	})

	return h
}

func (h *Handler) Handler() http.Handler {
	return h.router
}
