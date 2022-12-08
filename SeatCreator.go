package main

import (
	"fmt"
	"strings"

	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/emptychecker"
)

func twoDigit(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	}
	return fmt.Sprint(n)
}

func newSeat(base model.SeatBase, generateShortName func(_ model.SeatBase) string, generateName func(_ model.SeatBase) string) model.Seat {
	result := model.Seat{
		SeatBase:  base,
		ShortName: generateShortName(base),
		Name:      generateName(base),
	}
	return result
}

type Row struct {
	Seats []model.Seat
}

func (r Row) String() string {

	var result = make([]string, 0)

	for _, v := range r.Seats {
		result = append(result, v.String())
	}
	return strings.Join(result, ", ")
}

type RowInput struct {
	criteria      model.SeatBase
	initialNumber int
	amount        int
	emptyPos      []int
	columnNumber  int
}

func getPositionList(start int, amount int, emptyList []int) []int {
	result := make([]int, 0)
	emptyPosIndex := 0
	for seatNumber := 1; seatNumber <= amount; seatNumber++ {
		for emptyPosIndex < len(emptyList) && start == emptyList[emptyPosIndex] {
			emptyPosIndex++
			start++
		}
		result = append(result, start)
		start++
	}
	return result
}

func newVerticalRow(input RowInput) *Row {
	return nil
}

func newHorizontalRow(input RowInput) *Row {
	y := input.criteria.Y
	seatType := input.criteria.SeatType
	seats := make([]model.Seat, 0)

	posList := getPositionList(input.criteria.X, input.amount, input.emptyPos)

	for i, x := range posList {
		generateShortName := func(target model.SeatBase) string {
			return fmt.Sprintf("%c%s", rune(64+input.columnNumber), twoDigit(i+1))
		}

		generateName := func(target model.SeatBase) string {
			return fmt.Sprintf("%s %d열 %s", target.SeatType, input.columnNumber, twoDigit(i+1))
		}

		inputBase := model.SeatBase{
			Pos:      model.Pos{X: x, Y: y},
			SeatType: seatType,
		}
		seats = append(seats, newSeat(inputBase, generateShortName, generateName))
	}

	row := new(Row)

	row.Seats = seats
	return row
}

type Blocks interface {
	fmt.Stringer
}

type Block struct {
	row []*Row
}

func (b Block) String() string {
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
	emptyChecker  model.EmptyChecker
}

func newBlock(input BlockInput) *Block {
	rowInput := make([]RowInput, 0)

	for columnNumber := 1; columnNumber <= input.ySize; columnNumber++ {
		y := input.startingPoint.Y + columnNumber - 1
		emptyPos := make([]int, 0)
		for x := input.startingPoint.X; x < input.startingPoint.X+input.xSize; x++ {
			if input.emptyChecker.IsEmpty(x, y) {
				emptyPos = append(emptyPos, x)
			}
		}

		rowInput = append(rowInput, RowInput{
			criteria: model.SeatBase{
				Pos:      model.Pos{X: input.startingPoint.X, Y: y},
				SeatType: "A석",
			},
			initialNumber: 1,
			amount:        input.xSize - len(emptyPos),
			emptyPos:      emptyPos,
			columnNumber:  columnNumber,
		})
	}
	return newHorizontalBlock(rowInput...)
}

func newHorizontalBlock(row ...RowInput) *Block {
	block := new(Block)
	for _, v := range row {
		block.row = append(block.row, newHorizontalRow(v))
	}
	return block
}

func main() {
	hall := emptychecker.VerticalHallway(13, 14)
	rect1 := emptychecker.Rectangle(model.Pos{X: 3, Y: 4}, model.Pos{X: 6, Y: 5})
	rect2 := emptychecker.Rectangle(model.Pos{X: 4, Y: 15}, model.Pos{X: 7, Y: 15})
	specific := emptychecker.Position(model.Pos{X: 3, Y: 6})
	integrated := emptychecker.Integrated(hall, rect1, rect2, specific)

	blockInput := BlockInput{
		startingPoint: model.SeatBase{
			Pos:      model.Pos{X: 3, Y: 4},
			SeatType: "A석",
		},
		xSize:        21,
		ySize:        12,
		emptyChecker: integrated,
	}

	block := newBlock(blockInput)
	fmt.Println(block)
}
