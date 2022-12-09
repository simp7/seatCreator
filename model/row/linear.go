package row

import (
	"strings"

	"github.com/simp7/seatCreator/model"
)

func relativePositionList(start int, amount int, emptyList []int) []int {
	result := make([]int, 0)
	current := start
	emptyPosIndex := 0
	for i := 1; i <= amount; i++ {
		for emptyPosIndex < len(emptyList) && current == emptyList[emptyPosIndex] {
			emptyPosIndex++
			current++
		}
		result = append(result, current)
		current++
	}
	return result
}

type linear struct {
	seats         []model.Seat
	nameFormatter model.NameFormatter
}

func (r linear) String() string {

	var result = make([]string, 0)

	for _, v := range r.seats {
		result = append(result, v.String(r.nameFormatter))
	}
	return strings.Join(result, ", ")
}

type Input struct {
	Criteria      model.Seat
	NameFormatter model.NameFormatter
	Amount        int
	EmptyPos      []int
}

func Horizontal(input Input) model.Group {
	y := input.Criteria.Y
	seatType := input.Criteria.SeatType
	seats := make([]model.Seat, 0)

	posList := relativePositionList(input.Criteria.X, input.Amount, input.EmptyPos)

	for i, x := range posList {
		inputBase := model.NewSeatBase(x, y, seatType)
		seats = append(seats, model.NewSeat(inputBase, i+input.Criteria.Index-1, input.Criteria.Row))
	}

	row := new(linear)

	row.seats = seats
	row.nameFormatter = input.NameFormatter
	return row
}
