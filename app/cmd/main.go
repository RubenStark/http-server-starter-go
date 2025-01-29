package main

import (
	"flag"
	"fmt"

	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func main() {

	// Definir flag para el directorio base
	dir := flag.String("directory", "./", "Directorio base para servir archivos")
	flag.Parse()

	fmt.Println("Iniciando servidor con base en:", *dir)

	server.StartServer(*dir)
}
