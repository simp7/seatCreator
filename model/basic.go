package model

import (
	"fmt"

	"github.com/simp7/seatCreator/model/pos"
)

type SeatBase struct {
	pos.Absolute
	SeatType string
}

func NewSeatBase(x int, y int, seatType string) SeatBase {
	return SeatBase{Absolute: pos.Absolute{X: x, Y: y}, SeatType: seatType}
}

type Seat struct {
	SeatBase
	pos.Relative
}

func (s Seat) String(name NameFormatter) string {
	return fmt.Sprintf("{\n\tx: %d,\n\ty: %d,\n\tseat_type: \"%s\",\n\tshort_name: \"%s\",\n\tname: \"%s\"\n}", s.X, s.Y, s.SeatType, name.Short(s), name.Long(s))
}

func NewSeat(base SeatBase, relativePos pos.Relative) Seat {
	result := Seat{
		SeatBase: base,
		Relative: relativePos,
	}
	return result
}
