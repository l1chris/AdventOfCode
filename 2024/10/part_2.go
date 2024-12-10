package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var directions = [][2]int{
	{0, 1},  // Right
	{0, -1}, // Left
	{1, 0},  // Down
	{-1, 0}, // Up
}

func findRoutes(grid [][]int) [][][2]int {
	rows := len(grid)
	cols := len(grid[0])
	var routes [][][2]int
	var currentRoute [][2]int

	var dfs func(x, y int)

	dfs = func(x, y int) {
		// Add the current position to the route
		currentRoute = append(currentRoute, [2]int{x, y})

		// If we reach a cell with value 9, store the route
		if grid[x][y] == 9 {
			// Copy the current route to avoid mutation
			routeCopy := make([][2]int, len(currentRoute))
			copy(routeCopy, currentRoute)
			routes = append(routes, routeCopy)
		} else {
			// Explore all valid neighbors
			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]
				// Check boundaries and ensure the next digit is strictly greater
				if nx >= 0 && nx < rows && ny >= 0 && ny < cols && (grid[nx][ny]-grid[x][y] == 1) {
					dfs(nx, ny)
				}
			}
		}

		// Remove last element from the slice to backtrack
		currentRoute = currentRoute[:len(currentRoute)-1]
	}

	// Start DFS from all positions containing 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				dfs(i, j)
			}
		}
	}

	return routes
}

func groupRoutes(routes [][][2]int) map[string][][][2]int {
	groupedRoutes := make(map[string][][][2]int)

	for _, route := range routes {
		if len(route) < 1 {
			continue // Skip routes with no entries
		}

		// Create a key based on the first
		key := fmt.Sprintf("%d", route[0])

		// Add the route to the group
		groupedRoutes[key] = append(groupedRoutes[key], route)
	}

	return groupedRoutes
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var array2D [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var row []int

		for _, ch := range line {
			num, err := strconv.Atoi(string(ch))
			if err != nil {
				return
			}
			row = append(row, num)
		}

		array2D = append(array2D, row)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the 2D array
	//for _, row := range array2D {
	//	fmt.Println(row)
	//}

	routes := findRoutes(array2D)

	// Print the routes
	fmt.Println("All Routes:")
	for _, route := range routes {
		fmt.Println(route)
	}

	grouped := groupRoutes(routes)

	// Trailhead: Find distinct endings for a starting point like [0 2]

	var sumOfScores int

	for key, group := range grouped {
		fmt.Printf("Key: %s\n", key)
		fmt.Println("Routes:")
		for _, route := range group {
			fmt.Println(route)
		}

		fmt.Println("len(group):%d", len(group))

		sumOfScores += len(group)
	}

	fmt.Printf("SumofScores: %d\n", sumOfScores)
}
