package validator

import (
	"fmt"

	models "github.com/ChristianAlexa/GoGo/models"
)

// IsEmptyIntersection returns if there's an empty space for a stone
func IsEmptyIntersection(board models.Board, choice models.Intersection) bool {

	intersec := board.Intersections[choice.XCoor-1][choice.YCoor-1]

	if intersec.Stone.Color == "empty" {
		return true
	}

	fmt.Println("There's already a stone there")
	return false
}

func getNeighbors(board models.Board, choice models.Intersection) models.Neighbors {
	neighbors := models.Neighbors{}

	rowChoice := choice.XCoor - 1
	colChoice := choice.YCoor - 1

	// get left neighbor interesection
	isLeftMostStone := false
	if colChoice == 0 {
		isLeftMostStone = true
	}
	if !isLeftMostStone {
		leftNeighbor := board.Intersections[rowChoice][colChoice-1]
		neighbors.Left = leftNeighbor
	}

	// get right neighbor interesection
	isRightMostStone := false
	if colChoice == len(board.Intersections)-1 {
		isRightMostStone = true
	}
	if !isRightMostStone {
		rightNeighbor := board.Intersections[rowChoice][colChoice+1]
		neighbors.Right = rightNeighbor
	}

	// get top neighbor intersection
	isTopRow := false
	if rowChoice == 0 {
		isTopRow = true
	}
	if !isTopRow {
		aboveNeighbor := board.Intersections[rowChoice-1][colChoice]
		neighbors.Above = aboveNeighbor
	}

	// get bottom neighbor intersection
	isBottomRow := false
	if rowChoice == len(board.Intersections)-1 {
		isBottomRow = true
	}
	if !isBottomRow {
		belowNeighbor := board.Intersections[rowChoice+1][colChoice]
		neighbors.Below = belowNeighbor
	}

	return neighbors
}

// IsSurroundedByEnemies returns if an empty interesection has surrounding enemies
func IsSurroundedByEnemies(board models.Board, choice models.Intersection) bool {

	n := getNeighbors(board, choice)

	enemyCount := 0

	if n.Above.Stone.Color != choice.Stone.Color && n.Above.Stone.Color != "empty" {
		enemyCount++
	}

	if n.Below.Stone.Color != choice.Stone.Color && n.Below.Stone.Color != "empty" {
		enemyCount++
	}

	if n.Left.Stone.Color != choice.Stone.Color && n.Left.Stone.Color != "empty" {
		enemyCount++
	}

	if n.Right.Stone.Color != choice.Stone.Color && n.Right.Stone.Color != "empty" {
		enemyCount++
	}

	if enemyCount == 4 {
		fmt.Println("illegal placement: no liberties available.")
		return true
	}

	return false
}
