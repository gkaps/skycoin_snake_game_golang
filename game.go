package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Assignment of default board dimensions
var defaultHeight, defaultWidth = 10, 10

type Coordinate struct {
	x int
	y int
}

// Enumeration of board cell values
const (
	Empty    int = 0
	Body     int = 1
	Head     int = 2
	Food     int = 3
	DeadHead int = 4
)

// Game state struct that contains all the needed game state information
type SnakeGameState struct {
	Board          [][]int
	BoardDimension [2]int
	Round          int
	Score          int
	SnakeHead      Coordinate
	SnakeHeadingTo string
	SnakeBody      []Coordinate
	SnakeLength    int
	FoodLocation   Coordinate
}

// This function takes dimensions given from program execution arguments (if any) as input
// Otherwise the default board dimension are used
// The function returns the initial game state
func InitializeGameState(args []string) SnakeGameState {

	var height, width int
	var err1, err2 error

	if len(args) < 3 {
		fmt.Println("Using the default board dimensions")
		height, width = defaultHeight, defaultWidth
	} else {
		height, err1 = strconv.Atoi(args[1])
		width, err2 = strconv.Atoi(args[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Using the default board dimensions")
			height, width = defaultHeight, defaultWidth
		} else {
			fmt.Printf("Using %d by %d board dimensions\n", height, width)
		}
	}

	// Initial Snake Game State
	state := SnakeGameState{
		BoardDimension: [2]int{height, width},
		Round:          0,
		Score:          0,
		SnakeLength:    2,
		SnakeHead:      Coordinate{height / 2, width / 2},
		SnakeBody:      []Coordinate{{height / 2, width/2 - 1}},
		SnakeHeadingTo: "S",
	}

	// Create Board 2D table
	state.Board = make([][]int, height)
	for i := 0; i < height; i++ {
		state.Board[i] = make([]int, width)
	}

	// Get random food location
	// This line is idiomatic Golang. The first part (ok := GetNewFoodLocation(&state))
	// is the assignment and the second part (!ok) is the condition
	if ok := GetNewFoodLocation(&state); !ok {
		fmt.Println("could not get food location while initializing")
	}

	// Populate snake game board
	// Board[y][x] instead of Board[x][y] because height matches to rows and width to columns
	state.Board[state.SnakeHead.y][state.SnakeHead.x] = Head
	state.Board[state.SnakeBody[0].y][state.SnakeBody[0].x] = Body

	return state
}

// The user inputs move direction W (Up) S (Down) A (Left) D (Right) or inputs nothing to keep going in current direction
// The function returns string "W" "A" "S" "D" or ""
func RequireNewMove() string {

	fmt.Printf("Input new direction and press the Enter key: ")
	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()
	userChoice := strings.ToUpper(userInput.Text())

	if userChoice == "W" || userChoice == "S" || userChoice == "A" || userChoice == "D" {
		return userChoice
	}

	return ""
}

// This function takes as input the current state and move, it updates the SnakeGameState
// and returns a bool ok which is false for game over move
func UpdateGameState(state *SnakeGameState, move string) bool {
	// If there is no valid input continue moving to the previous direction
	if move == "" {
		move = state.SnakeHeadingTo
	}

	// For valid input assign to previous head the current head
	// and update snake head values for respective directions
	// If there is opposite move to the current heading
	// don't update snake head and exit the function by returning true value
	previousHead := state.SnakeHead
	switch move {
	case "W":
		if state.SnakeHeadingTo == "S" {
			return true
		}
		state.SnakeHead.y--
	case "S":
		if state.SnakeHeadingTo == "W" {
			return true
		}
		state.SnakeHead.y++
	case "D":
		if state.SnakeHeadingTo == "A" {
			return true
		}
		state.SnakeHead.x++
	case "A":
		if state.SnakeHeadingTo == "D" {
			return true
		}
		state.SnakeHead.x--
	}

	// Check if snake head pass board dimensions
	// in this case end the game (game over)
	if state.SnakeHead.x >= state.BoardDimension[1] || state.SnakeHead.x < 0 || state.SnakeHead.y >= state.BoardDimension[0] || state.SnakeHead.y < 0 {
		state.Board[previousHead.y][previousHead.x] = DeadHead
		return false
	}
	// Check if snake head finds snake body
	// in this case end the game (game over)
	if state.Board[state.SnakeHead.y][state.SnakeHead.x] == Body {
		state.Board[state.SnakeHead.y][state.SnakeHead.x] = DeadHead
		return false
	}

	state.SnakeHeadingTo = move

	// Move snake body
	state.SnakeBody = append(state.SnakeBody, previousHead)

	// Check for food and move snake head
	// Otherwise empty the board element matched to the position related to previous first element of SnakeBody
	// and update SnakeBody by shifting elements' position by 1
	if state.Board[state.SnakeHead.y][state.SnakeHead.x] == Food {
		state.Score++
		ok := GetNewFoodLocation(state)
		if !ok {
			fmt.Println("could not retrieve food location")
		}
	} else {
		state.Board[state.SnakeBody[0].y][state.SnakeBody[0].x] = Empty
		state.SnakeBody = state.SnakeBody[1:len(state.SnakeBody)]
	}

	state.Board[state.SnakeHead.y][state.SnakeHead.x] = Head
	state.Board[previousHead.y][previousHead.x] = Body

	// Update round
	state.Round++

	return true

}

// This function takes state as input and prints board state to terminal
func DisplayGameState(state SnakeGameState) {
	// Display top border as an line of underscores (_)
	topBorder := " "
	for i := 0; i < state.BoardDimension[1]; i++ {
		topBorder += "_"
	}
	fmt.Printf("%s \n", topBorder)
	// Display the main board
	// For empty space " ", for snake body "#", for food location "o" for snake head "&" and for dead snake head "X"
	for i, row := range state.Board {
		fmt.Printf("|")
		for _, col := range row {
			switch col {
			case Empty:
				fmt.Printf(" ")
			case Body:
				fmt.Printf("#")
			case Food:
				fmt.Printf("o")
			case Head:
				fmt.Printf("&")
			case DeadHead:
				fmt.Printf("X")
			}
		}
		fmt.Printf("|")
		// Print round, score and controls right to the board
		switch i {
		case 0:
			fmt.Printf("\tRound: %d", state.Round)
		case 1:
			fmt.Printf("\tScore: %d", state.Score)
		case 2:
			fmt.Printf("\tControls: W: up |S: down |D: right |A: left |other/no input: move on same heading")
		}
		fmt.Printf("\n")
	}
	// Display bottom border as an line of carets (^)
	bottomBorder := " "
	for i := 0; i < state.BoardDimension[1]; i++ {
		bottomBorder += "^"
	}
	fmt.Printf("%s \n", bottomBorder)
}

// This function takes as input a game state and returns a random food location
func GetNewFoodLocation(state *SnakeGameState) bool {
	h, w := state.BoardDimension[0], state.BoardDimension[1]

	for count := 0; count < 10*h*w; count++ {
		y, x := rand.Intn(h), rand.Intn(w)

		// Assign new food location only in empty board cells
		if state.Board[y][x] == Empty {
			state.FoodLocation = Coordinate{y, x}
			state.Board[y][x] = Food
			return true
		}
	}
	return false
}

func main() {
	state := InitializeGameState(os.Args)
	DisplayGameState(state)
	move := RequireNewMove()

	// infinite loop until game over
	for {
		if ok := UpdateGameState(&state, move); !ok {
			DisplayGameState(state)
			fmt.Printf("GAME OVER! The Final Score is: %d\n", state.Score)
			break
		}

		DisplayGameState(state)
		move = RequireNewMove()
	}
}
