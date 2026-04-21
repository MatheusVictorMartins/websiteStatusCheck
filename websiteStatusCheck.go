package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
		imprimeLogs()
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
	// sites := []string{"https://www.alura.com.br", "https://www.youtube.com", "https://www.github.com"}
	sites := leSitesArquivo()
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
	requisicao, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if requisicao.StatusCode == 200 {
		fmt.Println(index, "-", "Site:", site, "Foi carregado com sucesso!\nCódigo:", requisicao.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println(index, "-", "Site:", site, "Possuí problemas.\nCódigo:", requisicao.StatusCode)
		registraLog(site, false)
	}
}

func leSitesArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	leitor := bufio.NewReader(arquivo)

	for {
		texto, err := leitor.ReadString('\n')
		texto = strings.TrimSpace(texto)
		sites = append(sites, texto)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + " - Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	_, err := os.Stat("log.txt")

	if !os.IsNotExist(err) {
		var dados []string
		arquivo, err := os.Open("log.txt")
		if err != nil {
			fmt.Println("Ocorreu um erro:", err)
		}

		leitor := bufio.NewReader(arquivo)

		for {
			texto, err := leitor.ReadString('\n')
			texto = strings.TrimSpace(texto)
			dados = append(dados, texto)
			if err == io.EOF {
				break
			}
		}
		fmt.Println("========================== LOG ==========================")
		for _, dado := range dados {
			fmt.Println(dado)
		}
		fmt.Println("=========================================================")
		arquivo.Close()
	} else {
		fmt.Println("Ops! Parece que você não fez nenhum monitoramento ainda! Realize um antes de verificar o log dos seus monitoramentos!")
	}
}
