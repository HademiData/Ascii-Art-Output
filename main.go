package main

import (
	"log"
	"os"
	"strings"
)

// -------- OPTION PARSER --------

func parseOption(arg string) string {

	if !strings.HasPrefix(arg, "--output=") || !strings.HasSuffix(arg, ".txt") {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

	return strings.TrimPrefix(arg, "--output=")
}

// -------- BANNER PARSER --------

func parseBanner(lines []string) map[rune][]string {

	charMap := make(map[rune][]string)

	for i := 0; i < 95; i++ {

		start := i*9 + 1
		end := start + 8

		charMap[rune(i+32)] = lines[start:end]

	}

	return charMap
}

// -------- ASCII RENDERER --------

func render(text string, banner map[rune][]string) string {

	lines := strings.Split(text, "\n")
	var result []string

	for _, line := range lines {

		if line == "" {
			result = append(result, "")
			continue
		}

		for row := 0; row < 8; row++ {

			var out string

			for _, char := range line {

				if block, ok := banner[char]; ok {
					out += block[row]
				}

			}

			result = append(result, out)

		}

	}

	return strings.Join(result, "\n")
}

// -------- MAIN --------

func main() {

	if len(os.Args) != 4 {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

	// parse flag
	outputFile := parseOption(os.Args[1])

	// handle \n
	input := strings.ReplaceAll(os.Args[2], "\\n", "\n")

	// banner file
	bannerFile := os.Args[3] + ".txt"

	data, err := os.ReadFile(bannerFile)
	if err != nil {
		log.Fatal("Error reading banner file:", err)
	}

	lines := strings.Split(string(data), "\n")

	charMap := parseBanner(lines)

	result := render(input, charMap)

	// write output
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		log.Fatal("Error writing output file:", err)
	}
}
