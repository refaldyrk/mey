package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/mey/model"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func Send(filePath string) {
	fmt.Println("Waiting For Receiving Data...")

	url := os.Getenv("CORE_SHAREF_URL")

	if url == "" {
		fmt.Println("Sharef URL: not set")
		return
	}

	urlReceive := url + "/receive"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file to form:", err)
		return
	}

	writer.Close()

	request, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var resp model.ResponseSend
	json.Unmarshal(body, &resp)

	fmt.Println("Code : " + resp.Data.Code)

	chWaiting := make(chan bool)

	go func() {
		for {
			r, err := http.Get(urlReceive + "/" + resp.Data.Code)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer r.Body.Close()

			if r.StatusCode != 200 {
				chWaiting <- true
			}
		}
	}()

	<-chWaiting
	fmt.Println("Success Send: " + filePath)
}

func Receive(code string) {

	url := os.Getenv("CORE_SHAREF_URL")

	if url == "" {
		fmt.Println("Sharef URL: not set")
		return
	}

	resp, err := http.Get(url + "/file/" + code)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	var get model.ResponseGet

	err = json.Unmarshal(body, &get)

	if err != nil {
		fmt.Println(err)
		return
	}

	d, _ := http.Get(get.Data)

	outputFile, err := os.Create(get.Name)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}

	_, err = io.Copy(outputFile, d.Body)
	if err != nil {
		fmt.Println("Error copying content to file:", err)
		return
	}

	fmt.Println("Success Receive: " + code)
}
