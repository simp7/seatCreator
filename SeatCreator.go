package main

import (
	"fmt"
	"strings"

	"github.com/simp7/seatCreator/model"
	emptychecker "github.com/simp7/seatCreator/model/emptyChecker"
)

func twoDigit(n int) string {
	if n < 10 {
		return fmt.Sprintf("0%d", n)
	}
	return fmt.Sprint(n)
}

type SeatBase struct {
	x        int
	y        int
	seatType string
}

type Seat struct {
	X         int
	Y         int
	SeatType  string `json:"seat_type"`
	ShortName string `json:"short_type"`
	Name      string
}

func (s Seat) String() string {
	return fmt.Sprintf("{\n\tx: %d,\n\ty: %d,\n\tseat_type: \"%s\",\n\tshort_name: \"%s\",\n\tname: \"%s\"\n}", s.X, s.Y, s.SeatType, s.ShortName, s.Name)
}

func newSeat(base SeatBase, generateShortName func(_ SeatBase) string, generateName func(_ SeatBase) string) Seat {
	var result Seat = Seat{
		X:         base.x,
		Y:         base.y,
		SeatType:  base.seatType,
		ShortName: generateShortName(base),
		Name:      generateName(base),
	}
	return result
}

type Row struct {
	seats []Seat
}

func (r Row) String() string {

	var result = make([]string, 0)

	for _, v := range r.seats {
		result = append(result, v.String())
	}
	return strings.Join(result, ", ")
}

type RowInput struct {
	criteria      SeatBase
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
	y := input.criteria.y
	seatType := input.criteria.seatType
	seats := make([]Seat, 0)

	posList := getPositionList(input.criteria.x, input.amount, input.emptyPos)

	for i, x := range posList {
		generateShortName := func(target SeatBase) string {
			return fmt.Sprintf("%c%s", rune(64+input.columnNumber), twoDigit(i+1))
		}

		generateName := func(target SeatBase) string {
			return fmt.Sprintf("%s %d열 %s", target.seatType, input.columnNumber, twoDigit(i+1))
		}

		inputBase := SeatBase{
			x:        x,
			y:        y,
			seatType: seatType,
		}
		seats = append(seats, newSeat(inputBase, generateShortName, generateName))
	}

	row := new(Row)

	row.seats = seats
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
	startingPoint SeatBase
	xSize         int
	ySize         int
	emptyChecker  emptychecker.EmptyChecker
}

func newBlock(input BlockInput) *Block {
	rowInput := make([]RowInput, 0)

	for columnNumber := 1; columnNumber <= input.ySize; columnNumber++ {
		y := input.startingPoint.y + columnNumber - 1
		emptyPos := make([]int, 0)
		for x := input.startingPoint.x; x < input.startingPoint.x+input.xSize; x++ {
			if input.emptyChecker.IsEmpty(x, y) {
				emptyPos = append(emptyPos, x)
			}
		}

		rowInput = append(rowInput, RowInput{
			criteria: SeatBase{
				x:        input.startingPoint.x,
				y:        y,
				seatType: "A석",
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
	hall := emptychecker.NewVerticalHallwayChecker(13, 14)
	specific := emptychecker.NewPositionChecker(model.Pos{X: 3, Y: 4}, model.Pos{X: 4, Y: 4}, model.Pos{X: 5, Y: 4}, model.Pos{X: 6, Y: 4}, model.Pos{X: 3, Y: 5}, model.Pos{X: 4, Y: 5}, model.Pos{X: 5, Y: 5}, model.Pos{X: 6, Y: 5}, model.Pos{X: 3, Y: 6},
		model.Pos{X: 4, Y: 15}, model.Pos{X: 5, Y: 15}, model.Pos{X: 6, Y: 15}, model.Pos{X: 7, Y: 15})
	integrated := emptychecker.NewIntegratedChecker(hall, specific)

	blockInput := BlockInput{
		startingPoint: SeatBase{
			x:        3,
			y:        4,
			seatType: "A석",
		},
		xSize:        21,
		ySize:        12,
		emptyChecker: integrated,
	}

	block := newBlock(blockInput)
	fmt.Println(block)

	rect1 := emptychecker.NewRectangleChecker(model.Pos{X: 3, Y: 4}, model.Pos{X: 6, Y: 5})
	rect2 := emptychecker.NewRectangleChecker(model.Pos{X: 4, Y: 15}, model.Pos{X: 7, Y: 15})
	specific = emptychecker.NewPositionChecker(model.Pos{X: 3, Y: 6})
	integrated = emptychecker.NewIntegratedChecker(hall, rect1, rect2, specific)
	blockInput = BlockInput{
		startingPoint: SeatBase{
			x:        3,
			y:        4,
			seatType: "A석",
		},
		xSize:        21,
		ySize:        12,
		emptyChecker: integrated,
	}

	block = newBlock(blockInput)
	fmt.Println(block)
}
