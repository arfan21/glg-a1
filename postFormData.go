package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	formData := url.Values{
		"name":  {"tes"},
		"email": {"email@gmail.com"},
	}

	// processing data http post function to make request"
	resp, err := http.PostForm("https://postman-echo.com/post", formData)

	// handle err
	if err != nil {
		log.Fatalf("An error occured %v", err)
	}

	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
