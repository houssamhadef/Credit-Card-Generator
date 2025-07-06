package main

import (
	"bufio"
	"cc-gen.com/m/v2/functions"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	boldWhite := color.New(color.FgWhite, color.Bold).SprintFunc()

	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()

		color.Magenta("╔══════════════════════════════════════════════╗")
		color.Magenta("║ ")
		fmt.Print(boldWhite("  Welcome to the Credit Card Generator  "))
		color.Magenta("║")
		color.Magenta("╚══════════════════════════════════════════════╝")

		fmt.Println()
		fmt.Println(cyan("1.") + " Generate Card from BIN")
		fmt.Println(cyan("2.") + " Generate Random BIN")
		fmt.Println(cyan("0.") + " Exit")
		fmt.Print(green("Choose an option: "))

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			color.Yellow("Enter Bin : ")
			bin, _ := reader.ReadString('\n')
			bin = strings.TrimSpace(bin)
			fmt.Print(green("How many cards you want to generate: "))
			totalStr, _ := reader.ReadString('\n')
			totalStr = strings.TrimSpace(totalStr)
			var total int
			fmt.Sscanf(totalStr, "%d", &total)
			functions.GenCards(bin, total)
		case "2":
			functions.GenerateBinMenu(reader)
		case "0":
			color.HiGreen("Goodbye!")
			return
		default:
			color.Red("Invalid input. Please choose a valid option.")
		}

		fmt.Println("\nPress ENTER to continue...")
		reader.ReadString('\n')
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
