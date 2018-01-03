package dotMvc

import (
	"html/template"
	"log"
	"net/http"
)

//IController :
type IController interface {
	actionInvoker(w http.ResponseWriter, r *http.Request, vars map[string]string)
	BeginExecute()
	OnAuthentication() bool
	OnAuthorization() bool
	OnActionExecuting()
	OnActionExecuted()
	OnResultExecuting()
	OnResultExecuted()
	OnException(err interface{})
	EndExecute()
	Get() IActionResult
	Post() IActionResult
	Put() IActionResult
	Delete() IActionResult
	Head() IActionResult
	Patch() IActionResult
	Options() IActionResult
}

//Controller :
type Controller struct {
	IController
	ResponseWriter *http.ResponseWriter
	Request        *http.Request
	Vars           map[string]string
}

//New for Route
func (c *Controller) actionInvoker(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	log.Println(r.Method, r.Header, *r)
	c.Vars = vars
	c.ResponseWriter = &w
	c.Request = r
}

//View :
func (c *Controller) View(viewName string, model interface{}) IActionResult {
	t, _ := template.ParseFiles("views/" + viewName + ".html") // 讀取 HTML 檔案
	t.Execute(*c.ResponseWriter, model)
	return ActionResult{}
}

//HTTPNotFound :
func (c *Controller) HTTPNotFound() IActionResult {
	http.Error(*c.ResponseWriter, httpStatusCode[404], 404)
	return ActionResult{}
}

//BeginExecute :
func (c *Controller) BeginExecute() {

}

//OnAuthentication :
func (c *Controller) OnAuthentication() bool {
	return true
}

//OnAuthorization :
func (c *Controller) OnAuthorization() bool {
	return true
}

//OnActionExecuting :
func (c *Controller) OnActionExecuting() {

}

//OnActionExecuted :
func (c *Controller) OnActionExecuted() {

}

//OnResultExecuting :
func (c *Controller) OnResultExecuting() {

}

//OnResultExecuted :
func (c *Controller) OnResultExecuted() {

}

//OnException :
func (c *Controller) OnException(err interface{}) {

}
