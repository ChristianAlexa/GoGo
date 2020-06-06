package validator

import (
	"testing"

	"github.com/ChristianAlexa/GoGo/models"
)

func TestIsEmptyIntersection(t *testing.T) {

	t.Run("An empty intersection should return true", func(t *testing.T) {
		mockIntersection := models.Intersection{
			XCoor: 1,
			YCoor: 1,
			Stone: models.Stone{Color: "empty", LibertyCount: 4},
		}

		mockRow := []models.Intersection{mockIntersection}

		mockBoard := models.Board{
			Intersections: [][]models.Intersection{mockRow},
			WhiteGroups:   []models.Group{},
			BlackGroups:   []models.Group{},
		}

		mockChoice := models.Intersection{
			XCoor: 1,
			YCoor: 1,
			Stone: models.Stone{Color: "white", LibertyCount: 4},
		}

		actual := IsEmptyIntersection(mockBoard, mockChoice)
		expected := true

		if actual != expected {
			t.Errorf("got %v and expected %v", actual, expected)
		}
	})

	t.Run("An occupied intersection should return false", func(t *testing.T) {
		mockIntersection := models.Intersection{
			XCoor: 1,
			YCoor: 1,
			Stone: models.Stone{Color: "white", LibertyCount: 4},
		}

		mockRow := []models.Intersection{mockIntersection} // [ {XCoor, YCoor, Stone} ]

		mockBoard := models.Board{
			Intersections: [][]models.Intersection{mockRow},
			WhiteGroups:   []models.Group{},
			BlackGroups:   []models.Group{},
		}

		mockChoice := models.Intersection{
			XCoor: 1,
			YCoor: 1,
			Stone: models.Stone{Color: "white", LibertyCount: 4},
		}

		actual := IsEmptyIntersection(mockBoard, mockChoice)
		expected := false

		if actual != expected {
			t.Errorf("got %v and expected %v", actual, expected)
		}
	})

}

func TestIsEnemyStone(t *testing.T) {

	t.Run("A black stone and a white stone are enemies", func(t *testing.T) {

		mockBlackStone := models.Stone{
			Color:        "black",
			LibertyCount: 4,
		}

		mockWhiteStone := models.Stone{
			Color:        "white",
			LibertyCount: 4,
		}

		actual := isEnemyStone(mockBlackStone.Color, mockWhiteStone.Color)
		expected := true

		if actual != expected {
			t.Errorf("got %v and expected %v", actual, expected)
		}
	})

	t.Run("Two white stones are not enemies", func(t *testing.T) {

		mockWhiteStone1 := models.Stone{
			Color:        "white",
			LibertyCount: 4,
		}

		mockWhiteStone2 := models.Stone{
			Color:        "white",
			LibertyCount: 4,
		}

		actual := isEnemyStone(mockWhiteStone1.Color, mockWhiteStone2.Color)
		expected := false

		if actual != expected {
			t.Errorf("got %v and expected %v", actual, expected)
		}
	})

	t.Run("If one stone color is empty, there is no enemy", func(t *testing.T) {

		mockEmptyStone := models.Stone{
			Color:        "empty",
			LibertyCount: 4,
		}

		mockWhiteStone := models.Stone{
			Color:        "white",
			LibertyCount: 4,
		}

		actual := isEnemyStone(mockEmptyStone.Color, mockWhiteStone.Color)
		expected := false

		if actual != expected {
			t.Errorf("got %v and expected %v", actual, expected)
		}
	})
}

func TestIsFriendlyStone(t *testing.T) {

	t.Run("A black stone and a white stone are not friends", func(t *testing.T) {

		mockBlackStone := models.Stone{
			Color:        "black",
			LibertyCount: 4,
		}

		mockWhiteStone := models.Stone{
			Color:        "white",
			LibertyCount: 4,
		}

		actual := isFriendlyStone(mockBlackStone.Color, mockWhiteStone.Color)
		expected := false

		if actual != expected {
			t.Errorf("got %v and expected %v", actual, expected)
		}
	})

	t.Run("Two white stones are friends", func(t *testing.T) {

		mockWhiteStone1 := models.Stone{
			Color:        "white",
			LibertyCount: 4,
		}

		mockWhiteStone2 := models.Stone{
			Color:        "white",
			LibertyCount: 4,
		}

		actual := isFriendlyStone(mockWhiteStone1.Color, mockWhiteStone2.Color)
		expected := true

		if actual != expected {
			t.Errorf("got %v and expected %v", actual, expected)
		}
	})

	t.Run("If one stone color is empty, there is no friend", func(t *testing.T) {

		mockEmptyStone := models.Stone{
			Color:        "empty",
			LibertyCount: 4,
		}

		mockWhiteStone := models.Stone{
			Color:        "white",
			LibertyCount: 4,
		}

		actual := isFriendlyStone(mockEmptyStone.Color, mockWhiteStone.Color)
		expected := false

		if actual != expected {
			t.Errorf("got %v and expected %v", actual, expected)
		}
	})
}

func TestGetNeighbors(t *testing.T) {
	t.Run("Should return neighbors directly surrounding a player intersection choice", func(t *testing.T) {

		boardSize := 19
		var intersections = make([][]models.Intersection, boardSize)
		mockBoard := models.Board{
			Intersections: intersections,
			WhiteGroups:   []models.Group{},
			BlackGroups:   []models.Group{},
		}

		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				st := models.Stone{Color: "empty", LibertyCount: 4}
				intersection := models.Intersection{XCoor: i + 1, YCoor: j, Stone: st}
				intersections[i] = append(intersections[i], intersection)
			}
		}

		mockChoice := models.Intersection{
			XCoor: 4,
			YCoor: 4,
			Stone: models.Stone{Color: "white", LibertyCount: 4},
		}

		actual := getNeighbors(mockBoard, mockChoice)
		expected := map[string]models.Intersection{
			"ABOVE": {XCoor: 3, YCoor: 3, Stone: models.Stone{Color: "empty", LibertyCount: 4}},
			"BELOW": {XCoor: 5, YCoor: 3, Stone: models.Stone{Color: "empty", LibertyCount: 4}},
			"LEFT":  {XCoor: 4, YCoor: 2, Stone: models.Stone{Color: "empty", LibertyCount: 4}},
			"RIGHT": {XCoor: 4, YCoor: 4, Stone: models.Stone{Color: "empty", LibertyCount: 4}},
		}
		if actual["ABOVE"] != expected["ABOVE"] {
			t.Errorf("got %v and expected %v", actual["ABOVE"], expected["ABOVE"])
		}
		if actual["BELOW"] != expected["BELOW"] {
			t.Errorf("got %v and expected %v", actual["BELOW"], expected["BELOW"])
		}
		if actual["LEFT"] != expected["LEFT"] {
			t.Errorf("got %v and expected %v", actual["LEFT"], expected["LEFT"])
		}
		if actual["RIGHT"] != expected["RIGHT"] {
			t.Errorf("got %v and expected %v", actual["RIGHT"], expected["RIGHT"])
		}
	})
}
