package controller

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	valid "github.com/asaskevich/govalidator"
	"github.com/noppawitt/admintools/model"
	"github.com/noppawitt/admintools/service"
	"github.com/noppawitt/admintools/util"
)

// ApplicationController is an application controller
type ApplicationController struct {
	*Controller
	Service service.ApplicationService
}

// NewApplicationController returns application controller
func NewApplicationController(s service.ApplicationService) *ApplicationController {
	return &ApplicationController{Service: s}
}

// Router returns application router
func (c *ApplicationController) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", c.Create)
	r.Get("/", c.Get)
	r.Get("/{id}", c.GetByID)
	r.Get("/page/{page}", c.GetPage)
	r.Put("/{id}", c.Update)
	r.Delete("/{id}", c.Delete)
	r.Get("/{id}/sp", c.GetStoredProcedure)
	r.Get("/{id}/view", c.GetView)
	r.Get("/{id}/sp/{name}/parameter", c.GetParameter)
	r.Get("/{id}/view/{name}/parameter", c.GetParameter)
	return r
}

// Create creates an application
func (c *ApplicationController) Create(w http.ResponseWriter, r *http.Request) {
	application := model.Application{}
	if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
		panic(err)
	}
	_, err := valid.ValidateStruct(&application)
	if err != nil {
		c.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := c.Service.Create(&application); err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, &application)
}

// Get returns applications
func (c *ApplicationController) Get(w http.ResponseWriter, r *http.Request) {
	applications, err := c.Service.Repo().Find()
	if err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, applications)
}

// GetByID returns an application with giving id
func (c *ApplicationController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	application, err := c.Service.Repo().FindOne(id)
	if err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, application)
}

// GetPage returns application page
func (c *ApplicationController) GetPage(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Applications []model.Application `json:"applications"`
		Count        int                 `json:"count"`
		TotalPage    int                 `json:"totalPage"`
	}
	pageSize := 10
	var err error
	name := r.FormValue("name")
	dbName := r.FormValue("dbName")
	sortBy := util.ToSnakeCase(r.FormValue("sortBy"))
	if sortBy == "" {
		sortBy = "id"
	}
	direction := r.FormValue("direction")
	if direction == "" {
		direction = "desc"
	}
	page, err := strconv.Atoi(chi.URLParam(r, "page"))
	if err != nil {
		panic(err)
	}
	if r.FormValue("pageSize") != "" {
		pageSize, err = strconv.Atoi(r.FormValue("pageSize"))
		if err != nil {
			panic(err)
		}
	}
	applications, err := c.Service.Repo().Paginate(name, dbName, sortBy, direction, pageSize, (page-1)*pageSize)
	if err != nil {
		panic(err)
	}
	count, err := c.Service.Repo().Count(name, dbName)
	if err != nil {
		panic(err)
	}
	response := Response{
		Applications: applications,
		Count:        count,
		TotalPage:    int(math.Ceil(float64(count) / float64(pageSize))),
	}
	c.JSON(w, http.StatusOK, response)
}

// Update updates an application
func (c *ApplicationController) Update(w http.ResponseWriter, r *http.Request) {
	request := model.Application{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		panic(err)
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	application, err := c.Service.Repo().FindOne(id)
	if err != nil {
		panic(err)
	}
	application.Name = request.Name
	application.Host = request.Host
	application.Port = request.Port
	application.Username = request.Username
	application.Password = request.Password
	application.DBName = request.DBName
	if err := c.Service.Repo().Save(application); err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, application)
}

// Delete deletes an application
func (c *ApplicationController) Delete(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Message string `json:"message"`
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	c.Service.Repo().Remove(id)
	response := Response{
		Message: "Success",
	}
	c.JSON(w, http.StatusOK, response)
}

// GetStoredProcedure returns external stored procedure
func (c *ApplicationController) GetStoredProcedure(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	storedProcedures, err := c.Service.GetStoredProcedure(id)
	if err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, storedProcedures)
}

// GetView returns external views
func (c *ApplicationController) GetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	views, err := c.Service.GetView(id)
	if err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, views)
}

// GetParameter returns external parameters
func (c *ApplicationController) GetParameter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	name := chi.URLParam(r, "name")
	parameters, err := c.Service.GetParameter(id, name)
	if err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, parameters)
}
