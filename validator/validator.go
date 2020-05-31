package validator

import (
	"fmt"

	models "github.com/ChristianAlexa/GoGo/models"
)

// IsEmptyIntersection returns if there's an empty space for a stone
func IsEmptyIntersection(board models.Board, choice models.Intersection) bool {

	chosenIntersection := board.Intersections[choice.XCoor-1][choice.YCoor-1]

	if chosenIntersection.Stone.Color == "empty" {
		return true
	}

	fmt.Println(IntersectionOccupied)
	return false
}

// getNeighbors returns a map of neighbors surrounding the chosen intersection
func getNeighbors(board models.Board, choice models.Intersection) map[string]models.Intersection {

	neighborsMap := map[string]models.Intersection{
		"ABOVE": {},
		"BELOW": {},
		"LEFT":  {},
		"RIGHT": {},
	}

	rowChoice := choice.XCoor - 1
	colChoice := choice.YCoor - 1

	// get left neighbor intersection
	isLeftMostStone := false
	if colChoice == 0 {
		isLeftMostStone = true
	}
	if !isLeftMostStone {
		leftNeighbor := board.Intersections[rowChoice][colChoice-1]
		neighborsMap["LEFT"] = leftNeighbor
	}

	// get right neighbor intersection
	isRightMostStone := false
	if colChoice == len(board.Intersections)-1 {
		isRightMostStone = true
	}
	if !isRightMostStone {
		rightNeighbor := board.Intersections[rowChoice][colChoice+1]
		neighborsMap["RIGHT"] = rightNeighbor
	}

	// get top neighbor intersection
	isTopRow := false
	if rowChoice == 0 {
		isTopRow = true
	}
	if !isTopRow {
		aboveNeighbor := board.Intersections[rowChoice-1][colChoice]
		neighborsMap["ABOVE"] = aboveNeighbor
	}

	// get bottom neighbor intersection
	isBottomRow := false
	if rowChoice == len(board.Intersections)-1 {
		isBottomRow = true
	}
	if !isBottomRow {
		belowNeighbor := board.Intersections[rowChoice+1][colChoice]
		neighborsMap["BELOW"] = belowNeighbor
	}

	return neighborsMap
}

// isEnemyStone returns if 2 stones are the same color or not
func isEnemyStone(stoneColor1 string, stoneColor2 string) bool {

	if stoneColor1 == "empty" || stoneColor2 == "empty" {
		return false
	}

	if stoneColor1 != stoneColor2 {
		return true
	}

	return false
}

// IsSurroundedByEnemies returns if an empty intersection has surrounding enemies
func IsSurroundedByEnemies(board models.Board, choice models.Intersection) bool {
	nMap := getNeighbors(board, choice)

	enemyCount := 0

	for _, n := range nMap {
		if isEnemyStone(n.Stone.Color, choice.Stone.Color) {
			enemyCount++
		}
	}

	if enemyCount == 4 {
		fmt.Println(NoLibertiesAvailable)
		return true
	}

	return false
}
