package LikeMvc

import (
	"html/template"
	"log"
	"net/http"
)

//IController :
type IController interface {
	new(w http.ResponseWriter, r *http.Request)
	Get()
	Post()
	Put()
	Delete()
	Head()
	Patch()
	Options()
}

//Controller :
type Controller struct {
	ResponseWriter *http.ResponseWriter
	Request        *http.Request
}

//New for Route
func (c *Controller) new(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.Header, *r)
	c.ResponseWriter = &w
	c.Request = r
}

//View :
func (c *Controller) View(viewName string, model interface{}) {
	t, _ := template.ParseFiles("views/" + viewName + ".html") // 讀取 HTML 檔案
	t.Execute(*c.ResponseWriter, model)
}
