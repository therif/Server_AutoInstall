package main

import (
	"bufio"
	"fmt"
	"os"

	"autoinstaller/configure"
	"autoinstaller/install"
)

const thefancy = "" +
	"     ████████╗██╗  ██╗███████╗██████╗ ██╗███████╗ \n" +
	"     ╚══██╔══╝██║  ██║██╔════╝██╔══██╗██║██╔════╝ \n" +
	"        ██║   ███████║█████╗  ██████╔╝██║█████╗   \n" +
	"        ██║   ██╔══██║██╔══╝  ██╔══██╗██║██╔══╝   \n" +
	"        ██║   ██║  ██║███████╗██║  ██║██║██║      \n" +
	"        ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝╚═╝      "

func main() {

	defer fmt.Println("\nMain Process End -> Exit!")

	skipexec := false
	if len(os.Args) > 1 {
		if os.Args[1] == "-t" || os.Args[1] == "test" {
			skipexec = true
		} else if os.Args[1] == "-s" || os.Args[1] == "skip" {
			skipexec = true
		} else if os.Args[1] == "-h" || os.Args[1] == "help" {
			fmt.Println("\nUsage :")
			fmt.Println("\n-s or skip : For skip execution")
			fmt.Println("\n-t or test : For skip execution, just emulate script")
			fmt.Println("\n-h or help : For skip execution, just emulate script")
			fmt.Println("")
			fmt.Println("")
			os.Exit(0)
		} else {
			fmt.Println("\nWrong Arguments, use -h for help")
			os.Exit(0)
		}
	}

	fmt.Println("")
	fmt.Println(thefancy)
	fmt.Println("                     by therif")
	fmt.Println("                 github.com/therif")
	fmt.Println("")

	fmt.Println("======== SERVER TOOLS ========")
	fmt.Println("1. Install")
	fmt.Println("2. Configuration")
	fmt.Println("3. Sites Operation")

	fmt.Println("--")
	fmt.Println("0. Exit")
	fmt.Println("")
	fmt.Print("Fill according to choice : ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	jawaban := input.Text()
	if jawaban == "1" {
		install.Start(skipexec)
	} else if jawaban == "2" {
		configure.Start(skipexec)
	} else {
		//
	}

}
