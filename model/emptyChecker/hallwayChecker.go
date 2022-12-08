package emptychecker

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

func NewVerticalHallwayChecker(positions ...int) EmptyChecker {
	checker := verticalHallwayChecker{
		positions: positions,
	}
	return checker
}

func NewHorizontalHallwayChecker(positions ...int) EmptyChecker {
	checker := horizontalHallwayChecker{
		positions: positions,
	}
	return checker
}
