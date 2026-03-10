package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ----------- PARSER ----------------

func parseBanner(banner []string) map[rune][]string {
	charMap := make(map[rune][]string)
	var block []string
	code := 32 //ASCII space

	for _, line := range banner {
		line = strings.TrimRight(line, "\r")
		if line == "" {
			if len(block) > 0 {
				charMap[rune(code)] = block
				block = []string{}
				code++
			}
		} else {
			block = append(block, line)
		}
	}
	if len(block) > 0 {
		charMap[rune(code)] = block
	}
	return charMap
}

func printBannertoArt(text string, charMap map[rune][]string) string {
	lines := strings.Split(text, "\n")
	var result []string

	for _, line := range lines {

		if line == "" {
			result = append(result, "")
			continue
		}

		maxHeight := 0
		for _, char := range line {
			if block, ok := charMap[char]; ok {
				if len(block) > maxHeight {
					maxHeight = len(block)
				}
			}
		}

		for row := 0; row < maxHeight; row++ {

			var builder strings.Builder

			for _, char := range line {
				block, ok := charMap[char]

				if ok {
					if row < len(block) {
						builder.WriteString(block[row])
					} else {
						builder.WriteString(strings.Repeat(" ", len(block[0])))
					}
				}
			}

			result = append(result, builder.String())
		}
	}

	return strings.Join(result, "\n")
}

func main() {

	if len(os.Args) != 4 || !(len(os.Args[1]) > 9) {
		log.Fatal("\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

	//------------ OPTIONS ----------------
	options := os.Args[1]

	outputFileName := ""

	if options[:9] == "--output=" && strings.HasSuffix(options, ".txt") {
		outputFileName = strings.TrimPrefix(options, "--output=")
	} else {
		log.Fatal("\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

	//---------- INPUT TEXT ----------------
	inputText := strings.ReplaceAll(os.Args[2], "\\n", "\n")

	//--------- BANNER FILE-----------------
	bannerFile := os.Args[3] + ".txt"

	buffer, err := os.ReadFile(bannerFile)

	if err != nil {
		log.Fatal("Error reading Banner file: ", err)
	}
	banner := strings.Split(string(buffer), "\n")

	finalProcessedResult := printBannertoArt(inputText, parseBanner(banner))

	outputFile, err := os.Create(outputFileName)

	if err != nil {
		log.Fatal("Error creating output file: ", err)
	}

	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	_, err = writer.WriteString(finalProcessedResult)

	if err != nil {
		log.Fatalf("Error Writing to output file: %v\n", err)
	}

	writer.Flush()

}
