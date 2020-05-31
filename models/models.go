package models

// Stone is a go stone that can either be white, black, or empty
type Stone struct {
	Color        string
	LibertyCount int
}

// Intersection is where a stone is placed
type Intersection struct {
	XCoor int
	YCoor int
	Stone Stone
}

// Group is a collection of friendly stones that are touching
type Group struct {
	Intersections []Intersection
	LibertyCount  int
}

// Board contains the intersections data
type Board struct {
	Intersections [][]Intersection
	WhiteGroups   Group
	BlackGroups   Group
}
