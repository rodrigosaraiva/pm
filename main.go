package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Constants were defined by the byte value from the movement character N,S,E,O
const north byte = 78
const south byte = 83
const east byte = 69
const west byte = 79

var invalidMovements int64 = 0

type position struct {
	x int64
	y int64
}

// Assuming by example that the Ash was born or dropped at a random point on the grid and that according to the
// first line of the example at that time he already captures a Pokemon because with just one move passed to the
// program two Pokemons were captured
var capturedPokemons int64 = 1
var ashPosition = position{0, 0}
var gridDriven = []position{{0, 0}}

// Main program
func main() {
	// get start time
	start := time.Now()
	// read movements parameter
	movementsFlag := flag.String("movs", "", "String that describes Ash movements.")
	flag.Parse()
	movements := []byte(strings.ToUpper(*movementsFlag))
	// move Ash
	moveAsh(movements, &ashPosition)
	// get final time
	elapsed := time.Since(start)
	// get memory information
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// print results
	fmt.Println("Movements list:", strings.ToUpper(*movementsFlag))
	fmt.Println("Captured Pokemons:", capturedPokemons)
	fmt.Println("Invalid Movements:", invalidMovements)
	fmt.Println("Elapsed Time:", elapsed)
	fmt.Println("Total memory allocated in Bytes:", m.TotalAlloc)
}

// function that moves Ash in the grid
func moveAsh(movements []byte, ashPosition *position) {
	for _, movement := range movements {
		isValidMovement := getNextValidGridMovement(movement, ashPosition)
		if isValidMovement == false {
			continue
		}
		ashWasHere := checkIfAshWasHere(gridDriven, *ashPosition)
		if ashWasHere == false {
			capturedPokemons++
			gridDriven = append(gridDriven, *ashPosition)
		}
	}
}

// Get the next grid position for Adh and verify if the movement is invalid
// if the movement is invalid the invalidMovements variable is increased
// otherwise the ashPosition variable is updated
func getNextValidGridMovement(movement byte, ashPosition *position) bool {
	switch movement {
	case north:
		ashPosition.y++
		return true
	case south:
		ashPosition.y--
		return true
	case east:
		ashPosition.x++
		return true
	case west:
		ashPosition.x--
		return true
	default:
		invalidMovements += 1
		return false
	}
}

// Function that checks if Ahs already was in this grid position
func checkIfAshWasHere(gridDriven []position, actualPosition position) bool {
	for _, historicalPosition := range gridDriven {
		if historicalPosition == actualPosition {
			return true
		}
	}
	return false
}
