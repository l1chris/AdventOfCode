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

	do := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		pattern := `do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`
		regex := regexp.MustCompile(pattern)
		matches := regex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			fmt.Println("Valid match:", match)
			if match[0] == "do()" {
				do = true
			} else if match[0] == "don't()" {
				do = false
			} else {
				if do {
					num1, err1 := strconv.Atoi(match[1])
					num2, err2 := strconv.Atoi(match[2])
					if err1 != nil || err2 != nil {
						log.Fatalf("Error converting numbers on line: %s", line)
					}

					sum = num1 * num2
					totalSum += sum
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %d ", totalSum)
}
