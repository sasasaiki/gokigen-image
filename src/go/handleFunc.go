package imageFileServer

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type HandlingFuncI interface {
	save(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
	get(w http.ResponseWriter, r *http.Request)
}

type prodHandlingFunc struct {
}

type Handlers struct {
	index templeteHandler
}

func (h *prodHandlingFunc) save(w http.ResponseWriter, r *http.Request) {

}

func (h *prodHandlingFunc) update(w http.ResponseWriter, r *http.Request) {

}

func (h *prodHandlingFunc) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *prodHandlingFunc) get(w http.ResponseWriter, r *http.Request) {

}

func (t *templeteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//一度だけテンプレートを読み込む
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("views",
				t.fileName)))
	})

	e := t.templ.Execute(w, nil)

	if e != nil {
		fmt.Println("テンプレートの読み込みに失敗しています")
	}
}

type templeteHandler struct {
	once     sync.Once
	fileName string
	templ    *template.Template
}

//NewProdHandler 本番用ハンドラーを作成
func NewProdHandler() (*HandlingFuncI, *Handlers) {
	var hf HandlingFuncI
	hf = new(prodHandlingFunc)
	hs := Handlers{
		index: templeteHandler{fileName: "main/index.html"},
	}
	return &hf, &hs
}
