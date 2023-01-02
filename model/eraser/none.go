package eraser

import "github.com/simp7/seatCreator/model"

type none struct{}

func (n none) IsEmpty(x int, y int) bool {
	return false
}

func None() model.Eraser {
	return none{}
}
