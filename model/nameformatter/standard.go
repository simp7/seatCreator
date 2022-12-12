package nameformatter

import (
	"fmt"

	"github.com/simp7/seatCreator/model"
)

func twoDigit(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	}
	return fmt.Sprint(n)
}

type standardFormatter struct{}

func (s standardFormatter) Long(input model.Seat) string {
	return fmt.Sprintf("%s %d열 %s", input.SeatType, input.Row, twoDigit(input.Index+1))
}

func (s standardFormatter) Short(input model.Seat) string {
	return fmt.Sprintf("%c%s", rune(64+input.Row), twoDigit(input.Index+1))
}

func Standard() model.NameFormatter {
	return standardFormatter{}
}
