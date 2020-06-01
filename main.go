package main

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	models "github.com/ChristianAlexa/GoGo/models"
	validator "github.com/ChristianAlexa/GoGo/validator"
)

// initBoard creates a go board of any size populated with 'empty' stones
func initBoard(boardSize int, logger zap.Logger) models.Board {

	logger.Info("initializing board")

	if !(boardSize == 9 || boardSize == 13 || boardSize == 19) {
		fmt.Println(validator.InvalidBoardSize)
		os.Exit(1)
	}

	var intersections = make([][]models.Intersection, boardSize)
	var ret = models.Board{
		Intersections: intersections,
		WhiteGroups:   models.Group{},
		BlackGroups:   models.Group{},
	}

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			st := models.Stone{Color: "empty", LibertyCount: 4}
			intersection := models.Intersection{XCoor: i + 1, YCoor: j, Stone: st}
			intersections[i] = append(intersections[i], intersection)
		}
	}

	return ret
}

// printBoardUI prints each row as a new line to the console
func printBoardUI(b models.Board) {

	// clear console after each played move
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// display Board info
	fmt.Printf("%[1]vx%[1]v Board\n", len(b.Intersections))

	fmt.Printf("--  ")

	// display col board letters, supporting up to max board size (19x19)
	boardLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S"}

	dynamicBoardLetters := boardLetters[:len(b.Intersections)]

	for l := 0; l < len(dynamicBoardLetters); l++ {
		fmt.Printf("%v ", dynamicBoardLetters[l])
	}
	fmt.Printf("\n")

	middleCoor := (len(b.Intersections) + 1) / 2

	// display row number and intersections
	for i, r := range b.Intersections {

		// consistent row number spacing
		if i+1 < 10 {
			fmt.Printf("0%v ", i+1)
		} else {
			fmt.Printf("%v ", i+1)
		}

		for _, intersec := range r {
			switch intersec.Stone.Color {
			case "white":
				fmt.Printf("%s", "⚪")
			case "black":
				fmt.Printf("%s", "⚫")
			case "empty":
				if intersec.XCoor == middleCoor && intersec.YCoor+1 == middleCoor {
					fmt.Printf("%s", "🔆")
				} else {
					fmt.Printf("%s", "➕")
				}
			}
		}
		fmt.Printf("\n")
	}
}

// playMove finds player's chosen intersection and updates that intersection with a stone
func playMove(b models.Board, choice models.Intersection) models.Board {
	fmt.Printf("\n%s plays at %d %d\n", choice.Stone.Color, choice.XCoor, choice.YCoor)

	rowChoice := choice.XCoor - 1
	colChoice := choice.YCoor - 1

	// jump to row, jump to column, and update intersection
	targetRow := b.Intersections[rowChoice]
	targetRow[colChoice].Stone = choice.Stone

	return b
}

// promptRowChoice prompts the user for the row they would like to play
func promptRowChoice(b models.Board) string {
	fmt.Println(">> Choose a row number between 1 and", len(b.Intersections))
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	return input
}

// promptColChoice prompts the user for the column they would like to play
func promptColChoice() string {
	fmt.Println("pick col:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	return input
}

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("can't initialize zap logger")
	}
	defer logger.Sync()

	hasWinner := false
	currentTurnColor := "black"

	// TODO: prompt user for what size game they want to play
	b := initBoard(19, *logger)

	printBoardUI(b)

	for !hasWinner {

		var rowChoice int
		var colChoice int
		var playerChoice models.Intersection

		// continue to prompt row and col choice if choices are invalid
		for {

			// prompt player for row choice
			for {

				rowInput := promptRowChoice(b)

				rowInputInt, err := strconv.Atoi(rowInput)
				if err != nil {
					fmt.Println(">> Invalid input")
				}

				isValidRow := rowInputInt >= 1 && rowInputInt <= len(b.Intersections)

				if isValidRow {
					rowChoice = rowInputInt
					break
				}
			}

			// prompt player for column choice
			for {
				colInput := promptColChoice()

				colMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19}
				colChoiceInt := colMap[colInput]

				isValidCol := colChoiceInt >= 1 && colChoiceInt <= len(b.Intersections)

				if isValidCol {
					colChoice = colChoiceInt
					break
				}

				fmt.Println("Choose a column letter")
			}

			playerChoice = models.Intersection{
				XCoor: rowChoice,
				YCoor: colChoice,
				Stone: models.Stone{Color: currentTurnColor, LibertyCount: 4},
			}

			isLegalMove := validator.IsEmptyIntersection(b, playerChoice) &&
				!validator.IsSurroundedByEnemies(b, playerChoice)

			if isLegalMove {
				break
			}
		}

		playMove(b, playerChoice)
		printBoardUI(b)

		// toggle player move
		switch currentTurnColor {
		case "white":
			currentTurnColor = "black"
		case "black":
			currentTurnColor = "white"
		}

		// hasWinner = true
	}
}
