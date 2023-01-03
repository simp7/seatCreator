package model

import (
	"fmt"
	"sync"

	"github.com/simp7/seatCreator/model/pos"
)

var (
	once      sync.Once
	colorList []string
	colorMap  map[string]string
)

type SeatBase struct {
	pos.Absolute
	SeatType string
}

func NewSeatBase(x int, y int, seatType string) SeatBase {
	return SeatBase{Absolute: pos.Absolute{X: x, Y: y}, SeatType: seatType}
}

type Seat struct {
	SeatBase
	pos.Relative
}

func (s Seat) String(name NameFormatter) string {
	return fmt.Sprintf("{\n\tx: %d,\n\ty: %d,\n\tseat_type: \"%s\",\n\tshort_name: \"%s\",\n\tname: \"%s\"\n}", s.X, s.Y, s.SeatType, name.Short(s), name.Long(s))
}

func getColor(seatType string) string {
	once.Do(func() {
		colorList = []string{"#C00", "#D80", "#DD0", "#8D0", "#08D", "#00D", "#808", "#888"}
		colorMap = make(map[string]string)
	})
	value, ok := colorMap[seatType]
	if !ok {
		value = colorList[len(colorMap)%8]
		colorMap[seatType] = value
	}
	return value
}

func (s Seat) Html(name NameFormatter) string {
	top := s.Y * 36
	left := s.X * 36
	return fmt.Sprintf(`<div style="position: absolute; top: %d; left: %d; display: flex; width:36px; height:36px; align-items: center; justify-content: center">
	<div style="display: flex; width: 30px; height: 30px; background-color:%s; border-radius:5px; align-items: center; justify-content: center">
		<text style="font-size:13; color: white">%s</text>
	</div>
</div>`, top, left, getColor(s.SeatType), name.Short(s))
}

func NewSeat(base SeatBase, index int, row int) Seat {
	result := Seat{
		SeatBase: base,
		Relative: pos.Relative{
			Index: index,
			Row:   row,
		},
	}
	return result
}
