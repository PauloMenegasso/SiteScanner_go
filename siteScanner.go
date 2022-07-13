package main

import (
	"fmt"
	"net/http"
	"os"
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
			fmt.Println("Showing the logs:")
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
	resp, _ := http.Get(site)
	fmt.Println("Site:", site, " . Status Code:", resp.StatusCode)

	if resp.StatusCode != 200 {
		fmt.Println("Atenção, site", site, "está com problemas!")
	}
}

func readFile() []string {
	var sites []string

	file, _ := os.ReadFile("sites.txt")
	fmt.Println(file)
	return sites
}
