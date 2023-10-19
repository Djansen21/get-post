package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fileContent, err := os.ReadFile("eye.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fileContent[:]))

	bodyReader := bytes.NewReader(fileContent)

	requestURL := "http://localhost:8080/concat"
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	//req.Header.Set("Content-type", "application/json")

	client := &http.Client{} // создал http client
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(bodyReader, "\n\n", resp)

}
