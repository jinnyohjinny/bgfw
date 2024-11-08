package internal

import (
	"bufio"
	"fmt"
	"os"
)

func ReadDomains(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka file: %v", err)
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	// Mengecek error saat membaca file
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("gagal membaca file: %v", err)
	}

	return domains, nil
}
