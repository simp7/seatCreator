package emptychecker

import "github.com/simp7/seatCreator/model"

type hallwayChecker struct {
	positions []int
}

type verticalHallwayChecker hallwayChecker

func (v verticalHallwayChecker) IsEmpty(x int, y int) bool {
	for _, pos := range v.positions {
		if pos == x {
			return true
		}
	}
	return false
}

type horizontalHallwayChecker hallwayChecker

func (h horizontalHallwayChecker) IsEmpty(x int, y int) bool {
	for _, pos := range h.positions {
		if pos == x {
			return true
		}
	}
	return false
}

func VerticalHallway(positions ...int) model.EmptyChecker {
	checker := verticalHallwayChecker{
		positions: positions,
	}
	return checker
}

func HorizontalHallway(positions ...int) model.EmptyChecker {
	checker := horizontalHallwayChecker{
		positions: positions,
	}
	return checker
}
