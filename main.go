package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/noppawitt/admintools/config"
	"github.com/noppawitt/admintools/controller"
	"github.com/noppawitt/admintools/infrastructure"
	"github.com/noppawitt/admintools/middleware"
	"github.com/noppawitt/admintools/repository"
	"github.com/noppawitt/admintools/service"

	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	// ENV is a application's environment
	ENV string
	err error
)

func init() {
	flag.StringVar(&ENV, "env", "development", "Application's environment")
	flag.Parse()
	fmt.Printf("Server is using %s environment\n", ENV)
}

func main() {
	var cfg *config.Config
	if ENV == "development" {
		cfg, err = config.Dev()
	} else if ENV == "production" {
		cfg, err = config.Prod()
	} else {
		panic("Incorrect environment")
	}
	if err != nil {
		panic(err)
	}

	db, err := infrastructure.Connect("mssql", cfg.DBURL)
	defer db.Close()
	if err != nil {
		panic(err)
	}
  // db.LogMode(true)
	infrastructure.AutoMigrate(db)

	// Repository
	applicationRepository := repository.NewApplicationRepository(db)
	functionRepository := repository.NewFunctionRepository(db)
	parameterRepository := repository.NewParameterRepository(db)
	externalRepository := repository.NewExternalRepository()

	// Service
	applicationService := service.NewApplicationService(applicationRepository, externalRepository)
	functionService := service.NewFunctionService(functionRepository)
	parameterService := service.NewParameterService(parameterRepository)

	// Controller
	applicationController := controller.NewApplicationController(applicationService)
	functionController := controller.NewFunctionController(functionService, parameterService)

	// Router
	r := chi.NewRouter()
	r.Use(middleware.Cors)
	r.Use(middleware.LoggingMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Admin Tools"))
	})
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/application", applicationController.Router())
		r.Mount("/function", functionController.Router())
	})

	port := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("Server is listening on port %d\n", cfg.Port)
	log.Fatal(http.ListenAndServe(port, r))
}
