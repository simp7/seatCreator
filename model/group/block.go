package group

import (
	"strings"

	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/pos"
	"github.com/simp7/seatCreator/model/row"
)

type Block interface {
	String() string
}

type horizontalBlock struct {
	row        []model.Group
	startPoint pos.Absolute
}

func (b horizontalBlock) String() string {
	result := make([]string, 0)
	for _, v := range b.row {
		result = append(result, v.String())
	}
	return strings.Join(result, ", ")
}

type BlockInput struct {
	Criteria      model.Seat
	XSize         int
	YSize         int
	EmptyChecker  model.EmptyChecker
	NameGenerator model.NameFormatter
}

func HorizontalBlock(input BlockInput) Block {
	rowInput := make([]row.Input, 0)
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
		rowInput = append(rowInput, row.Input{
			Criteria:      model.NewSeat(base, input.Criteria.Index, rowNumber),
			NameFormatter: input.NameGenerator,
			Amount:        input.XSize - len(emptyPos),
			EmptyPos:      emptyPos,
		})
		rowNumber++
	}

	for dy := 1; dy <= input.YSize; dy++ {
		appendRowInput(dy)
	}

	block := new(horizontalBlock)
	for _, v := range rowInput {
		block.row = append(block.row, row.Horizontal(v))
	}
	block.startPoint = input.Criteria.Absolute

	return block
}
