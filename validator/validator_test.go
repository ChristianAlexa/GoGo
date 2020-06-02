package validator

import (
	"testing"

	models "github.com/ChristianAlexa/GoGo/models"
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
