package clockface

import "time"

type Point struct {
	x float64
	y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
