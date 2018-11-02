package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/noppawitt/admintools/config"
	"github.com/noppawitt/admintools/controller"
	"github.com/noppawitt/admintools/infrastructure"
	"github.com/noppawitt/admintools/middleware"
	"github.com/noppawitt/admintools/repository"
	"github.com/noppawitt/admintools/service"
	"github.com/noppawitt/admintools/util"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func init() {
	var dotEnv bool

	flag.BoolVar(&dotEnv, "dotenv", false, "Load environents from .env")
	flag.Parse()

	if dotEnv {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Register custom validator
	util.RegisterCustomValidator()
}

func main() {
	// Init config
	cfg := config.New()

	// Connect to database
	db, err := infrastructure.Connect("mssql", cfg.DBURL)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	// db.LogMode(true)
	// Migrate database
	infrastructure.AutoMigrate(db)

	// Repository
	authAgent := repository.NewAuthAgent(cfg)
	userRepository := repository.NewUserRepository(db)
	applicationRepository := repository.NewApplicationRepository(db)
	functionRepository := repository.NewFunctionRepository(db)
	parameterRepository := repository.NewParameterRepository(db)
	externalRepository := repository.NewExternalRepository()

	// Service
	authService := service.NewAuthService(authAgent, userRepository, cfg)
	applicationService := service.NewApplicationService(applicationRepository, externalRepository, cfg.EncryptionKey)
	functionService := service.NewFunctionService(functionRepository)
	parameterService := service.NewParameterService(parameterRepository)

	// Controller
	authController := controller.NewAuthController(authService)
	applicationController := controller.NewApplicationController(applicationService)
	functionController := controller.NewFunctionController(functionService, parameterService)

	// Router
	r := chi.NewRouter()
	r.Use(middleware.Cors)
	r.Use(middleware.LoggingMiddleware)

	r.Mount("/auth", authController.Router())
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.AuthVerify(cfg.Secret))
		r.Mount("/application", applicationController.Router())
		r.Mount("/function", functionController.Router())
	})

	// Serve
	port := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("Server is listening on port %d\n", cfg.Port)
	log.Fatal(http.ListenAndServe(port, r))
}
