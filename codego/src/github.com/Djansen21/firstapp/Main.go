package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func squaring(writer http.ResponseWriter, req *http.Request) {
	val := req.PostFormValue("Number")
	log.Println(val)
	digit, err := strconv.Atoi(val)
	log.Println(digit)
	if err != nil {
		io.WriteString(writer, "Value not supported")
		log.Println(err)

	}
	digit = digit * digit
	log.Println(digit)
	digitStr := strconv.Itoa(digit)
	writer.Write([]byte(digitStr))
}

func main() {
	port := ":8080"
	http.HandleFunc("/square", squaring)
	log.Println("Starting server at " + port)
	http.ListenAndServe(port, nil)
}

//Написать клиент, который отправляет запрос на этот сервер//
//чтоб она принимала данные в json
// клиент --> json
// в json 2 поля берет - контатинирует.
// сделать мэпу

// Сделать на клиенте: мэпу
// заинкодить в json
// отдать серверу - клиент парсит все поля и закотатинирует и отдает строчку, где сложены строчки

//
