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
	t, err := template.ParseFiles("views/" + viewName + ".html") // 讀取 HTML 檔案
	if err != nil {
		panic(err)
	}

	t.Execute(*c.ResponseWriter, model)
	return ActionResult{}
}

//ReturnCode :
func (c *Controller) ReturnCode(code int) IActionResult {
	word, ok := httpStatusCode[code]
	if !ok {
		panic("Not exist status code.")
	}
	http.Error(*c.ResponseWriter, word, code)
	return ActionResult{}
}

//ReturnMessage :
func (c *Controller) ReturnMessage(code int, message string) IActionResult {
	http.Error(*c.ResponseWriter, message, code)
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

//Get :
func (c *Controller) Get() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Post :
func (c *Controller) Post() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Put :
func (c *Controller) Put() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Delete :
func (c *Controller) Delete() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Head :
func (c *Controller) Head() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Patch :
func (c *Controller) Patch() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}

// Options :
func (c *Controller) Options() IActionResult {
	return methodNotAllowed(c.ResponseWriter)
}
