package models

// Stone is a go stone that can either be white, black, or empty
type Stone struct {
	Color        string
	LibertyCount uint8
}

// Intersection is where a stone is placed
type Intersection struct {
	XCoor int
	YCoor int
	Stone Stone
}

// Board contains the intersections data
type Board struct {
	Intersections [][]Intersection
}

// The shape of Board data is a single slice of many row slices that
// contain Intersection structs. This data shape 'resembles' the actual board.
//
// [
//     [ Intersection{XCoor YCoor Stone{}}, Intersection{XCoor YCoor Stone{}} ],
//     [ Intersection{XCoor YCoor Stone{}}, Intersection{XCoor YCoor Stone{}} ],
//     [ Intersection{XCoor YCoor Stone{}}, Intersection{XCoor YCoor Stone{}} ],
//     [ Intersection{XCoor YCoor Stone{}}, Intersection{XCoor YCoor Stone{}} ],
// ]
