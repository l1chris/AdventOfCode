package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var input string

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		input = scanner.Text()
		fmt.Println("First line:", input)
	}

	type block struct {
		val int
	}

	blocks := []block{}
	currNum := 0
	doSpace := false

	for _, c := range input { // For each char in the disk map
		num, _ := strconv.Atoi(string(c))
		for k := 0; k < num; k++ { // Use the num to control the num of chars appearing
			if doSpace {
				blocks = append(blocks, block{-1})
			} else {
				blocks = append(blocks, block{currNum})
			}
		}
		if !doSpace {
			currNum += 1
		}
		doSpace = !doSpace
	}
	fmt.Println("blocks:", blocks)

	// Two pointers
	left := 0
	right := len(blocks) - 1

	for left < right {
		// Find the next dot from the left
		for left < len(blocks) && blocks[left].val != -1 {
			left++
		}

		// Find the next non-dot from the right
		for right >= 0 && blocks[right].val == -1 {
			right--
		}

		// If pointers are valid, replace the dot with the right-hand character
		if left < right {
			blocks[left] = blocks[right]
			blocks[right].val = -1
			left++
			right--
		}
	}

	fmt.Println("blocks:", blocks)

	var checksum int

	for i := 0; i < len(blocks)-1; i++ {
		if blocks[i].val != -1 {
			checksum += i * blocks[i].val
		}
	}

	fmt.Println("checksum:", checksum)

}
