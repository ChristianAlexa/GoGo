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
