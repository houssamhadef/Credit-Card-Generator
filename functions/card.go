package functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func GenCards(bin string, total int) any {
	start := time.Now()
	generated := make(map[string]bool, total)
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("❌ Failed to get current directory:", err)
	}
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	folderName := filepath.Join(cwd, "logs", "cards_", timestamp)
	err = os.MkdirAll(folderName, 0755)
	if err != nil {
		fmt.Println("❌ Failed to create folder:", err)
	}
	filePath := filepath.Join(folderName, "generated_cards.txt")
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for len(generated) < total {
		card := GenLuhnCard(bin)
		if !generated[card] {
			exp := GenExpirationDate()
			cvv := GenCVV()
			line := card + "|" + exp + "|" + cvv + "\n"
			writer.WriteString(line)
			generated[card] = true
		}
	}
	fmt.Println("✅ Generated " + fmt.Sprint(total) + " in " + time.Since(start).String())
	return writer.Flush()
}

func GenLuhnCard(bin string) string {
	numbers := "0123456789"
	card := bin

	for len(card) < 15 {
		card += string(numbers[rand.Intn(10)])
	}
	checkDigit := calculateLuhnCheckDigit(card)
	card += string('0' + checkDigit)

	return card
}

func GenExpirationDate() string {
	monthsAhead := rand.Intn(60) + 1
	expiry := time.Now().AddDate(0, monthsAhead, 0)
	return expiry.Format("01|06") // MM|YY
}

func GenCVV() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

func calculateLuhnCheckDigit(number string) int {
	sum := 0
	reversed := reverse(number)

	for i, r := range reversed {
		n := int(r - '0')

		if i%2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}
	return (10 - (sum % 10)) % 10
}
