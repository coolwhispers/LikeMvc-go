package LikeMvc

import (
	"log"
	"net/http"
	"os"
)

var customRouter map[string]func(w http.ResponseWriter, r *http.Request)

//Run : Start Web Server
func Run() {
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

	log.Println("Server Start")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
