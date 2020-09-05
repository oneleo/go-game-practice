package transform

// Position is
type Position struct {
	Tx float64
	Ty float64
}

// Scale is
type Scale struct {
	Width  int
	Height int
}

// Rotation rotates the matrix by theta.
// The unit is radian.
type Rotation struct {
	Theta float64
}
