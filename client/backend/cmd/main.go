package main

import "flag"

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "cli", "Режим работы программы (cli/gui)")
	flag.Parse()
}
