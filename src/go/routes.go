package imageFileServer

import (
	"github.com/gorilla/mux"
)

//CreateRoute 渡されたhandlerとfuncについてrouteを設定する
func CreateRoute(hf *HandlingFuncI, h *Handlers) *mux.Router {
	r := mux.NewRouter()

	setAPIRoute(r, hf)
	setRouteExistHandler(r, h)

	// 404のときのハンドラ
	//r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	return r
}

//Handlerが必要ないrouteの設定
func setAPIRoute(r *mux.Router, hp *HandlingFuncI) {
	h := *hp
	r.HandleFunc("/save", h.save).Methods("POST")
	r.HandleFunc("/get", h.get).Methods("GET")
	r.HandleFunc("/update", h.update).Methods("PUT")
	r.HandleFunc("/delete", h.delete).Methods("DELETE")
}

//Handlerが必要なrouteの設定
//templete読み込みなど
func setRouteExistHandler(r *mux.Router, h *Handlers) {
	r.Handle("/index", &(h.index))
}
