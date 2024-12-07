package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getMiddle(numbers []string) int {

	middleIndex := (len(numbers) - 1) / 2
	middleString := numbers[middleIndex]
	middleNumber, error := strconv.Atoi(middleString)

	if error != nil {
		log.Fatalf("Error converting string to number: %s", middleString)
	}
	return middleNumber
}

func contains(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	relations := make(map[int][]int)
	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				log.Fatalf("Error converting numbers on line: %s", line)
			}
			relations[num1] = append(relations[num1], num2)

		} else {

			isValidUpdate := true

			parts := strings.Split(line, ",")
			for i := 0; i < len(parts)-1; i++ {
				num1, err1 := strconv.Atoi(parts[i])
				num2, err2 := strconv.Atoi(parts[i+1])

				if err1 != nil || err2 != nil {
					log.Fatalf("Error converting numbers on line: %s", line)
				}

				if related, exists := relations[num1]; exists {
					fmt.Printf("The numbers related to %d are %d\n", num1, related)
					if !contains(related, num2) {
						isValidUpdate = false
					}
				} else {
					fmt.Printf("No number related to %d found.\n", num1)
					isValidUpdate = false
				}

			}
			fmt.Printf("Update valid: %t \n", isValidUpdate)

			if isValidUpdate && len(parts) > 1 {
				sum += getMiddle(parts)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %d ", sum)
}
