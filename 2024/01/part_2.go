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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var array1, array2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into parts
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("Invalid line format: %s", line)
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Error converting numbers on line: %s", line)
		}

		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	countMap := make(map[int]int) // Map to count occurrences

	// Count occurrences in array2
	for _, num := range array2 {
		countMap[num]++
	}

	var sum int

	for _, num := range array1 {
		sum += num * countMap[num]
	}

	fmt.Printf("Sum: %d", sum)
}
