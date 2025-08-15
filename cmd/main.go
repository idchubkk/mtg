package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: my-tg-proxy <command> [args]")
		fmt.Println("Commands: run")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		run()
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
