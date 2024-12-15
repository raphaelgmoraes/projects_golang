package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showInfo()
	for {
		showMenu()
		command := readCommand()
		switch command {
		case 0:
			fmt.Println("Closing the program...")
			os.Exit(0)
		case 1:
			startMonitor()
		case 2:
			fmt.Println("Show Logging...")
		case 999:
			showHelper()
		default:
			fmt.Println("Error: command not accepted")
			os.Exit(-1)
		}
	}
}

func showHelper() {
	fmt.Println("Info program")
	fmt.Println("0: This command exit program now!")
	fmt.Println("1: This command running program now!")
	fmt.Println("2: This command show log system in real time!")
}

func showInfo() {
	name := "Raphael"
	version_app := 1.1
	commandHelper := 999
	fmt.Println("Hi, Mr/Mrs ", name)
	fmt.Println("The version program is:", version_app)
	fmt.Println("View Documentation press command:", commandHelper)
}

func showMenu() {
	fmt.Println("1- Start Monitor")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The position address in memory is:", &command)
	fmt.Println("The value input captured is: ", command)
	return command
}

func startMonitor() {
	fmt.Println("Running monitor website..")
	site := "https://httpbin.org"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("[Server UP] :: Site: ", site)
	} else {
		fmt.Println("[Server DOWN] :: Site: ", site)
	}
}
