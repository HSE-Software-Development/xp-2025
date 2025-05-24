package main

import (
	"flag"
	"fmt"

	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/cli"
	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/gui"
)

func main() {
	mode := flag.String("mode", "cli", "Режим запуска: 'cli' или 'gui'")
	flag.Parse()

	fmt.Println(*mode)
	if *mode == "cli" {
		fmt.Println("Запуск в CLI режиме...")
		cli.RunCLI()
	} else if *mode == "gui" {
		fmt.Println("Запуск в GUI режиме...")
		gui.RunServer()
	} else {
		fmt.Println("Error")
	}
}
