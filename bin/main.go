package bin

import (
	"github.com/gorilla/mux"
)

func CreateServer() *mux.Router {
	mux := mux.NewRouter().StrictSlash(true)
	return mux
}
