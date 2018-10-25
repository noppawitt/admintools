package controller

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/noppawitt/admintools/util"

	"github.com/go-chi/chi"
	"github.com/noppawitt/admintools/model"

	"github.com/noppawitt/admintools/service"
)

// FunctionController is a function controller
type FunctionController struct {
	*Controller
	Service          service.FunctionService
	ParameterService service.ParameterService
}

// NewFunctionController returns function controller
func NewFunctionController(s service.FunctionService, p service.ParameterService) *FunctionController {
	return &FunctionController{Service: s, ParameterService: p}
}

// Router returns function router
func (c *FunctionController) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", c.Create)
	r.Get("/{id}", c.GetByID)
	r.Get("/page/{page}", c.GetPage)
	r.Put("/{id}", c.Update)
	r.Delete("/{id}", c.Delete)
	r.Get("/{id}/parameter", c.GetParameter)
	return r
}

// Create creates a function
func (c *FunctionController) Create(w http.ResponseWriter, r *http.Request) {
	function := model.Function{}
	if err := json.NewDecoder(r.Body).Decode(&function); err != nil {
		panic(err)
	}
	c.Service.Repo().Create(&function)
	c.JSON(w, http.StatusOK, &function)
}

// GetByID returns a function
func (c *FunctionController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	function, err := c.Service.Repo().FindOneIncludeApplication(id)
	if err != nil {
		panic(err)
	}
	c.JSON(w, http.StatusOK, &function)
}

// GetPage returns function page
func (c *FunctionController) GetPage(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Functions []model.Function `json:"functions"`
		Count     int              `json:"count"`
		TotalPage int              `json:"totalPage"`
	}
	pageSize := 10
	var err error
	name := r.FormValue("name")
	appName := r.FormValue("appName")
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
	functions, err := c.Service.Repo().Paginate(name, appName, sortBy, direction, pageSize, (page-1)*pageSize)
	if err != nil {
		panic(err)
	}
	count, err := c.Service.Repo().Count(name, appName)
	if err != nil {
		panic(err)
	}
	response := Response{
		Functions: functions,
		Count:     count,
		TotalPage: int(math.Ceil(float64(count) / float64(pageSize))),
	}
	c.JSON(w, http.StatusOK, &response)
}

// Update updates a function
func (c *FunctionController) Update(w http.ResponseWriter, r *http.Request) {
	request := model.Function{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		panic(err)
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	function, err := c.Service.Repo().FindOne(id)
	if err != nil {
		panic(err)
	}
	function.Name = request.Name
	function.Remarks = request.Remarks
	function.AllowMultiple = request.AllowMultiple
	function.StoredProcedureName = request.StoredProcedureName
	function.ViewName = request.ViewName
	c.Service.Repo().Save(function)
	c.JSON(w, http.StatusOK, function)
}

// Delete deletes a function
func (c *FunctionController) Delete(w http.ResponseWriter, r *http.Request) {
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

// GetParameter returns parameters
func (c *FunctionController) GetParameter(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		panic(err)
	}
	parameters, err := c.ParameterService.Repo().FindByFunctionID(id)
	c.JSON(w, http.StatusOK, parameters)
}
