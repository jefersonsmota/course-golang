package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const MONITORINGS = 3
const DELAY_MONITORING = 5

func main() {
	for {
		showMenu()

		processCommand(readCommand())
	}
}

func processCommand(opt int) {
	switch opt {
	case 1:
		startMonitor()
	case 2:
		clearConsole(runtime.GOOS)
		showListLogs()
	case 0:
		clearConsole(runtime.GOOS)
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		clearConsole(runtime.GOOS)
		fmt.Println("Opção invalida", opt)
		os.Exit(-1)
	}
}

func startMonitor() {
	clearConsole(runtime.GOOS)
	fmt.Println("Iniciando monitoramento...")

	sites := getSites()
	logFile := createFileLog()
	defer logFile.Close()

	for i := 0; i < MONITORINGS; i++ {
		for _, site := range sites {
			logSite := verifyPage(site)
			fmt.Println(logSite)
			logFile.WriteString(logSite)
		}
		time.Sleep(DELAY_MONITORING * time.Second)
		logFile.WriteString("\n")
	}
	fmt.Println("Log salvo!")
}

func verifyPage(pageUrl string) string {

	http.Header.Add(http.Header{}, "Cache-Control", "no-store,no-cache,max-age=0,must-revalidate,proxy-revalidate")
	resp, _ := http.Get(pageUrl)

	if resp.StatusCode != 200 {
		return fmt.Sprintf("[%s][FAIL][%d]\n", pageUrl, resp.StatusCode)
	}

	return fmt.Sprintf("[%s][OK]\n", pageUrl)
}

func showListLogs() {
	listaLogs := getListLogs()
	lenLista := len(listaLogs)

	skip := 0
	take := 5

	for {

		paginated := paginatedList(listaLogs, skip, take)

		fmt.Println("===================================")
		for i, item := range paginated {
			fmt.Printf("[%d] %s \n", i+1, item)
		}
		fmt.Println("Digite 9 para continuar ou 0 pra sair...")

		comm := readCommand()

		if comm == 0 {
			break
		}

		if comm <= len(paginated) {
			showLog(paginated[comm-1])
		}

		fmt.Println(skip)

		if comm == 9 && (skip+len(paginated)) < lenLista {
			skip += take
		} else if comm == 9 {
			skip = 0
		}
	}
}

func paginatedList(list []string, skip int, take int) []string {
	if skip > len(list) {
		skip = len(list)
	}

	endList := skip + take
	if endList > len(list) {
		endList = len(list)
	}

	return list[skip:endList]
}

func showLog(logPath string) {
	fmt.Println("\n================", logPath)
	file, err := os.Open("logs\\" + logPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func showMenu() {
	fmt.Println("===================================")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
	fmt.Println("===================================")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return int(command)
}

func processaErro(err error) {
	log.Fatal(err)
}

func getSites() []string {
	file, err := os.Open("sites.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sites := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sites = append(sites, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sites
}

func createFileLog() *os.File {
	file, err := os.Create("logs/log-" + time.Now().Format("02-01-2006-15-04-05") + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func clearConsole(osterminal string) {
	if osterminal == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osterminal == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func getListLogs() []string {
	dir, err := os.ReadDir("logs/")
	if err != nil {
		log.Fatal(err)
	}

	files := []string{}

	for _, de := range dir {
		files = append(files, de.Name())
	}

	return files
}
