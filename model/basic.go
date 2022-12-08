package model

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
