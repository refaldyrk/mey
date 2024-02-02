package command

import (
	"fmt"
	"net/http"
	"os"
)

func TestConnection() {
	url := os.Getenv("CORE_SHAREF_URL")

	if url == "" {
		fmt.Println("Sharef URL: not set")
		return
	}
	url = url + "/test"

	fmt.Println("Test Connect: " + url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Success!")
	}
}
