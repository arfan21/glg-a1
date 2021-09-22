package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListLanguages struct {
	List []Language `json:"languages"`
}

func (ll *ListLanguages) AddLanguage(data Language) {
	ll.List = append(ll.List, data)
}

func (ll *ListLanguages) GetLanguage(name string) *Language {
	for _, val := range ll.List {
		if val.Language == name {
			return &val
		}
	}

	return nil
}

func (ll *ListLanguages) GetAllLanguage() []Language {

	return ll.List
}

type Language struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}

func (L Language) Print() {
	fmt.Println(L)
}

type Relation struct {
	InfluencedBy []string `json:"influenced-by,"`
	Influences   []string `json:"influences"`
}

var DataLanguange ListLanguages

type Request struct {
	Param string `json:"param"`
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/language", HandlerLanguage)
	http.HandleFunc("/palindrom", HandlerPalindrom)

	fmt.Println("Http Server listening on http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}

func HandlerLanguage(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		newLanguage := new(Language)
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		err := decoder.Decode(newLanguage)
		if err != nil {
			WriteResponse(rw, 400, Response{400, err})
			return
		}

		DataLanguange.AddLanguage(*newLanguage)
		newLanguage.Print()
		WriteResponse(rw, 200, Response{200, newLanguage})
		return
	}

	if r.Method == "GET" {
		name := r.URL.Query().Get("name")

		if name == "" {
			WriteResponse(rw, http.StatusOK, Response{http.StatusOK, DataLanguange.GetAllLanguage()})
			return
		}

		data := DataLanguange.GetLanguage(name)

		if data == nil {
			WriteResponse(rw, http.StatusNotFound, Response{http.StatusNotFound, "Not Found"})
			return
		}

		WriteResponse(rw, http.StatusOK, Response{http.StatusOK, data})
		return
	}

	WriteResponse(rw, http.StatusMethodNotAllowed, Response{http.StatusMethodNotAllowed, "method not allowed"})
}

func HandlerPalindrom(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		WriteResponse(rw, http.StatusMethodNotAllowed, Response{http.StatusMethodNotAllowed, "method not allowed"})
		return
	}
	text := r.URL.Query().Get("text")

	if text == "" {
		WriteResponse(rw, http.StatusBadRequest, Response{http.StatusBadRequest, nil})
		return
	}

	result := isPalindrom(text)

	WriteResponse(rw, 200, Response{Status: 200, Data: result})
}

func isPalindrom(text string) bool {
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}
	return true
}

func WriteResponse(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(data)
}
