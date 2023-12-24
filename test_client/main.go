package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type request struct {
	FROM     string `json:"from"`
	TO       string `json:"to"`
	ClientID int    `json:"client_id"`
}

func createOffer() {
	var req = &request{
		FROM:     "Sochi",
		TO:       "sdjf",
		ClientID: 1,
	}

	bytesRepresentation, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post("http://127.0.0.1:8080/createOffer", "application/json", bytes.NewBuffer(bytesRepresentation))
	defer response.Body.Close()

	if err != nil {
		fmt.Println(err)
	}
	bytesResp, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Ответ сервера:", string(bytesResp))
}

func main() {
	createOffer()
}
