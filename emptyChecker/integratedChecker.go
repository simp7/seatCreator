package emptychecker

type IntegratedChecker struct {
	checkers []EmptyChecker
}

func (i IntegratedChecker) IsEmpty(x int, y int) bool {
	for _, checker := range i.checkers {
		if checker.IsEmpty(x, y) {
			return true
		}
	}
	return false
}

func NewIntegratedChecker(checkers ...EmptyChecker) EmptyChecker {
	checker := new(IntegratedChecker)
	checker.checkers = checkers
	return checker
}
