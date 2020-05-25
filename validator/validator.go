package validator

import models "github.com/ChristianAlexa/GoGo/models"

// IsEmptyIntersection determines if an intersection can receive a stone
func IsEmptyIntersection(board models.Board, choice models.Intersection) bool {

	intersec := board.Intersections[choice.XCoor-1][choice.YCoor-1]

	if intersec.Stone.Color == "empty" {
		return true
	}

	return false
}
