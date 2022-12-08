package main

import (
	"fmt"
	"strings"
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
	base         SeatBase
	xSize        int
	ySize        int
	emptyChecker EmptyChecker
}

func newHorizontalBlock(row ...RowInput) *Block {
	block := new(Block)
	for _, v := range row {
		block.row = append(block.row, newHorizontalRow(v))
	}
	return block
}

func main() {
	rowInput := make([]RowInput, 0)

	rowInput = append(rowInput, RowInput{
		criteria: SeatBase{
			x:        7,
			y:        4,
			seatType: "A석",
		},
		initialNumber: 1,
		amount:        15,
		emptyPos:      []int{13, 14},
		columnNumber:  1,
	}, RowInput{
		criteria: SeatBase{
			x:        7,
			y:        5,
			seatType: "A석",
		},
		initialNumber: 1,
		amount:        15,
		emptyPos:      []int{13, 14},
		columnNumber:  2,
	}, RowInput{
		criteria: SeatBase{
			x:        4,
			y:        6,
			seatType: "A석",
		},
		initialNumber: 1,
		amount:        18,
		emptyPos:      []int{13, 14},
		columnNumber:  3,
	})

	for columnNumber := 4; columnNumber < 11; columnNumber++ {
		rowInput = append(rowInput, RowInput{
			criteria: SeatBase{
				x:        3,
				y:        columnNumber + 3,
				seatType: "A석",
			},
			initialNumber: 1,
			amount:        19,
			emptyPos:      []int{13, 14},
			columnNumber:  columnNumber,
		})
	}

	rowInput = append(rowInput, RowInput{
		criteria: SeatBase{
			x:        3,
			y:        16,
			seatType: "A석",
		},
		initialNumber: 1,
		amount:        15,
		emptyPos:      []int{4, 5, 6, 7, 13, 14},
		columnNumber:  12,
	})

	block := newHorizontalBlock(rowInput...)
	fmt.Println(block)
}
