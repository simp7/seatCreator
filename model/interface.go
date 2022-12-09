package model

import "github.com/simp7/seatCreator/model/pos"

type EmptyChecker interface {
	IsEmpty(x int, y int) bool
}

type NameInput struct {
	pos.Relative
	SeatType string
}

type NameFormatter interface {
	Long(input NameInput) string
	Short(input NameInput) string
}
