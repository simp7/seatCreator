package model

import "fmt"

type Pos struct {
	X int
	Y int
}

type SeatBase struct {
	Pos
	SeatType string `json:"seat_type"`
}

type Seat struct {
	SeatBase
	ShortName string `json:"short_type"`
	Name      string
}

func (s Seat) String() string {
	return fmt.Sprintf("{\n\tx: %d,\n\ty: %d,\n\tseat_type: \"%s\",\n\tshort_name: \"%s\",\n\tname: \"%s\"\n}", s.X, s.Y, s.SeatType, s.ShortName, s.Name)
}
