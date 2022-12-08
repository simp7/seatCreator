package emptychecker

import "github.com/simp7/seatCreator/model"

type integratedChecker struct {
	checkers []model.EmptyChecker
}

func (i integratedChecker) IsEmpty(x int, y int) bool {
	for _, checker := range i.checkers {
		if checker.IsEmpty(x, y) {
			return true
		}
	}
	return false
}

func Integrated(checkers ...model.EmptyChecker) model.EmptyChecker {
	checker := new(integratedChecker)
	checker.checkers = checkers
	return checker
}
