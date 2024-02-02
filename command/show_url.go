package command

import (
	"fmt"
	"os"
)

func ShowUrl() {
	url := os.Getenv("CORE_SHAREF_URL")

	if url == "" {
		fmt.Println("Sharef URL: not set")
		return
	}

	fmt.Println("Sharef URL: " + url)
}
