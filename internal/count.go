package internal

import (
	"bufio"
	"fmt"
	"os"
)

func Count(domain, fileName string) int {
	// Membuka file
	filepath := fmt.Sprintf("output/%s/%s.txt", domain, fileName)
	file, err := os.Open(filepath)
	if err != nil {
		return 0
	}
	defer file.Close()

	r := bufio.NewScanner(file)
	r.Split(bufio.ScanLines)

	var count int
	for r.Scan() {
		count++
	}
	if err := r.Err(); err != nil {
		return 0
	}
	return count

}
