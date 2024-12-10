package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stringToInt(strArr []string) []int {
	var intArr []int

	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		intArr = append(intArr, num)
	}

	return intArr
}

func getMiddleNumber(numbers []int) int {
	middleIndex := (len(numbers) - 1) / 2
	return numbers[middleIndex]
}

func contains(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}

func checkUpdateIsValid(update []int, relations map[int][]int) bool {
	for i := 0; i < len(update)-1; i++ {
		num1 := update[i]
		num2 := update[i+1]

		if related, exists := relations[num1]; exists {
			fmt.Printf("The numbers related to %d are %d\n", num1, related)
			if !contains(related, num2) {
				return false
			}
		} else {
			fmt.Printf("No number related to %d found.\n", num1)
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	relations := make(map[int][]int)

	var sum int
	var invalidUpdates [][]int

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
			parts := strings.Split(line, ",")
			partsInt := stringToInt(parts)

			isValidUpdate := checkUpdateIsValid(partsInt, relations)

			if !isValidUpdate {
				invalidUpdates = append(invalidUpdates, partsInt)
			}
		}
	}

	for i := 0; i < len(invalidUpdates); i++ {
		fmt.Printf("Invalid update: %d \n", invalidUpdates[i])
		update := invalidUpdates[i]

		for true {
			isValidUpdate := true

			for i := 0; i < len(update)-1; i++ {
				num1 := update[i]
				num2 := update[i+1]

				if related, exists := relations[num1]; exists {
					//fmt.Printf("The numbers related to %d are %d\n", num1, related)
					if !contains(related, num2) {
						isValidUpdate = false
						// Swap
						update[i] = num2
						update[i+1] = num1
					}
				} else {
					fmt.Printf("No number related to %d found.\n", num1)
					isValidUpdate = false
				}
			}

			if isValidUpdate {
				break
			}
		}
		sum += getMiddleNumber(update)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum: %d ", sum)
}
