package dotMvc

import "net/http"

//Hub : WebStocket Hub
type hub struct {
	client *map[string]hubClient
}

func (h *hub) Execute(w http.ResponseWriter, r *http.Request) {

}

type hubClient struct {
}
