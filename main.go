package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	fileName := flag.String("i", "", "fileName to be shrunk")
	percentage := flag.Float64("p", 0.5, "percentage of the fileName to keep")
	skip := flag.Int("s", 0, "lines to skip")
	flag.Parse()

	if *fileName == "" {
		log.Fatalf("No input fileName was given")
	} else {
		log.Printf("Provided file: %v", *fileName)
	}

	if *percentage < 0 || *percentage > 1.0 {
		log.Fatalf("Invalid percentage: expecting (0, 1.0), got %v", *percentage)
	}

	if *skip < 0 {
		log.Fatalf("Invalid number for lines to skip: expecting > 0, got %v", *skip)
	}

	shrink(*fileName, *percentage, *skip)
}

func shrink(fileName string, percentage float64, skip int) {

	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file: %v, got error: %v", fileName, err)
	}
	defer sourceFile.Close()

	targetFileName := fileName + "." + strconv.FormatFloat(percentage, 'f', 2, 64)
	targetFile, err := os.Create(targetFileName)
	if err != nil {
		log.Fatalf("Failed to create target file: %v, got error: %v", targetFileName, err)
	}
	defer targetFile.Close()

	scanner := bufio.NewScanner(sourceFile)
	i := 0
	for scanner.Scan() {
		if i < skip {
			_, _ = targetFile.WriteString(scanner.Text() + "\n")
			i++
		} else {
			if randomize(percentage) {
				_, _ = targetFile.WriteString(scanner.Text() + "\n")
			}
		}
	}
}
