package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Filters an entire file for a keyword.
// Options:
// --file (required)
// --keyword (required)
func Filter(filePath string, keyword string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Println(line)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching lines found.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// Displays the last N lines from the specified file.
func GetLastNLines(filePath string, n int) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) > n {
			lines = lines[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return lines, nil
}

// Tails a file and prints lines containing the keyword (non-live version).
func Tail(filePath string, keyword string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if keyword == "" || strings.Contains(line, keyword) {
			fmt.Println(line)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching lines found.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
