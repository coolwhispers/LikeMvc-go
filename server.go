package dotMvc

import (
	"log"
	"net/http"
	"os"
)

var customRouter map[string]func(w http.ResponseWriter, r *http.Request)

var srv *http.Server

//Start Web Server
func Start() *http.Server {

	srv = &http.Server{Addr: ":8080"}

	folders := []string{"conf", "controllers", "models", "static", "utils", "views"}

	for _, path := range folders {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, os.ModePerm)
		}
	}

	for k, v := range customRouter {
		http.HandleFunc(k, v)
	}

	if _, ok := customRouter["/"]; ok {
		http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	} else {
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	}

	if err := srv.ListenAndServe(); err != nil {
		// cannot panic, because this probably is an intentional close
		log.Printf("Httpserver: ListenAndServe() error: %s", err)
	} else {
		log.Println("Server Start")
	}

	return srv
}

//Stop Web Server
func Stop() {
	if err := srv.Shutdown(nil); err != nil {
		panic(err)
	}
}
