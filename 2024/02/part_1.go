package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkAdjacent(num1 int, num2 int) bool {
	diff := Abs(num1 - num2)
	if diff < 1 || diff > 3 {
		return false
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		isAdjacent := true

		parts := strings.Fields(line)

		// Check if numbers are adjacent
		for i := 0; i < len(parts)-1; i++ {

			num1, err1 := strconv.Atoi(parts[i])
			num2, err2 := strconv.Atoi(parts[i+1])
			if err1 != nil || err2 != nil {
				log.Fatalf("Error converting numbers on line: %s", line)
			}

			if !checkAdjacent(num1, num2) {
				isAdjacent = false
			}
		}

		// Check if numbers are all increasing or all decreasing
		isIncreasing := true
		isDecreasing := true
		for i := 0; i < len(parts)-1; i++ {
			num1, err1 := strconv.Atoi(parts[i])
			num2, err2 := strconv.Atoi(parts[i+1])
			if err1 != nil || err2 != nil {
				log.Fatalf("Error converting numbers on line: %s", line)
			}

			if num1 > num2 {
				isIncreasing = false
			}
			if num1 < num2 {
				isDecreasing = false
			}
		}

		if isAdjacent && (isIncreasing || isDecreasing) {
			sum++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %d ", sum)
}
