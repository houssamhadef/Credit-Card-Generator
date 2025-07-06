package functions

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func GenerateBinMenu(reader *bufio.Reader) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	//yellow := color.New(color.FgYellow).SprintFunc()

	for {
		fmt.Println()
		color.Magenta("╔════════════════════════════════╗")
		color.Magenta("║        Select Card Type        ║")
		color.Magenta("╚════════════════════════════════╝")

		fmt.Println(cyan("1.") + " Visa")
		fmt.Println(cyan("2.") + " Mastercard")
		fmt.Println(cyan("3.") + " American Express")
		fmt.Println(cyan("0.") + " Return to main menu")
		fmt.Print("Your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		color.Yellow("Enter total bin : ")
		totalStr, _ := reader.ReadString('\n')
		totalStr = strings.TrimSpace(totalStr)
		var total int
		fmt.Sscanf(totalStr, "%d", &total)
		switch choice {
		case "1":
			color.Green("You selected Visa.")
			GenBins(total, "VISA")

		case "2":
			color.Green("You selected Mastercard.")
			GenBins(total, "MC")

		case "3":
			color.Green("You selected American Express.")
			GenBins(total, "AMEX")
		case "0":
			color.HiBlue("Returning to main menu...")
			return ""
		default:
			color.Red("Invalid choice. Please select 1, 2, or 3.")
		}
	}
}
