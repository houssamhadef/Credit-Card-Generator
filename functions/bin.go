package functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const numbers = "0123456789"

func GenBins(total int, CardType string) any {
	prefix := map[string]string{
		"VISA": "4",
		"MC":   "5",
		"AMEX": "6",
	}
	start := time.Now()
	generated := make(map[string]bool, total)
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("❌ Failed to get current directory:", err)
	}
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	folderName := filepath.Join(cwd, "logs", "bins_", timestamp)
	err = os.MkdirAll(folderName, 0755)
	if err != nil {
		fmt.Println("❌ Failed to create folder:", err)
	}
	filePath := filepath.Join(folderName, "generated_bins.txt")
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for len(generated) < total {
		bin := GenBin()

		if !generated[bin] {
			line := prefix[CardType] + bin + "\n"
			writer.WriteString(line)
			generated[bin] = true

		}
	}
	fmt.Println("✅ Generated " + fmt.Sprint(total) + " in " + time.Since(start).String())
	return writer.Flush()
}

func GenBin() string {
	result := ""
	for i := 0; i < 5; i++ {
		result += string(numbers[rand.Intn(10)])
	}
	return result
}
