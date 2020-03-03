package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!", "go")

	/*
		comandos de terminal

		go help
		go version
		godoc -http=:6060 = habilita documentação offline
		go env = informações do ambiente
		go doc cmd/vet
		go vet  = verifica se há erro no código - ex.: go vet comandos.go
		go build comandos.go = comando de construir os programas
		go run comandos.go = compila e ja executa em um único programa
	*/
}
