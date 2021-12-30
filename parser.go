package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

// interesting text we want to save seems to always start with "article class="Box-row"

func main() {
	currentTime := time.Now()
	todaysDate := currentTime.Format("2006-01-02")
	// open the file
	f, err := os.Open("results.txt")
	if err != nil {
		log.Fatal(err)
	}
	of, oferr := os.Create(todaysDate)
	if oferr != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	defer of.Close()

	scanner := bufio.NewScanner(f)
	lineNumber := 0
	//interestingLineIDs := []int{}
	interestingLineFound := false
	chunkCounter := 0
	for scanner.Scan() {

		// do something with a line
		// program finds line of interest and logs it ...
		// need to log the next ten lines after each line of interest and then skip

		if strings.Contains(scanner.Text(), "article class=\"Box-row") {
			//fmt.Printf("line:%d, %s\n", lineNumber, scanner.Text())
			//interestingLineIDs = append(interestingLineIDs, lineNumber)
			interestingLineFound = true
		}
		if interestingLineFound && chunkCounter < 10 {
			if strings.Contains(scanner.Text(), "a href=") &&
				!strings.Contains(scanner.Text(), "a href=\"/login?") &&
				!strings.Contains(scanner.Text(), "network/members") &&
				!strings.Contains(scanner.Text(), "docs.github.com/en/") &&
				!strings.Contains(scanner.Text(), "g-emoji") &&
				!strings.Contains(scanner.Text(), "span") &&
				!strings.Contains(scanner.Text(), "data-hydro-click") {
				outputLine := strings.Fields(scanner.Text())
				outputLine[1] = strings.ReplaceAll(outputLine[1], "href=\"", "www.github.com")
				outputLine[1] = strings.ReplaceAll(outputLine[1], "stargazers\"", "")
				//fmt.Printf("line:%d, %s\n", lineNumber, scanner.Text())
				of.WriteString(outputLine[1])
				of.WriteString("\n")
				chunkCounter++
			}
		}
		if chunkCounter >= 10 {
			interestingLineFound = false
			chunkCounter = 0
		}
		lineNumber++
	}
	//fmt.Printf("interesting line IDs: %v", interestingLineIDs)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
