package emptychecker

type integratedChecker struct {
	checkers []EmptyChecker
}

func (i integratedChecker) IsEmpty(x int, y int) bool {
	for _, checker := range i.checkers {
		if checker.IsEmpty(x, y) {
			return true
		}
	}
	return false
}

func NewIntegratedChecker(checkers ...EmptyChecker) EmptyChecker {
	checker := new(integratedChecker)
	checker.checkers = checkers
	return checker
}
