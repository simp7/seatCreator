package nameformatter

import (
	"fmt"

	"github.com/simp7/seatCreator/model"
)

type prefixFormatter struct {
	prefix  string
	index   int
	hashMap map[model.Seat]int
}

func (p *prefixFormatter) Long(input model.Seat) string {
	return fmt.Sprintf("%s %s", input.SeatType, p.Short(input))
}

func (p *prefixFormatter) Short(input model.Seat) string {
	index, ok := p.hashMap[input]
	if !ok {
		p.hashMap[input] = p.index
		index = p.index
		p.index++
	}
	return fmt.Sprintf("%s%s", p.prefix, twoDigit(index))
}

func Prefix(prefix string) model.NameFormatter {
	hashMap := make(map[model.Seat]int)
	return &prefixFormatter{
		prefix:  prefix,
		index:   1,
		hashMap: hashMap,
	}
}
