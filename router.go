package dotMvc

import (
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type routerControl struct {
	t reflect.Type
}

func (c *routerControl) hundler(w http.ResponseWriter, r *http.Request) {

	for k, v := range Header {
		w.Header().Set(k, v)
	}

	control := reflect.New(c.t).Elem().Interface().(IController)

	defer func() {
		if err := recover(); err != nil {
			control.OnException(err)
		}

		http.Error(w, httpStatusCode[500], 500)
	}()

	vars := mux.Vars(r)

	control.BeginExecute()

	control.actionInvoker(w, r, vars)

	control.OnAuthentication()

	control.OnAuthorization()

	control.OnActionExecuting()

	var actionResult IActionResult

	switch r.Method {
	case "GET":
		actionResult = control.Get()
	case "POST":
		actionResult = control.Post()
	case "PUT":
		actionResult = control.Put()
	case "DELETE":
		actionResult = control.Delete()
	case "HEAD":
		actionResult = control.Head()
	case "PATCH":
		actionResult = control.Patch()
	case "OPTIONS":
		actionResult = control.Options()
	default:
		actionResult = methodNotAllowed(&w)
	}

	control.OnActionExecuted()

	control.EndExecute()

	control.OnResultExecuting()

	returnAction(actionResult)

	control.OnResultExecuted()
}

func methodNotAllowed(w *http.ResponseWriter) IActionResult {
	http.Error(*w, httpStatusCode[405], 405)
	return ActionResult{}
}

func returnAction(action IActionResult) {

}
