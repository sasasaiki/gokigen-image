package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

func main() {

	//cssやjsを読み込めるようにするHandler
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	//indexへのルーティング
	http.Handle("/", &templeteHandler{fileName: "main/index.html"})

	//webServerの開始
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
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
