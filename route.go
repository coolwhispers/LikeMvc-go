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
		panic("not a controller")
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

func (rt *routeTemplate) run(w http.ResponseWriter, r *http.Request) {
	c := reflect.New(rt.controllerType).Elem().Interface().(IController)
	defer func() {
		if err := recover(); err != nil {
			c.OnException(err)
		}
	}()

	c.actionInvoker(w, r)

	c.BeginExecute()
	if !c.OnAuthentication() {

	}

	if !c.OnAuthorization() {

	}

	c.OnActionExecuting()

	switch r.Method {
	case "GET":
		c.Get()
	case "POST":
		c.Post()
	case "PUT":
		c.Put()
	case "DELETE":
		c.Delete()
	case "HEAD":
		c.Head()
	case "PATCH":
		c.Patch()
	case "OPTIONS":
		c.Options()
	}
	c.OnActionExecuted()

	c.EndExecute()

	c.OnResultExecuting()
	c.OnResultExecuted()
}
