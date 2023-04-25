package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "claudio"}`))
	res, err := c.Post("http://uol.com.br", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
