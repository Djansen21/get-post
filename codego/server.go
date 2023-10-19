package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func concat(writer http.ResponseWriter, req *http.Request) {

	body, err := io.ReadAll(req.Body) // почему без io не вычитывает бади? почему просто нельзя вычитать бади, без io
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(body) === выводит в байтах
	//fmt.Println(string(body))
	content := string(body)

	// bodyReader := bytes.NewReader(fileContent)
	// writer.Write([]byte(digitStr))
	type Row struct {
		eye string
	}
	var row Row
	json.Unmarshal([]byte(content), &row)
	fmt.Println(content)

}

func main() {
	port := ":8080"
	http.HandleFunc("/concat", concat)
	log.Println("Starting server at " + port)
	http.ListenAndServe(port, nil)

}
