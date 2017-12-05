package imageFileServer

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sasasaiki/gkgimg"
)

//HandlerFuncI ハンドリングすべき全てのfuncを持つ
type HandlerFuncI interface {
	save(w http.ResponseWriter, r *http.Request)
	update(w http.ResponseWriter, r *http.Request)
	delete(w http.ResponseWriter, r *http.Request)
	get(w http.ResponseWriter, r *http.Request)
}

func (h *prodHandlerFunc) save(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, e := r.FormFile("image")
	//本当はハンドラーにもたせた方がいい気がする
	im := gkgimg.DirImgStorage{}
	e = im.SaveResizedImage(file, fileHeader.Filename, "newFile", "testDir", 400, 0, 90)
	if e != nil {
		outputError(&w, e, "fileの保存")
		return
	}
	io.WriteString(w, "保存成功")
}

func (h *prodHandlerFunc) update(w http.ResponseWriter, r *http.Request) {

}

func (h *prodHandlerFunc) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *prodHandlerFunc) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars["name"])
}

func outputError(w *http.ResponseWriter, e error, message string) {
	io.WriteString(*w, e.Error())
	log.Println(message, " エラーが発生しました:", e)
}

type prodHandlerFunc struct {
}
