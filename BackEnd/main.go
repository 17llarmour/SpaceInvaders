package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var grid [][]string
var shootyGrid [][]string
var direction string

func main() {
	go runServer()
	grid = buildGrid()
	shootyGrid = buildGrid()
	newLevel()
	for {
		shiftCheck()
		printGrid(grid)
		fmt.Println("------------SPLIT---------------")
		printGrid(shootyGrid)
		time.Sleep(2 * time.Second)
	}
}

func buildGrid() [][]string { // This is going cause issues when resetting...no way to reassign the global grid to be empty...although maybe it's not needed

	//for i := 0; i < 30; i++ { // Pointed to same memory address, so when manipulating, updated ever array rather than just the intended one
	//	x = append(x, " ")
	//}
	var tempGrid [][]string
	for z := 0; z < 15; z++ {
		var x []string
		for i := 0; i < 30; i++ { // Space between each item in the array - means the bullet can not collide with anything
			x = append(x, " ")
		}
		tempGrid = append(tempGrid, x)
	}
	return tempGrid
}

func printGrid(grid [][]string) {
	for i := 0; i < 15; i++ {
		printLn := "["
		for x := 0; x < 30; x++ {
			printLn += grid[i][x]
		}
		fmt.Println(printLn + "]")
	}
}

func newLevel() { // For the future if in even array space one sprite, odd the other
	for i := 1; i < 6; i++ {
		for x := 4; x < 26; x += 2 {
			grid[i][x] = strconv.Itoa(i) // string(i) makes fun symbols for some reason
		}
	}
	addShelter()
}

func addShelter() { // Come back to me
	for offset := 4; offset < 30; offset += 7 {
		return
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
			grid[i][x-1] = " "
		}
	}
}

func shiftLeft() {
	for i := 10; i > 0; i-- {
		for x := 0; x < 29; x++ {
			grid[i][x] = grid[i][x+1]
			grid[i][x+1] = " "
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

func placeUser(box int) {
	for i := 0; i < 29; i++ {
		grid[14][i] = " "
	}
	grid[14][box] = "0"
}

func clearTop() {
	for i := 0; i > 29; i++ {
		grid[0][i] = " "
	}
}

func playerBullet() { // Change how this is done to have a separate grid for bullets
	var pos int
	for i := 0; i < 29; i++ {
		if grid[14][i] == "0" {
			pos = i
			break
		}
	}
	grid[13][pos] = "y"
	for y := 13; y > 0; y-- {
		if grid[y-1][pos] != " " {
			grid[y][pos] = " "
			grid[y-1][pos] = " "
			break
		}
		grid[y-1][pos] = grid[y][pos]
		grid[y][pos] = " "
		//time.Sleep(2 * time.Second)
	}
	clearTop()
}

func runServer() {
	http.HandleFunc("/state", getState)
	http.HandleFunc("/playerPos", updatePos)
	http.HandleFunc("/shoot", playerShot)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(grid)
	if err != nil {
		return
	}
}

func updatePos(w http.ResponseWriter, r *http.Request) { // For pygame, use mouse.get_pos() DIV 30 rounded to get pos
	targetPos := r.URL.Query()["pos"]
	fmt.Println(targetPos)
	box, err := strconv.Atoi(targetPos[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	placeUser(box)

}

func playerShot(w http.ResponseWriter, r *http.Request) {
	playerBullet()
}
