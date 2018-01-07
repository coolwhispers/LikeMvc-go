package dotMvc

import (
	"net/http"
	"os"
	"reflect"
)

var (
	httpStatusCode = map[int]string{
		403: "403 Forbidden",
		404: "404 Not Found",
		405: "405 Method Not Allowed",
		500: "500 Internal Server Error",
	}
	rc = make(map[reflect.Type]*routerControl)

	//Header : HTTP Header
	Header = make(map[string]string)
)

//New :
func New(c IController) func(w http.ResponseWriter, r *http.Request) {

	t := reflect.TypeOf(c)

	f, ok := rc[t]

	if !ok {
		f := &routerControl{t}
		rc[t] = f
	}

	return f.hundler
}

//DefaultHeader :
func DefaultHeader() {
	Header["Access-Control-Allow-Origin"] = "*"
	Header["Access-Control-Allow-Headers"] = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
}

//CreateFolder :
func CreateFolder() {
	folders := []string{"conf", "controllers", "models", "static", "utils", "views"}

	for _, path := range folders {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModePerm)
		}
	}
}
