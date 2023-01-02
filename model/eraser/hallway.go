package eraser

import "github.com/simp7/seatCreator/model"

type hallway struct {
	positions []int
}

type verticalHallway hallway

func (v verticalHallway) IsEmpty(x int, y int) bool {
	for _, pos := range v.positions {
		if pos == x {
			return true
		}
	}
	return false
}

type horizontalHallway hallway

func (h horizontalHallway) IsEmpty(x int, y int) bool {
	for _, pos := range h.positions {
		if pos == y {
			return true
		}
	}
	return false
}

func VerticalHallway(positions ...int) model.Eraser {
	return verticalHallway{
		positions: positions,
	}
}

func HorizontalHallway(positions ...int) model.Eraser {
	return horizontalHallway{
		positions: positions,
	}
}
