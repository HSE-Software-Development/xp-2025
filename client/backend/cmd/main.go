package main

import (
	"flag"
	"fmt"

	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/cli"
	"github.com/HSE-Software-Development/xp-2025/client/backend/internal/gui"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "cli", "Режим работы программы (cli/gui)")
	flag.Parse()

	if mode == "cli" {
		fmt.Println("Запуск в CLI режиме...")
		cli.RunCLI()
	} else if mode == "gui" {
		fmt.Println("Запуск в GUI режиме...")
		gui.RunServer()
	} else {
		panic("Неизвестный режим: " + mode)
	}
}
