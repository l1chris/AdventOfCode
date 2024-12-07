package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func generatePermutations(operators []string, setSize int, current []string, result *[][]string) {
	// Base case: if the current permutation is the desired size
	if len(current) == setSize {
		// Append a copy of the current permutation to the result
		temp := make([]string, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	// Recursive case: iterate through each operator
	for _, op := range operators {
		// Add the operator to the current permutation
		current = append(current, op)
		// Recurse to fill the next position
		generatePermutations(operators, setSize, current, result)
		// Backtrack: remove the last added operator
		current = current[:len(current)-1]
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalSum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")

		if len(parts) != 2 {
			fmt.Println("Invalid input format")
			return
		}

		num1, err1 := strconv.Atoi(parts[0])
		if err1 != nil {
			log.Fatalf("Error converting numbers on line: %s", line)
		}

		numStrings := strings.Fields(strings.TrimSpace(parts[1]))
		numbers := make([]int, len(numStrings))
		for i, numStr := range numStrings {
			numbers[i], err = strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting number:", err)
				return
			}
		}

		fmt.Println("Number before colon:", num1)
		fmt.Println("Array of numbers:", numbers)

		operators := []string{"+", "*"}
		numberOfCombinations := (len(numbers) - 1)

		var permutations [][]string

		generatePermutations(operators, numberOfCombinations, []string{}, &permutations)

		// Print permutations
		//for _, perm := range permutations {
		//	fmt.Println(perm)
		//}

		for i := 0; i < len(permutations); i++ {
			sum := 0

			for j := 0; j < len(permutations[i]); j++ {

				if j == 0 {
					if permutations[i][j] == "+" {
						sum += numbers[j] + numbers[j+1]
					} else {
						sum += numbers[j] * numbers[j+1]
					}
				} else {
					if permutations[i][j] == "+" {
						sum += numbers[j+1]
					} else {
						sum *= numbers[j+1]
					}
				}
			}

			if sum == num1 {
				fmt.Println("Equation is valid")
				totalSum += sum
				break
			}
		}
	}

	fmt.Printf("Total Sum: %d", totalSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
