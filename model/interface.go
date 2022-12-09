package model

type EmptyChecker interface {
	IsEmpty(x int, y int) bool
}

type NameFormatter interface {
	Long(seat Seat) string
	Short(seat Seat) string
}

type Row interface {
	String() string
}
