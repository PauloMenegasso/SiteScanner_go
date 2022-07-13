package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringTimes = 3
const monitoringDelay = 5

func main() {
	//sites := []string{"https://autenticacao.cappta.com.br/", "https://www.linkedin.com/feed/", "https://www.cappta.com.br/"}

	sites := readFile()

	showIntro()
	for {
		showMenu()
		command := selectCommand()

		switch command {
		case 1:
			startMonitoring(sites)
		case 2:
			printLogs()
		case 0:
			fmt.Println("Closing the program.")
			fmt.Println("Thanks for using.")
			os.Exit(0)
		default:
			fmt.Println("Error: Unknown command")
		}
	}
}

func showIntro() {
	version := 1.0
	fmt.Println("This program runs at version", version)

	fmt.Println("Please insert your name")
	fmt.Println("")

	var name string
	fmt.Scan(&name)

	fmt.Println("Hello", name)
	fmt.Println("")
}

func showMenu() {
	fmt.Println("Select an option:")
	fmt.Println("1 - Monitor sites")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit program")
	fmt.Println("")
}

func selectCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)
	return command
}

func startMonitoring(sites []string) {
	fmt.Println("Monitoring...")

	for j := 0; j < monitoringTimes; j++ {
		for i := 0; i < len(sites); i++ {
			testSite(sites[i])
		}

		time.Sleep(monitoringDelay * time.Second)
		fmt.Println("")
	}
}

func testSite(site string) {
	resp, err := http.Get(site)
	logError(err)

	statusCode := resp.StatusCode

	fmt.Println("Site:", site, " . Status Code:", statusCode)

	createLogFile(site, statusCode)

	if statusCode != 200 {
		fmt.Println("Atenção, site", site, "está com problemas!")
	}
}

func readFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")
	logError(err)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		logError(err)

		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func createLogFile(site string, statusCode int) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logError(err)

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - Site: " + site + ". Resposta: " + strconv.Itoa(statusCode) + "\n")

	file.Close()
}

func logError(err error) {
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
}

func printLogs() {
	fmt.Println("Showing the logs:")

	arquivo, err := ioutil.ReadFile("logs.txt")
	logError(err)

	fmt.Println(string(arquivo))
}
