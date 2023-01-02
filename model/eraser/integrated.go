package eraser

import "github.com/simp7/seatCreator/model"

type integrated struct {
	erasers []model.Eraser
}

func (i integrated) IsEmpty(x int, y int) bool {
	for _, eraser := range i.erasers {
		if eraser.IsEmpty(x, y) {
			return true
		}
	}
	return false
}

func Integrated(erasers ...model.Eraser) model.Eraser {
	output := new(integrated)
	output.erasers = erasers
	return output
}
