package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mostraIntroducao()
	for {
		mostraMenu()
		comando := leComando()
		executaComando(comando)
	}
}

func mostraIntroducao() {
	nome := "Matheus"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func mostraMenu() {
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair")
}

func executaComando(valor int) {
	switch valor {
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	default:
		fmt.Println("Não reconheço esse comando.")
		os.Exit(-1)
	}
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")

	requisicao, _ := http.Get(a.sites[1])
	if requisicao.StatusCode == 200 {
		fmt.Println("Site:", a.sites[1], "Foi carregado com sucesso!\nCódigo:", requisicao.StatusCode)
	} else {
		fmt.Println("Site:", a.sites[1], "Possuí problemas.\nCódigo:", requisicao.StatusCode)
	}
}
