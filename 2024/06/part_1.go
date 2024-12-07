package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = map[string][2]int{
	"^": {-1, 0}, // Up
	">": {0, 1},  // Right
	"v": {1, 0},  // Down
	"<": {0, -1}, // Left
}

var transitions = map[string]string{
	"^": ">",
	">": "v",
	"v": "<",
	"<": "^",
}

func nextPostitionHasObstacle(grid [][]string, currentRow int, currentCol int, dir [2]int) bool {
	newRow := currentRow + dir[0]
	newCol := currentCol + dir[1]

	// checkObstacle
	if grid[newRow][newCol] == "#" {
		fmt.Printf("Obstacle found at row %d and column %d\n", newRow, newCol)
		return true
	}
	return false
}

func findPosition(grid [][]string, direction string) [][2]int {
	var position [][2]int
	rows, cols := len(grid), len(grid[0])

	// Iterate through every cell in the grid
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == direction {
				position = append(position, [2]int{row, col})
				return position
			}
		}
	}
	return position
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var array2D [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Convert the line into a slice of bytes (each character)
		row := []byte(line)

		// Convert to a string slice
		var rowStrings []string
		for _, b := range row {
			rowStrings = append(rowStrings, string(b))
		}

		array2D = append(array2D, rowStrings)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Determine bounds
	numberOfRows := len(array2D) - 1
	numberOfColumns := len(array2D[0]) - 1
	fmt.Printf("numberOfRows %d\n", numberOfRows)
	fmt.Printf("numberOfColumns %d\n", numberOfColumns)

	currentDirection := "^"

	var steps int

	// Find and mark the current position
	currentPosition := findPosition(array2D, currentDirection)
	curRow := currentPosition[0][0]
	curCol := currentPosition[0][1]

	array2D[curRow][curCol] = "x"

	for {
		// Get current direction
		dir := directions[currentDirection]

		// Check if next position has obstacle
		// if yes, change direction
		// else, mark postion as visited, set next position
		if nextPostitionHasObstacle(array2D, curRow, curCol, dir) {
			currentDirection = transitions[currentDirection]
		} else {
			curRow = curRow + dir[0]
			curCol = curCol + dir[1]

			array2D[curRow][curCol] = "x"
		}

		// Break condition to exit the loop
		if curRow == numberOfRows || curRow == 0 || curCol == numberOfColumns || curCol == 0 {
			fmt.Println("Exiting loop")
			break
		}
	}
	// Print the 2D array
	for i, row := range array2D {
		fmt.Printf("Row %d: %q\n", i, row)
	}

	// Count visited positions
	for i := 0; i < len(array2D); i++ {
		for j := 0; j < len(array2D[i]); j++ {
			if array2D[i][j] == "x" {
				steps++
			}
		}
	}

	fmt.Printf("Steps made %d", steps)
}
