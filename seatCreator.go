package main

import (
	"fmt"
	"strings"

	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/emptychecker"
	"github.com/simp7/seatCreator/model/nameformatter"
	"github.com/simp7/seatCreator/model/pos"
	"github.com/simp7/seatCreator/model/row"
)

type Block interface {
	fmt.Stringer
}

type HorizontalBlock struct {
	row        []model.Row
	startPoint pos.Absolute
}

func (b HorizontalBlock) String() string {
	result := make([]string, 0)
	for _, v := range b.row {
		result = append(result, v.String())
	}
	return strings.Join(result, ", ")
}

type BlockInput struct {
	startingPoint model.SeatBase
	xSize         int
	ySize         int
	startIndex    int
	startRow      int
	emptyChecker  model.EmptyChecker
	nameGenerator model.NameFormatter
}

func newHorizontalBlock(input BlockInput) Block {
	rowInput := make([]row.Input, 0)
	rowNumber := input.startRow

	for dy := 1; dy <= input.ySize; dy++ {
		y := input.startingPoint.Y + dy - 1
		emptyPos := make([]int, 0)
		for x := input.startingPoint.X; x < input.startingPoint.X+input.xSize; x++ {
			if input.emptyChecker.IsEmpty(x, y) {
				emptyPos = append(emptyPos, x)
			}
		}

		amount := input.xSize - len(emptyPos)
		if amount == 0 {
			continue
		}

		base := model.NewSeatBase(input.startingPoint.X, y, input.startingPoint.SeatType)
		seat := model.NewSeat(base, pos.Relative{Row: rowNumber, Index: input.startIndex})

		rowInput = append(rowInput, row.Input{
			Criteria:      seat,
			NameFormatter: input.nameGenerator,
			Amount:        input.xSize - len(emptyPos),
			EmptyPos:      emptyPos,
		})
		rowNumber++
	}

	block := new(HorizontalBlock)
	for _, v := range rowInput {
		block.row = append(block.row, row.Horizontal(v))
	}
	block.startPoint = input.startingPoint.Absolute

	return block
}

type integratedBlock struct {
	blocks []Block
}

func (b integratedBlock) String() string {
	result := make([]string, 0)
	for _, v := range b.blocks {
		result = append(result, v.String())
	}
	return strings.Join(result, ", ")
}

func NewIntegratedBlock(blocks ...Block) Block {
	result := new(integratedBlock)
	result.blocks = blocks
	return result
}

func ArtriumSmall() Block {
	hall := emptychecker.VerticalHallway(13, 14)
	rect1 := emptychecker.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 6, Y: 5})
	rect2 := emptychecker.Rectangle(pos.Absolute{X: 4, Y: 15}, pos.Absolute{X: 7, Y: 15})
	specific := emptychecker.Position(pos.Absolute{X: 3, Y: 6})
	integrated := emptychecker.Integrated(hall, rect1, rect2, specific)
	nameGenerator := nameformatter.Standard()

	blockInput := BlockInput{
		startingPoint: model.SeatBase{
			Absolute: pos.Absolute{X: 3, Y: 4},
			SeatType: "A석",
		},
		xSize:         21,
		ySize:         12,
		startIndex:    1,
		startRow:      1,
		emptyChecker:  integrated,
		nameGenerator: nameGenerator,
	}

	return newHorizontalBlock(blockInput)
}

func OperaHouse() Block {
	hall := emptychecker.HorizontalHallway(22)
	rect1 := emptychecker.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 6, Y: 5})
	rect2 := emptychecker.Rectangle(pos.Absolute{X: 4, Y: 15}, pos.Absolute{X: 7, Y: 15})
	specific := emptychecker.Position(pos.Absolute{X: 3, Y: 6})
	integrated := emptychecker.Integrated(hall, rect1, rect2, specific)
	nameGenerator := nameformatter.Standard()

	blockInput := BlockInput{
		startingPoint: model.SeatBase{
			Absolute: pos.Absolute{X: 3, Y: 4},
			SeatType: "A석",
		},
		xSize:         21,
		ySize:         12,
		startIndex:    1,
		startRow:      1,
		emptyChecker:  integrated,
		nameGenerator: nameGenerator,
	}

	return newHorizontalBlock(blockInput)
}

func ConcertHall() Block {
	vHall := emptychecker.VerticalHallway(14, 35)
	hHall := emptychecker.HorizontalHallway(16, 17, 18, 19)
	integrated := emptychecker.Integrated(vHall, hHall)
	nameGenerator := nameformatter.Standard()

	vipHall := emptychecker.VerticalHallway(3, 5, 7, 9, 11, 13, 36, 38, 40, 42, 44, 46)
	empty := emptychecker.Rectangle(pos.Absolute{X: 14, Y: 26}, pos.Absolute{X: 35, Y: 26})
	integrated2 := emptychecker.Integrated(vipHall, empty)

	blockInput := BlockInput{
		startingPoint: model.SeatBase{
			Absolute: pos.Absolute{X: 2, Y: 8},
			SeatType: "A석",
		},
		xSize:         46,
		ySize:         18,
		startIndex:    1,
		startRow:      1,
		emptyChecker:  integrated,
		nameGenerator: nameGenerator,
	}
	blockInput2 := BlockInput{
		startingPoint: model.SeatBase{
			Absolute: pos.Absolute{X: 15, Y: 26},
			SeatType: "A석",
		},
		xSize:         20,
		ySize:         2,
		startIndex:    13,
		startRow:      15,
		emptyChecker:  integrated,
		nameGenerator: nameGenerator,
	}

	blockInput3 := BlockInput{
		startingPoint: model.SeatBase{
			Absolute: pos.Absolute{X: 2, Y: 26},
			SeatType: "VIP석",
		},
		xSize:         46,
		ySize:         1,
		startIndex:    1,
		startRow:      15,
		emptyChecker:  integrated2,
		nameGenerator: nameGenerator,
	}

	block1 := newHorizontalBlock(blockInput)
	block2 := newHorizontalBlock(blockInput2)
	block3 := newHorizontalBlock(blockInput3)
	return NewIntegratedBlock(block1, block2, block3)
}

func main() {
	fmt.Println(ArtriumSmall())
	fmt.Println(ConcertHall())
}
