package main

import (
	"fmt"
	"github.com/refaldyrk/mey/command"
	"os"
)

func main() {
	arg := os.Args[1]
	switch arg {
	case "test":
		command.TestConnection()
		break
	case "show":
		command.ShowUrl()
		break
	case "send":
		command.Send(os.Args[2])
		break
	case "receive":
		command.Receive(os.Args[2])
		break
	case "help":
		fmt.Println("Usage: mey [command] [url] [payload] [options]\n\nAvailable commands:\n  test\n  show\n  send\n  receive\n  help")
	}
}
