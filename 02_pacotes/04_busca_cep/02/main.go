package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCEP)
	http.ListenAndServe(":8080", nil)
}

func BuscaCEP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World"))
}
