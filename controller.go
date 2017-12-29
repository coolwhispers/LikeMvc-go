package LikeMvc

import (
	"html/template"
	"log"
	"net/http"
)

//IController :
type IController interface {
	actionInvoker(w http.ResponseWriter, r *http.Request)
	BeginExecute()
	OnAuthentication() bool
	OnAuthorization() bool
	OnActionExecuting()
	OnActionExecuted()
	OnResultExecuting()
	OnResultExecuted()
	OnException(err interface{})
	EndExecute()
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
	IController
	ResponseWriter *http.ResponseWriter
	Request        *http.Request
}

//New for Route
func (c *Controller) actionInvoker(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.Header, *r)
	c.ResponseWriter = &w
	c.Request = r
}

//View :
func (c *Controller) View(viewName string, model interface{}) {
	t, _ := template.ParseFiles("views/" + viewName + ".html") // 讀取 HTML 檔案
	t.Execute(*c.ResponseWriter, model)
}

func (c *Controller) BeginExecute() {

}
func (c *Controller) OnAuthentication() bool {
	return true
}
func (c *Controller) OnAuthorization() bool {
	return true
}
func (c *Controller) OnActionExecuting() {

}
func (c *Controller) OnActionExecuted() {

}
func (c *Controller) OnResultExecuting() {

}
func (c *Controller) OnResultExecuted() {

}
func (c *Controller) OnException(err interface{}) {

}
