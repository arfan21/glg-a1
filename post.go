package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Encode data
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Arfan",
		"email": "arfan@gmail.com",
	})

	responseBody := bytes.NewBuffer(postBody)

	// processing data http post function to make request
	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)

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
