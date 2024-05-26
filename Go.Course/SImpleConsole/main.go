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

const secondsDelayBetweenCAll = 2

func main() {
	exibeIntroducao()

	for {
		exibeItens()

		comando := lerComando()

		exibeMenu(comando)
	}

}

func exibeIntroducao() {
	nome := "Gui"
	versao := 1.1

	println("versao:", versao)
	println("Ola,", nome)
}

func exibeItens() {
	println("1 - Iniciar monitoramento")
	println("2 - Exibir os logs")
	println("0 - Sair do programa")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func exibeMenu(comando int) {
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		println("Exibindo logs...")
		readLogs()
	case 0:
		println("Saindo do programa...")
		os.Exit(0)

	default:
		println("COmando nao conhecido")
		os.Exit(-1)
	}

}

func iniciarMonitoramento() {
	println("Monitorando...")

	sites := readFile()

	println("quantas vezes deseja que o teste seja executada?")
	times := 0
	fmt.Scan(&times)

	for i := 0; i < times; i++ {

		for _, site := range sites {
			response, _ := http.Get(site)

			if response.StatusCode == http.StatusOK {
				writeLog(site, true)
				println("tudo certo com o site", site)
			} else {
				println("houve um erro ao monitorar o site", site, "status:", response.StatusCode)
				writeLog(site, false)
			}
		}

		time.Sleep(secondsDelayBetweenCAll * time.Second)

	}

}

func readFile() []string {
	file, error := os.Open("sites.txt")
	slice := []string{}

	if error != nil {
		println("Ocorreu um erro:", error)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		slice = append(slice, strings.TrimSpace(line))
	}

	file.Close()

	return slice
}

func writeLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		println("ERRO " + err.Error())
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + "online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func readLogs() {
	arquivo, err := os.ReadFile("log.txt")

	if err != nil {
		println(err)
	}

	println(string(arquivo))
}
