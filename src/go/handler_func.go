package imageFileServer

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//HandlerFuncI ハンドリングすべき全てのfuncを持つ
type HandlerFuncI interface {
	save(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
	get(w http.ResponseWriter, r *http.Request)
}

func (h *prodHandlerFunc) save(w http.ResponseWriter, r *http.Request) {
	log.Println("saveにアクセス")
}

func (h *prodHandlerFunc) update(w http.ResponseWriter, r *http.Request) {

}

func (h *prodHandlerFunc) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *prodHandlerFunc) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["name"])
}

type prodHandlerFunc struct {
}
