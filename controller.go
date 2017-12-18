package LikeMvc

import (
	"html/template"
	"net/http"
)

//MvcControllerInterface :
type MvcControllerInterface interface {
	new(w http.ResponseWriter, r *http.Request)
	Get()
	Post()
	Put()
	Delete()
	Head()
	Patch()
	Options()
}

//MvcController :
type MvcController struct {
	ResponseWriter *http.ResponseWriter
	Request        *http.Request
}

//New for Route
func (c *MvcController) new(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = &w
	c.Request = r
}

//View :
func (c *MvcController) View(viewName string, model interface{}) {
	t, _ := template.ParseFiles("Views/" + viewName + ".html") // 讀取 HTML 檔案
	t.Execute(*c.ResponseWriter, model)
}
