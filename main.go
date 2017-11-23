package main

import (
	"fmt"
	"net/http"

	"github.com/sasasaiki/image-file-server/src/go"
)

func main() {
	r := imageFileServer.CreateRoute(imageFileServer.NewProdHandler())
	http.ListenAndServe(":8080", r)
	fmt.Println("owatta")
}
