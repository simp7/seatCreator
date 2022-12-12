package nameformatter

import (
	"fmt"

	"github.com/simp7/seatCreator/model"
)

type floorFormatter struct {
	model.NameFormatter
	prefix string
}

func (f floorFormatter) Long(seat model.Seat) string {
	return fmt.Sprintf("%s %s", f.prefix, f.NameFormatter.Long(seat))
}

func Floor(formatter model.NameFormatter, floor int) model.NameFormatter {
	prefix := ""
	if floor > 0 {
		prefix = fmt.Sprintf("%d층", floor)
	}
	if floor < 0 {
		prefix = fmt.Sprintf("지하 %d층", floor)
	}

	return floorFormatter{NameFormatter: formatter, prefix: prefix}
}
