package model

import (
	"fmt"

	"github.com/simp7/seatCreator/model/pos"
)

type SeatBase struct {
	pos.Absolute
	SeatType string
}

type Seat struct {
	SeatBase
	ShortName string
	Name      string
}

func (s Seat) String() string {
	return fmt.Sprintf("{\n\tx: %d,\n\ty: %d,\n\tseat_type: \"%s\",\n\tshort_name: \"%s\",\n\tname: \"%s\"\n}", s.X, s.Y, s.SeatType, s.ShortName, s.Name)
}

func NewSeat(base SeatBase, shortName string, name string) Seat {
	result := Seat{
		SeatBase:  base,
		ShortName: shortName,
		Name:      name,
	}
	return result
}
