package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"uralmelal/backend/internal/config"
	"uralmelal/backend/internal/database"
	"uralmelal/backend/internal/handlers"
	"uralmelal/backend/internal/middleware"
	"uralmelal/backend/internal/models"
	"uralmelal/backend/internal/repository"
	"uralmelal/backend/internal/services"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg := config.Load()
	db, err := database.NewDB(cfg)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.ProductAttribute{},
		&models.ContentBlock{}, &models.Lead{}, &models.AppliedMigration{}); err != nil {
		slog.Error("migration failed", "error", err)
		os.Exit(1)
	}
	if err := database.SeedAdmin(db, cfg); err != nil {
		slog.Error("seed failed", "error", err)
		os.Exit(1)
	}
	if err := database.SeedContentBlocks(db); err != nil {
		slog.Error("content seed failed", "error", err)
		os.Exit(1)
	}

	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService)
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db, cfg.DBType)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)
	productHandler := handlers.NewProductHandler(productRepo)
	uploadHandler := handlers.NewUploadHandler(cfg)
	productIOHandler := handlers.NewProductIOHandler(productRepo, categoryRepo, db)
	contentRepo := repository.NewContentRepository(db)
	leadRepo := repository.NewLeadRepository(db)
	contentHandler := handlers.NewContentHandler(contentRepo)
	leadHandler := handlers.NewLeadHandler(leadRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	authMW := middleware.Auth(cfg.JWTSecret)
	adminOnly := middleware.RequireRole(models.RoleAdmin)
	editorOrAdmin := middleware.RequireRole(models.RoleEditor, models.RoleAdmin)
	managerOrAdmin := middleware.RequireRole(models.RoleManager, models.RoleAdmin)

	r := chi.NewRouter()
	r.Use(middleware.Logging)
	r.Use(middleware.CORS)
	r.Use(chimw.Recoverer)
	r.Use(chimw.RealIP)

	r.Get("/api/health", handlers.Health)

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.With(authMW).Get("/me", authHandler.Me)
	})

	r.Get("/api/categories", categoryHandler.GetTree)
	r.Get("/api/products", productHandler.List)
	r.Get("/api/products/{id}", productHandler.Get)
	r.Get("/api/content/{page}", contentHandler.GetByPage)
	r.With(middleware.RateLimit(5, 60*time.Second)).Post("/api/leads", leadHandler.Create)

	r.Route("/api/admin", func(r chi.Router) {
		r.Use(authMW)
		r.Use(editorOrAdmin)
		r.Route("/categories", func(r chi.Router) {
			r.Post("/", categoryHandler.Create)
			r.Put("/{id}", categoryHandler.Update)
			r.Delete("/{id}", categoryHandler.Delete)
		})
		r.Route("/products", func(r chi.Router) {
			r.Post("/", productHandler.Create)
			r.Put("/{id}", productHandler.Update)
			r.Delete("/{id}", productHandler.Delete)
			r.Get("/export", productIOHandler.Export)
			r.Post("/import", productIOHandler.Import)
		})
		r.Post("/upload", uploadHandler.Upload)
		r.Get("/content", contentHandler.GetByPageAdmin)
		r.Put("/content/{id}", contentHandler.Update)
	})
	r.Route("/api/admin/leads", func(r chi.Router) {
		r.Use(authMW)
		r.Use(managerOrAdmin)
		r.Get("/", leadHandler.List)
		r.Patch("/{id}", leadHandler.UpdateStatus)
	})
	r.Route("/api/admin/users", func(r chi.Router) {
		r.Use(authMW)
		r.Use(adminOnly)
		r.Get("/", userHandler.List)
		r.Post("/", userHandler.Create)
		r.Put("/{id}", userHandler.Update)
		r.Delete("/{id}", userHandler.Delete)
	})

	addr := ":" + cfg.AppPort
	slog.Info("server starting", "addr", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
