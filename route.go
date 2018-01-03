package dotMvc

import (
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

var muxRouter *mux.Router

func init() {
	muxRouter = mux.NewRouter()
}

//AddRoutes :
func AddRoutes(s string, controllerType reflect.Type) {

	if _, ok := controllerType.(IController); !ok {
		panic("Not dotMvc controller")
	}

	if customRouter == nil {
		customRouter = make(map[string]func(w http.ResponseWriter, r *http.Request))
	}
	rt := routeTemplate{controllerType}
	customRouter[s] = rt.run
}

type routeTemplate struct {
	controllerType reflect.Type
}

func resultAction(actionResult *IActionResult) {

}

var (
	httpStatusCode = map[int]string{
		403: "403 Forbidden",
		404: "404 Not Found",
		405: "405 Method Not Allowed",
		500: "500 Internal Server Error",
	}
)

func (rt *routeTemplate) run(w http.ResponseWriter, r *http.Request) {
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c := reflect.New(rt.controllerType).Elem().Interface().(IController)
	defer func() {
		if err := recover(); err != nil {
			c.OnException(err)
		}

		http.Error(w, httpStatusCode[500], 500)
	}()

	vars := mux.Vars(r)
	c.actionInvoker(w, r, vars)

	c.BeginExecute()
	if !c.OnAuthentication() {

	}

	if !c.OnAuthorization() {

	}

	c.OnActionExecuting()

	var actionResult IActionResult

	switch r.Method {
	case "GET":
		actionResult = c.Get()
	case "POST":
		actionResult = c.Post()
	case "PUT":
		actionResult = c.Put()
	case "DELETE":
		actionResult = c.Delete()
	case "HEAD":
		actionResult = c.Head()
	case "PATCH":
		actionResult = c.Patch()
	case "OPTIONS":
		actionResult = c.Options()
	default:
		actionResult = methodNotAllowed(&w)
	}

	c.OnActionExecuted()

	c.EndExecute()

	c.OnResultExecuting()

	resultAction(&actionResult)

	c.OnResultExecuted()
}

func methodNotAllowed(w *http.ResponseWriter) IActionResult {
	http.Error(*w, httpStatusCode[405], 405)
	return ActionResult{}
}
