package model

type NameInput struct {
	Row      int
	Index    int
	SeatType string
}

type NameFormatter interface {
	Long(input NameInput) string
	Short(input NameInput) string
}
