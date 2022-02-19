package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http" //Pacote para fazer requisicao Web
	"os"       //Pacote para usar o sair
	"strconv"  // converter qualquer coisa para string
	"strings"
	"time" // para usar tempo e seus atributos
)

const delay = 3
const monitoramentoTime = 2

func main() {

	exibeIntroducao()
	fmt.Println("")

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome, idade := devolveNomeEIdade()
	fmt.Println("Olá, sr.", nome, "sua idade é", idade, "anos.")
	fmt.Println("Bem vindo!")

}

func exibeMenu() {
	fmt.Println("O que deseja fazer?")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O endereço da minha variavel comando é", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	// for i := 0; i <= len(sites); i++ {
	for i := 0; i < monitoramentoTime; i++ {
		for i, site := range sites { //FOR em range primeiro posicao e depois a informacao
			fmt.Println("Posição ", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func devolveNomeEIdade() (string, int) {
	nome := "Italo"
	idade := 28
	return nome, idade ///função que retornar 2 informações
	//PARA NAO PRECISAR USAR AS DUAS INFORMAÇÕES DA FUNÇÃO EU POSSO USAR O _ É UM OPERADOR EM BRANCO
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(string(arquivo))

}

//ARRAY AS DESVANTAGENS DE TRABALHAR COM ARRAY PORQUE TEM SETAR O TAMANHO FIXO DELES
// var sites [4]string
// 	sites[0] = "https://random-status-code.herokuapp.com/"
// 	sites[1] = "https://www.alura.com.br"
// 	sites[2] = "https://www.caelum.com.br"

//CRIAR UM SLICE

// func exibeNomes() {
// 	nomes := []string{
// 		"Douglas",
// 		"Daniel",
// 		"Bernado",
// 	}
// 	fmt.Println(nomes)
// }
