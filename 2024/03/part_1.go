package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum int
	var totalSum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		pattern := `mul\((\d+),(\d+)\)`

		re := regexp.MustCompile(pattern)

		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			// match[1] is x, match[2] is y
			num1, err1 := strconv.Atoi(match[1])
			num2, err2 := strconv.Atoi(match[2])
			if err1 != nil || err2 != nil {
				log.Fatalf("Error converting numbers on line: %s", line)
			}

			sum = num1 * num2
			totalSum += sum
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %d ", totalSum)
}
