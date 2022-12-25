package main

import (
	"fmt"
	"strconv"
)

var grid [][]string
var direction string

func main() {
	buildGrid()
	newLevel()
	fmt.Println(grid)
	shiftCheck()
	fmt.Println(grid)
}

func buildGrid() { // This is going cause issues when resetting...no way to reassign the global grid to be empty...although maybe it's not needed

	//for i := 0; i < 30; i++ { // Pointed to same memory address, so when manipulating, updated ever array rather than just the intended one
	//	x = append(x, " ")
	//}
	for z := 0; z < 15; z++ {
		var x []string
		for i := 0; i < 30; i++ { // Space between each item in the array - means the bullet can not collide with anything
			x = append(x, " ")
		}
		grid = append(grid, x)
	}
}

func newLevel() { // For the future if in even array space one sprite, odd the other
	for i := 1; i < 6; i++ {
		for x := 4; x < 26; x += 2 {
			grid[i][x] = strconv.Itoa(i) // string(i) makes fun symbols for some reason
		}
	}
}

func shiftDown() {
	for i := 10; i > 0; i-- {
		for x := 0; x < 30; x++ {
			grid[i+1][x] = grid[i][x]
			grid[i][x] = " "
		}
	}
}

func shiftRight() {
	for i := 10; i > 0; i-- {
		for x := 29; x > 0; x-- {
			grid[i][x] = grid[i][x-1]
		}
	}
}

func shiftLeft() {
	for i := 10; i > 0; i-- {
		for x := 0; x < 29; x++ {
			grid[i][x] = grid[i][x+1]
		}
	}
}

func shiftCheckDown() bool {
	for i := 10; i > 0; i-- {
		if grid[i][0] != " " {
			shiftDown()
			shiftRight()
			direction = "right"
			return true
		} else if grid[i][29] != " " {
			shiftDown()
			shiftLeft()
			direction = "left"
			return true
		}
	}
	return false
}

func shiftCheck() {
	if shiftCheckDown() {
		return
	}
	if direction == "right" {
		shiftRight()
		direction = "right"
	} else {
		shiftLeft()
		direction = "left"
	}
}
