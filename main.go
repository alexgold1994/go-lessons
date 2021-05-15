package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed eliftech.jpg
var img []byte

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Eliftech")
}

func handleMeme(w http.ResponseWriter, r *http.Request) {
	w.Write(img)
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/meme", handleMeme)
	http.ListenAndServe(":8080", nil)
}
