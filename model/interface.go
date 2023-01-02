package model

import "fmt"

type Eraser interface {
	IsEmpty(x int, y int) bool
}

type NameFormatter interface {
	Long(seat Seat) string
	Short(seat Seat) string
}

type Group interface {
	fmt.Stringer
}
