package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/noppawitt/admintools/util"

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
	// ENV is an application's environment
	ENV string
	err error
)

func init() {
	flag.StringVar(&ENV, "env", "development", "Application's environment")
	flag.Parse()
	fmt.Printf("Server is using %s environment\n", ENV)

	// Register custom validator
	util.RegisterCustomValidator()
}

func main() {
	// Init configuration
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
	applicationRepository := repository.NewApplicationRepository(db)
	functionRepository := repository.NewFunctionRepository(db)
	parameterRepository := repository.NewParameterRepository(db)
	externalRepository := repository.NewExternalRepository()

	// Service
	authService := service.NewAuthService(cfg)
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

	// FIXME: Move this to sub directory
	FileServer(r, "/css", http.Dir("dist/css"))
	FileServer(r, "/fonts", http.Dir("dist/fonts"))
	FileServer(r, "/img", http.Dir("dist/img"))
	FileServer(r, "/js", http.Dir("dist/js"))
	FileServer(r, "/favicon.ico", http.Dir("dist/favacon.ico"))
	index := template.Must(template.ParseFiles("dist/index.html"))
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		index.Execute(w, nil)
	})
	tmpl := template.Must(template.ParseFiles("view/login.html"))
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})
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

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
