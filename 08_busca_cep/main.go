package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := req.URL.Query().Get("cep")

	if cepParam == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, error := BuscaCep(cepParam)

	if error != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(cep)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	res, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

	if error != nil {
		return nil, error
	}

	defer res.Body.Close()

	body, error := ioutil.ReadAll(res.Body)

	if error != nil {
		return nil, error
	}

	var c ViaCEP

	error = json.Unmarshal(body, &c)

	if error != nil {
		return nil, error
	}

	return &c, nil
}
