package dotMvc

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var servers = make(map[string]*Server)

//Server : HTTP Server
type Server struct {
	s *http.Server
}

//Run :
func Run(addr string, r *mux.Router) *Server {
	if _, ok := servers[addr]; ok {
		panic("addr is exist")
	}

	srv := &http.Server{Addr: addr, Handler: r}

	if err := srv.ListenAndServe(); err != nil {
		// cannot panic, because this probably is an intentional close
		log.Printf("Httpserver: ListenAndServe() error: %s", err)
	} else {
		log.Println("Server Start")
	}

	// if _, ok := customRouter["/"]; ok {
	// 	muxRouter.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// } else {
	// 	muxRouter.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	// }

	servers[addr] = &Server{srv}

	return servers[addr]
}
