package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 5
const monitoramentoDelay = 5

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
	sites := []string{"https://www.alura.com.br", "https://www.youtube.com", "https://www.github.com"}
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			testaSite(i, site)
		}
		time.Sleep(monitoramentoDelay * time.Second)
		fmt.Println("\nTestando... Por favor não desligue seu computador")
	}
	fmt.Println("\nTeste completo! voltando para o menu inicial...")
}

func testaSite(index int, site string) {
	requisicao, _ := http.Get(site)
	if requisicao.StatusCode == 200 {
		fmt.Println(index, "-", "Site:", site, "Foi carregado com sucesso!\nCódigo:", requisicao.StatusCode)
	} else {
		fmt.Println(index, "-", "Site:", site, "Possuí problemas.\nCódigo:", requisicao.StatusCode)
	}
}
