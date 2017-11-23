package imageFileServer

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewProdHandler 本番用ハンドラーを作成
func NewProdHandler() (*HandlerFuncI, *Handlers) {
	var hf HandlerFuncI
	hf = new(prodHandlerFunc)
	hs := Handlers{
		index: templeteHandler{fileName: "main/index.html"},
	}
	return &hf, &hs
}

//CreateRoute 渡されたhandlerとfuncについてrouteを設定する
func CreateRoute(hf *HandlerFuncI, h *Handlers) *mux.Router {
	r := mux.NewRouter()

	setAPIRoute(r, hf)
	setRouteExistHandler(r, h)

	// 404のときのハンドラ
	//r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	return r
}

//Handlerが必要ないrouteの設定
func setAPIRoute(r *mux.Router, hp *HandlerFuncI) {
	h := *hp
	setHandler(r, "/save", h.save, "POST")
	setHandler(r, "/get/{name}/", h.get, "GET")
	setHandler(r, "/update", h.update, "PUT")
	setHandler(r, "/delete", h.delete, "DELETE")
}

func setHandler(r *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request), methods ...string) {
	lh := NewLogHandler(http.HandlerFunc(f))
	r.Handle(path, lh).Methods(methods...)
}

//Handlerが必要なrouteの設定
//templete読み込みなど
func setRouteExistHandler(r *mux.Router, h *Handlers) {
	r.Handle("/index", &(h.index))
}
