package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloMux01)
	http.ListenAndServe(":8081", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/teste", HelloMux02)
	http.ListenAndServe(":8082", mux2)
}

func HelloMux01(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello 01"))
}

func HelloMux02(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello 02"))
}
