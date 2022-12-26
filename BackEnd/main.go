package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var grid [][]string
var shootyGrid [][]string
var direction string
var lives = 3
var score int

func main() {
	go runServer()
	grid = buildGrid()
	shootyGrid = buildGrid()
	newLevel()

	for lives > 0 {
		shiftCheck()
		printGrid(grid)
		fmt.Println("------------SPLIT---------------")
		printGrid(shootyGrid)
		fmt.Println("------------SPLIT---------------")
		winCheck()
		//time.Sleep(2 * time.Second)
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

//func endLevel() {
//	//for i := 10; i > 0; i-- {
//	//	for x := 0; x < 30; x++ {
//	//		if grid[i][x] != " " {
//	//			return
//	//		}
//	//	}
//	//}
//
//	if winCheck() {
//		newLevel()
//	}
//}

func winCheck() {
	for x := 0; x < 30; x++ {
		if grid[11][x] != " " {
			lives -= 1
			fmt.Println("Life Lost")
			placeUser(0)
			clearGrid()
			newLevel()
			return
		}
	}
}

func clearGrid() {
	for i := 11; i > 0; i-- {
		for x := 29; x > 0; x-- {
			grid[i][x] = " "
			shootyGrid[i][x] = " "
		}
	}
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

func playerBullet() { // Change how this is done to have a separate grid for bullets
	var pos int
	for i := 0; i < 29; i++ {
		if grid[14][i] == "0" {
			pos = i
			break
		}
	}
	shootyGrid[13][pos] = "y"
	for y := 13; y > 0; y-- {
		if grid[y-1][pos] != " " {
			shootyGrid[y][pos] = " "
			pointsUpdate(y, pos)
			grid[y-1][pos] = " "
			break
		}
		shootyGrid[y-1][pos] = shootyGrid[y][pos]
		shootyGrid[y][pos] = " "

		time.Sleep(2 * time.Second) // This needs to sync with the normal game clock potentially not entirely sure...
	}
	clearTop()
}

func invaderBullet() {
	shootRow := rand.Intn(9) + 1
	shootCol := rand.Intn(29)
	if grid[shootRow][shootCol] == "1" || grid[shootRow][shootCol] == "2" {
		shootyGrid[shootRow][shootCol] = "p1"
	} else if grid[shootRow][shootCol] == "3" || grid[shootRow][shootCol] == "4" {
		shootyGrid[shootRow][shootCol] = "p2"
	} else if grid[shootRow][shootCol] == "5" {
		shootyGrid[shootRow][shootCol] = "p3"
	}
}

func bulletUpdates() {

}

func pointsUpdate(y int, x int) {
	if grid[y-1][x] == "1" || grid[y-1][x] == "2" {
		score += 30
	} else if grid[y-1][x] == "3" || grid[y-1][x] == "4" {
		score += 20
	} else if grid[y-1][x] == "5" {
		score += 10
	} else if grid[y-1][x] == "6" {
		multi := rand.Intn(3) + 1
		score += 100 * multi
	}
}

func clearTop() {
	for i := 0; i < 29; i++ { // Syntax, inequality sign was backwards -_-
		shootyGrid[0][i] = " "
	}
}

func runServer() {
	http.HandleFunc("/state", getState)
	http.HandleFunc("/shootyState", getShootyState)
	http.HandleFunc("/playerPos", updatePos)
	http.HandleFunc("/shoot", playerShot)
	http.HandleFunc("/reset", resetCheck)

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
		fmt.Println(err)
		return
	}
}

func getShootyState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(shootyGrid)
	if err != nil {
		fmt.Println(err)
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
	beenShot := r.URL.Query()["shoot"]
	fmt.Println(beenShot)
	shot := beenShot[0]
	if shot == "yes" {
		playerBullet()
	}
}

func resetCheck(w http.ResponseWriter, r *http.Request) {
	reset := r.URL.Query()["reset"]

	fmt.Println(reset)
	if reset[0] == "yes" {
		lives = 3
		score = 0
		newLevel()
		//player = ""

	}
}
