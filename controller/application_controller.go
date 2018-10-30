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
	r.Post("/", c.create)
	r.Get("/", c.get)
	r.Get("/{id}", c.getByID)
	r.Get("/page/{page}", c.getPage)
	r.Put("/{id}", c.update)
	r.Delete("/{id}", c.delete)
	r.Get("/{id}/sp", c.getStoredProcedure)
	r.Get("/{id}/view", c.getView)
	r.Get("/{id}/sp/{name}/parameter", c.getParameter)
	r.Get("/{id}/view/{name}/parameter", c.getParameter)
	return r
}

func (c *ApplicationController) create(w http.ResponseWriter, r *http.Request) {
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

func (c *ApplicationController) get(w http.ResponseWriter, r *http.Request) {
	applications, err := c.Service.Repo().Find()
	if err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, applications)
}

func (c *ApplicationController) getByID(w http.ResponseWriter, r *http.Request) {
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

func (c *ApplicationController) getPage(w http.ResponseWriter, r *http.Request) {
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

func (c *ApplicationController) update(w http.ResponseWriter, r *http.Request) {
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

func (c *ApplicationController) delete(w http.ResponseWriter, r *http.Request) {
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

func (c *ApplicationController) getStoredProcedure(w http.ResponseWriter, r *http.Request) {
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

func (c *ApplicationController) getView(w http.ResponseWriter, r *http.Request) {
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

func (c *ApplicationController) getParameter(w http.ResponseWriter, r *http.Request) {
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
