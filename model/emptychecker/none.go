package emptychecker

import "github.com/simp7/seatCreator/model"

type noneEmptyChecker struct{}

func (n noneEmptyChecker) IsEmpty(x int, y int) bool {
	return false
}

func None() model.EmptyChecker {
	return noneEmptyChecker{}
}
