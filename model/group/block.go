package group

import (
	"strings"

	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/pos"
)

type block struct {
	row        []model.Group
	startPoint pos.Absolute
}

func (b block) String() string {
	result := make([]string, 0)
	for _, v := range b.row {
		result = append(result, v.String())
	}
	return strings.Join(result, ", ")
}

func (b block) Html() string {
	result := make([]string, 0)
	for _, v := range b.row {
		result = append(result, v.Html())
	}
	return strings.Join(result, "\n")
}

type BlockInput struct {
	Criteria      model.Seat
	XSize         int
	YSize         int
	EmptyChecker  model.Eraser
	NameFormatter model.NameFormatter
	Reverse       bool
}

func HorizontalBlock(input BlockInput) model.Group {
	rowInput := make([]Input, 0)
	rowNumber := input.Criteria.Row

	filterEmptyPos := func(y int) []int {
		result := make([]int, 0)

		for x := input.Criteria.X; x < input.Criteria.X+input.XSize; x++ {
			if input.EmptyChecker.IsEmpty(x, y) {
				result = append(result, x)
			}
		}
		return result
	}

	appendRowInput := func(dy int) {
		y := input.Criteria.Y + dy - 1
		emptyPos := filterEmptyPos(y)
		amount := input.XSize - len(emptyPos)
		if amount == 0 {
			return
		}

		base := model.NewSeatBase(input.Criteria.X, y, input.Criteria.SeatType)
		rowInput = append(rowInput, Input{
			Criteria:      model.NewSeat(base, input.Criteria.Index, rowNumber),
			NameFormatter: input.NameFormatter,
			Amount:        input.XSize - len(emptyPos),
			EmptyPos:      emptyPos,
		})
		rowNumber++
	}

	for dy := 1; dy <= input.YSize; dy++ {
		appendRowInput(dy)
	}

	block := new(block)
	for _, v := range rowInput {
		if input.Reverse {
			block.row = append(block.row, HorizontalReverseRow(v))
			continue
		}
		block.row = append(block.row, HorizontalRow(v))
	}
	block.startPoint = input.Criteria.Absolute

	return block
}

func VerticalBlock(input BlockInput) model.Group {
	rowInput := make([]Input, 0)
	rowNumber := input.Criteria.Row

	filterEmptyPos := func(x int) []int {
		result := make([]int, 0)

		for y := input.Criteria.Y; y < input.Criteria.Y+input.YSize; y++ {
			if input.EmptyChecker.IsEmpty(x, y) {
				result = append(result, y)
			}
		}
		return result
	}

	appendRowInput := func(dx int) {
		x := input.Criteria.X + dx - 1
		emptyPos := filterEmptyPos(x)
		amount := input.XSize - len(emptyPos)
		if amount == 0 {
			return
		}

		base := model.NewSeatBase(x, input.Criteria.Y, input.Criteria.SeatType)
		rowInput = append(rowInput, Input{
			Criteria:      model.NewSeat(base, input.Criteria.Index, rowNumber),
			NameFormatter: input.NameFormatter,
			Amount:        input.YSize - len(emptyPos),
			EmptyPos:      emptyPos,
		})
		rowNumber++
	}

	for dx := 1; dx <= input.XSize; dx++ {
		appendRowInput(dx)
	}

	block := new(block)
	for _, v := range rowInput {
		block.row = append(block.row, VerticalRow(v))
	}
	block.startPoint = input.Criteria.Absolute

	return block
}
