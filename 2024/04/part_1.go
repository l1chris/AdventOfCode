package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][2]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Down-right diagonal
	{-1, -1}, // Up-left diagonal
	{-1, 1},  // Up-right diagonal
	{1, -1},  // Down-left diagonal
}

func searchFromPosition(grid [][]byte, word string, row, col int, dir [2]int) bool {
	wordLen := len(word)
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < wordLen; i++ {
		// Calculate next position
		newRow := row + i*dir[0]
		newCol := col + i*dir[1]

		// Check bounds
		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
			return false
		}

		// Check character match
		if grid[newRow][newCol] != word[i] {
			return false
		}
	}
	return true
}

func findSequence(grid [][]byte, word string) [][2]int {
	var result [][2]int
	rows, cols := len(grid), len(grid[0])

	// Iterate through every cell in the grid
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Check all 8 directions from the current position
			for _, dir := range directions {
				if searchFromPosition(grid, word, row, col, dir) {
					result = append(result, [2]int{row, col})
				}
			}
		}
	}
	return result
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var array2D [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Convert the line into a slice of bytes (each character)
		row := []byte(line)

		array2D = append(array2D, row)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the 2D array
	//for i, row := range array2D {
	//	fmt.Printf("Row %d: %q\n", i, row)
	//}

	word := "XMAS"

	// Find the sequence
	occurrences := findSequence(array2D, word)

	fmt.Printf("Found %d sequences", len(occurrences))
}
