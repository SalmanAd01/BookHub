package bin

import (
	"github.com/gorilla/mux"
)

func CreateServer() *mux.Router {
	server := mux.NewRouter().StrictSlash(true)
	return server
}
