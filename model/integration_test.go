package model_test

import (
	"testing"

	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/emptychecker"
	"github.com/simp7/seatCreator/model/group"
	"github.com/simp7/seatCreator/model/nameformatter"
	"github.com/simp7/seatCreator/model/pos"
	"github.com/stretchr/testify/assert"
)

func ArtriumSmall() model.Group {
	hall := emptychecker.VerticalHallway(13, 14)
	rect1 := emptychecker.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 6, Y: 5})
	rect2 := emptychecker.Rectangle(pos.Absolute{X: 4, Y: 15}, pos.Absolute{X: 7, Y: 15})
	specific := emptychecker.Position(pos.Absolute{X: 3, Y: 6})
	integrated := emptychecker.Integrated(hall, rect1, rect2, specific)
	nameGenerator := nameformatter.Standard()

	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         21,
		YSize:         12,
		EmptyChecker:  integrated,
		NameGenerator: nameGenerator,
	}

	return group.HorizontalBlock(blockInput)
}
func TestMain(t *testing.T) {
	const result = `{
	x: 7,
	y: 4,
	seat_type: "A석",
	short_name: "A01",
	name: "A석 1열 01"
}, {
	x: 8,
	y: 4,
	seat_type: "A석",
	short_name: "A02",
	name: "A석 1열 02"
}, {
	x: 9,
	y: 4,
	seat_type: "A석",
	short_name: "A03",
	name: "A석 1열 03"
}, {
	x: 10,
	y: 4,
	seat_type: "A석",
	short_name: "A04",
	name: "A석 1열 04"
}, {
	x: 11,
	y: 4,
	seat_type: "A석",
	short_name: "A05",
	name: "A석 1열 05"
}, {
	x: 12,
	y: 4,
	seat_type: "A석",
	short_name: "A06",
	name: "A석 1열 06"
}, {
	x: 15,
	y: 4,
	seat_type: "A석",
	short_name: "A07",
	name: "A석 1열 07"
}, {
	x: 16,
	y: 4,
	seat_type: "A석",
	short_name: "A08",
	name: "A석 1열 08"
}, {
	x: 17,
	y: 4,
	seat_type: "A석",
	short_name: "A09",
	name: "A석 1열 09"
}, {
	x: 18,
	y: 4,
	seat_type: "A석",
	short_name: "A10",
	name: "A석 1열 10"
}, {
	x: 19,
	y: 4,
	seat_type: "A석",
	short_name: "A11",
	name: "A석 1열 11"
}, {
	x: 20,
	y: 4,
	seat_type: "A석",
	short_name: "A12",
	name: "A석 1열 12"
}, {
	x: 21,
	y: 4,
	seat_type: "A석",
	short_name: "A13",
	name: "A석 1열 13"
}, {
	x: 22,
	y: 4,
	seat_type: "A석",
	short_name: "A14",
	name: "A석 1열 14"
}, {
	x: 23,
	y: 4,
	seat_type: "A석",
	short_name: "A15",
	name: "A석 1열 15"
}, {
	x: 7,
	y: 5,
	seat_type: "A석",
	short_name: "B01",
	name: "A석 2열 01"
}, {
	x: 8,
	y: 5,
	seat_type: "A석",
	short_name: "B02",
	name: "A석 2열 02"
}, {
	x: 9,
	y: 5,
	seat_type: "A석",
	short_name: "B03",
	name: "A석 2열 03"
}, {
	x: 10,
	y: 5,
	seat_type: "A석",
	short_name: "B04",
	name: "A석 2열 04"
}, {
	x: 11,
	y: 5,
	seat_type: "A석",
	short_name: "B05",
	name: "A석 2열 05"
}, {
	x: 12,
	y: 5,
	seat_type: "A석",
	short_name: "B06",
	name: "A석 2열 06"
}, {
	x: 15,
	y: 5,
	seat_type: "A석",
	short_name: "B07",
	name: "A석 2열 07"
}, {
	x: 16,
	y: 5,
	seat_type: "A석",
	short_name: "B08",
	name: "A석 2열 08"
}, {
	x: 17,
	y: 5,
	seat_type: "A석",
	short_name: "B09",
	name: "A석 2열 09"
}, {
	x: 18,
	y: 5,
	seat_type: "A석",
	short_name: "B10",
	name: "A석 2열 10"
}, {
	x: 19,
	y: 5,
	seat_type: "A석",
	short_name: "B11",
	name: "A석 2열 11"
}, {
	x: 20,
	y: 5,
	seat_type: "A석",
	short_name: "B12",
	name: "A석 2열 12"
}, {
	x: 21,
	y: 5,
	seat_type: "A석",
	short_name: "B13",
	name: "A석 2열 13"
}, {
	x: 22,
	y: 5,
	seat_type: "A석",
	short_name: "B14",
	name: "A석 2열 14"
}, {
	x: 23,
	y: 5,
	seat_type: "A석",
	short_name: "B15",
	name: "A석 2열 15"
}, {
	x: 4,
	y: 6,
	seat_type: "A석",
	short_name: "C01",
	name: "A석 3열 01"
}, {
	x: 5,
	y: 6,
	seat_type: "A석",
	short_name: "C02",
	name: "A석 3열 02"
}, {
	x: 6,
	y: 6,
	seat_type: "A석",
	short_name: "C03",
	name: "A석 3열 03"
}, {
	x: 7,
	y: 6,
	seat_type: "A석",
	short_name: "C04",
	name: "A석 3열 04"
}, {
	x: 8,
	y: 6,
	seat_type: "A석",
	short_name: "C05",
	name: "A석 3열 05"
}, {
	x: 9,
	y: 6,
	seat_type: "A석",
	short_name: "C06",
	name: "A석 3열 06"
}, {
	x: 10,
	y: 6,
	seat_type: "A석",
	short_name: "C07",
	name: "A석 3열 07"
}, {
	x: 11,
	y: 6,
	seat_type: "A석",
	short_name: "C08",
	name: "A석 3열 08"
}, {
	x: 12,
	y: 6,
	seat_type: "A석",
	short_name: "C09",
	name: "A석 3열 09"
}, {
	x: 15,
	y: 6,
	seat_type: "A석",
	short_name: "C10",
	name: "A석 3열 10"
}, {
	x: 16,
	y: 6,
	seat_type: "A석",
	short_name: "C11",
	name: "A석 3열 11"
}, {
	x: 17,
	y: 6,
	seat_type: "A석",
	short_name: "C12",
	name: "A석 3열 12"
}, {
	x: 18,
	y: 6,
	seat_type: "A석",
	short_name: "C13",
	name: "A석 3열 13"
}, {
	x: 19,
	y: 6,
	seat_type: "A석",
	short_name: "C14",
	name: "A석 3열 14"
}, {
	x: 20,
	y: 6,
	seat_type: "A석",
	short_name: "C15",
	name: "A석 3열 15"
}, {
	x: 21,
	y: 6,
	seat_type: "A석",
	short_name: "C16",
	name: "A석 3열 16"
}, {
	x: 22,
	y: 6,
	seat_type: "A석",
	short_name: "C17",
	name: "A석 3열 17"
}, {
	x: 23,
	y: 6,
	seat_type: "A석",
	short_name: "C18",
	name: "A석 3열 18"
}, {
	x: 3,
	y: 7,
	seat_type: "A석",
	short_name: "D01",
	name: "A석 4열 01"
}, {
	x: 4,
	y: 7,
	seat_type: "A석",
	short_name: "D02",
	name: "A석 4열 02"
}, {
	x: 5,
	y: 7,
	seat_type: "A석",
	short_name: "D03",
	name: "A석 4열 03"
}, {
	x: 6,
	y: 7,
	seat_type: "A석",
	short_name: "D04",
	name: "A석 4열 04"
}, {
	x: 7,
	y: 7,
	seat_type: "A석",
	short_name: "D05",
	name: "A석 4열 05"
}, {
	x: 8,
	y: 7,
	seat_type: "A석",
	short_name: "D06",
	name: "A석 4열 06"
}, {
	x: 9,
	y: 7,
	seat_type: "A석",
	short_name: "D07",
	name: "A석 4열 07"
}, {
	x: 10,
	y: 7,
	seat_type: "A석",
	short_name: "D08",
	name: "A석 4열 08"
}, {
	x: 11,
	y: 7,
	seat_type: "A석",
	short_name: "D09",
	name: "A석 4열 09"
}, {
	x: 12,
	y: 7,
	seat_type: "A석",
	short_name: "D10",
	name: "A석 4열 10"
}, {
	x: 15,
	y: 7,
	seat_type: "A석",
	short_name: "D11",
	name: "A석 4열 11"
}, {
	x: 16,
	y: 7,
	seat_type: "A석",
	short_name: "D12",
	name: "A석 4열 12"
}, {
	x: 17,
	y: 7,
	seat_type: "A석",
	short_name: "D13",
	name: "A석 4열 13"
}, {
	x: 18,
	y: 7,
	seat_type: "A석",
	short_name: "D14",
	name: "A석 4열 14"
}, {
	x: 19,
	y: 7,
	seat_type: "A석",
	short_name: "D15",
	name: "A석 4열 15"
}, {
	x: 20,
	y: 7,
	seat_type: "A석",
	short_name: "D16",
	name: "A석 4열 16"
}, {
	x: 21,
	y: 7,
	seat_type: "A석",
	short_name: "D17",
	name: "A석 4열 17"
}, {
	x: 22,
	y: 7,
	seat_type: "A석",
	short_name: "D18",
	name: "A석 4열 18"
}, {
	x: 23,
	y: 7,
	seat_type: "A석",
	short_name: "D19",
	name: "A석 4열 19"
}, {
	x: 3,
	y: 8,
	seat_type: "A석",
	short_name: "E01",
	name: "A석 5열 01"
}, {
	x: 4,
	y: 8,
	seat_type: "A석",
	short_name: "E02",
	name: "A석 5열 02"
}, {
	x: 5,
	y: 8,
	seat_type: "A석",
	short_name: "E03",
	name: "A석 5열 03"
}, {
	x: 6,
	y: 8,
	seat_type: "A석",
	short_name: "E04",
	name: "A석 5열 04"
}, {
	x: 7,
	y: 8,
	seat_type: "A석",
	short_name: "E05",
	name: "A석 5열 05"
}, {
	x: 8,
	y: 8,
	seat_type: "A석",
	short_name: "E06",
	name: "A석 5열 06"
}, {
	x: 9,
	y: 8,
	seat_type: "A석",
	short_name: "E07",
	name: "A석 5열 07"
}, {
	x: 10,
	y: 8,
	seat_type: "A석",
	short_name: "E08",
	name: "A석 5열 08"
}, {
	x: 11,
	y: 8,
	seat_type: "A석",
	short_name: "E09",
	name: "A석 5열 09"
}, {
	x: 12,
	y: 8,
	seat_type: "A석",
	short_name: "E10",
	name: "A석 5열 10"
}, {
	x: 15,
	y: 8,
	seat_type: "A석",
	short_name: "E11",
	name: "A석 5열 11"
}, {
	x: 16,
	y: 8,
	seat_type: "A석",
	short_name: "E12",
	name: "A석 5열 12"
}, {
	x: 17,
	y: 8,
	seat_type: "A석",
	short_name: "E13",
	name: "A석 5열 13"
}, {
	x: 18,
	y: 8,
	seat_type: "A석",
	short_name: "E14",
	name: "A석 5열 14"
}, {
	x: 19,
	y: 8,
	seat_type: "A석",
	short_name: "E15",
	name: "A석 5열 15"
}, {
	x: 20,
	y: 8,
	seat_type: "A석",
	short_name: "E16",
	name: "A석 5열 16"
}, {
	x: 21,
	y: 8,
	seat_type: "A석",
	short_name: "E17",
	name: "A석 5열 17"
}, {
	x: 22,
	y: 8,
	seat_type: "A석",
	short_name: "E18",
	name: "A석 5열 18"
}, {
	x: 23,
	y: 8,
	seat_type: "A석",
	short_name: "E19",
	name: "A석 5열 19"
}, {
	x: 3,
	y: 9,
	seat_type: "A석",
	short_name: "F01",
	name: "A석 6열 01"
}, {
	x: 4,
	y: 9,
	seat_type: "A석",
	short_name: "F02",
	name: "A석 6열 02"
}, {
	x: 5,
	y: 9,
	seat_type: "A석",
	short_name: "F03",
	name: "A석 6열 03"
}, {
	x: 6,
	y: 9,
	seat_type: "A석",
	short_name: "F04",
	name: "A석 6열 04"
}, {
	x: 7,
	y: 9,
	seat_type: "A석",
	short_name: "F05",
	name: "A석 6열 05"
}, {
	x: 8,
	y: 9,
	seat_type: "A석",
	short_name: "F06",
	name: "A석 6열 06"
}, {
	x: 9,
	y: 9,
	seat_type: "A석",
	short_name: "F07",
	name: "A석 6열 07"
}, {
	x: 10,
	y: 9,
	seat_type: "A석",
	short_name: "F08",
	name: "A석 6열 08"
}, {
	x: 11,
	y: 9,
	seat_type: "A석",
	short_name: "F09",
	name: "A석 6열 09"
}, {
	x: 12,
	y: 9,
	seat_type: "A석",
	short_name: "F10",
	name: "A석 6열 10"
}, {
	x: 15,
	y: 9,
	seat_type: "A석",
	short_name: "F11",
	name: "A석 6열 11"
}, {
	x: 16,
	y: 9,
	seat_type: "A석",
	short_name: "F12",
	name: "A석 6열 12"
}, {
	x: 17,
	y: 9,
	seat_type: "A석",
	short_name: "F13",
	name: "A석 6열 13"
}, {
	x: 18,
	y: 9,
	seat_type: "A석",
	short_name: "F14",
	name: "A석 6열 14"
}, {
	x: 19,
	y: 9,
	seat_type: "A석",
	short_name: "F15",
	name: "A석 6열 15"
}, {
	x: 20,
	y: 9,
	seat_type: "A석",
	short_name: "F16",
	name: "A석 6열 16"
}, {
	x: 21,
	y: 9,
	seat_type: "A석",
	short_name: "F17",
	name: "A석 6열 17"
}, {
	x: 22,
	y: 9,
	seat_type: "A석",
	short_name: "F18",
	name: "A석 6열 18"
}, {
	x: 23,
	y: 9,
	seat_type: "A석",
	short_name: "F19",
	name: "A석 6열 19"
}, {
	x: 3,
	y: 10,
	seat_type: "A석",
	short_name: "G01",
	name: "A석 7열 01"
}, {
	x: 4,
	y: 10,
	seat_type: "A석",
	short_name: "G02",
	name: "A석 7열 02"
}, {
	x: 5,
	y: 10,
	seat_type: "A석",
	short_name: "G03",
	name: "A석 7열 03"
}, {
	x: 6,
	y: 10,
	seat_type: "A석",
	short_name: "G04",
	name: "A석 7열 04"
}, {
	x: 7,
	y: 10,
	seat_type: "A석",
	short_name: "G05",
	name: "A석 7열 05"
}, {
	x: 8,
	y: 10,
	seat_type: "A석",
	short_name: "G06",
	name: "A석 7열 06"
}, {
	x: 9,
	y: 10,
	seat_type: "A석",
	short_name: "G07",
	name: "A석 7열 07"
}, {
	x: 10,
	y: 10,
	seat_type: "A석",
	short_name: "G08",
	name: "A석 7열 08"
}, {
	x: 11,
	y: 10,
	seat_type: "A석",
	short_name: "G09",
	name: "A석 7열 09"
}, {
	x: 12,
	y: 10,
	seat_type: "A석",
	short_name: "G10",
	name: "A석 7열 10"
}, {
	x: 15,
	y: 10,
	seat_type: "A석",
	short_name: "G11",
	name: "A석 7열 11"
}, {
	x: 16,
	y: 10,
	seat_type: "A석",
	short_name: "G12",
	name: "A석 7열 12"
}, {
	x: 17,
	y: 10,
	seat_type: "A석",
	short_name: "G13",
	name: "A석 7열 13"
}, {
	x: 18,
	y: 10,
	seat_type: "A석",
	short_name: "G14",
	name: "A석 7열 14"
}, {
	x: 19,
	y: 10,
	seat_type: "A석",
	short_name: "G15",
	name: "A석 7열 15"
}, {
	x: 20,
	y: 10,
	seat_type: "A석",
	short_name: "G16",
	name: "A석 7열 16"
}, {
	x: 21,
	y: 10,
	seat_type: "A석",
	short_name: "G17",
	name: "A석 7열 17"
}, {
	x: 22,
	y: 10,
	seat_type: "A석",
	short_name: "G18",
	name: "A석 7열 18"
}, {
	x: 23,
	y: 10,
	seat_type: "A석",
	short_name: "G19",
	name: "A석 7열 19"
}, {
	x: 3,
	y: 11,
	seat_type: "A석",
	short_name: "H01",
	name: "A석 8열 01"
}, {
	x: 4,
	y: 11,
	seat_type: "A석",
	short_name: "H02",
	name: "A석 8열 02"
}, {
	x: 5,
	y: 11,
	seat_type: "A석",
	short_name: "H03",
	name: "A석 8열 03"
}, {
	x: 6,
	y: 11,
	seat_type: "A석",
	short_name: "H04",
	name: "A석 8열 04"
}, {
	x: 7,
	y: 11,
	seat_type: "A석",
	short_name: "H05",
	name: "A석 8열 05"
}, {
	x: 8,
	y: 11,
	seat_type: "A석",
	short_name: "H06",
	name: "A석 8열 06"
}, {
	x: 9,
	y: 11,
	seat_type: "A석",
	short_name: "H07",
	name: "A석 8열 07"
}, {
	x: 10,
	y: 11,
	seat_type: "A석",
	short_name: "H08",
	name: "A석 8열 08"
}, {
	x: 11,
	y: 11,
	seat_type: "A석",
	short_name: "H09",
	name: "A석 8열 09"
}, {
	x: 12,
	y: 11,
	seat_type: "A석",
	short_name: "H10",
	name: "A석 8열 10"
}, {
	x: 15,
	y: 11,
	seat_type: "A석",
	short_name: "H11",
	name: "A석 8열 11"
}, {
	x: 16,
	y: 11,
	seat_type: "A석",
	short_name: "H12",
	name: "A석 8열 12"
}, {
	x: 17,
	y: 11,
	seat_type: "A석",
	short_name: "H13",
	name: "A석 8열 13"
}, {
	x: 18,
	y: 11,
	seat_type: "A석",
	short_name: "H14",
	name: "A석 8열 14"
}, {
	x: 19,
	y: 11,
	seat_type: "A석",
	short_name: "H15",
	name: "A석 8열 15"
}, {
	x: 20,
	y: 11,
	seat_type: "A석",
	short_name: "H16",
	name: "A석 8열 16"
}, {
	x: 21,
	y: 11,
	seat_type: "A석",
	short_name: "H17",
	name: "A석 8열 17"
}, {
	x: 22,
	y: 11,
	seat_type: "A석",
	short_name: "H18",
	name: "A석 8열 18"
}, {
	x: 23,
	y: 11,
	seat_type: "A석",
	short_name: "H19",
	name: "A석 8열 19"
}, {
	x: 3,
	y: 12,
	seat_type: "A석",
	short_name: "I01",
	name: "A석 9열 01"
}, {
	x: 4,
	y: 12,
	seat_type: "A석",
	short_name: "I02",
	name: "A석 9열 02"
}, {
	x: 5,
	y: 12,
	seat_type: "A석",
	short_name: "I03",
	name: "A석 9열 03"
}, {
	x: 6,
	y: 12,
	seat_type: "A석",
	short_name: "I04",
	name: "A석 9열 04"
}, {
	x: 7,
	y: 12,
	seat_type: "A석",
	short_name: "I05",
	name: "A석 9열 05"
}, {
	x: 8,
	y: 12,
	seat_type: "A석",
	short_name: "I06",
	name: "A석 9열 06"
}, {
	x: 9,
	y: 12,
	seat_type: "A석",
	short_name: "I07",
	name: "A석 9열 07"
}, {
	x: 10,
	y: 12,
	seat_type: "A석",
	short_name: "I08",
	name: "A석 9열 08"
}, {
	x: 11,
	y: 12,
	seat_type: "A석",
	short_name: "I09",
	name: "A석 9열 09"
}, {
	x: 12,
	y: 12,
	seat_type: "A석",
	short_name: "I10",
	name: "A석 9열 10"
}, {
	x: 15,
	y: 12,
	seat_type: "A석",
	short_name: "I11",
	name: "A석 9열 11"
}, {
	x: 16,
	y: 12,
	seat_type: "A석",
	short_name: "I12",
	name: "A석 9열 12"
}, {
	x: 17,
	y: 12,
	seat_type: "A석",
	short_name: "I13",
	name: "A석 9열 13"
}, {
	x: 18,
	y: 12,
	seat_type: "A석",
	short_name: "I14",
	name: "A석 9열 14"
}, {
	x: 19,
	y: 12,
	seat_type: "A석",
	short_name: "I15",
	name: "A석 9열 15"
}, {
	x: 20,
	y: 12,
	seat_type: "A석",
	short_name: "I16",
	name: "A석 9열 16"
}, {
	x: 21,
	y: 12,
	seat_type: "A석",
	short_name: "I17",
	name: "A석 9열 17"
}, {
	x: 22,
	y: 12,
	seat_type: "A석",
	short_name: "I18",
	name: "A석 9열 18"
}, {
	x: 23,
	y: 12,
	seat_type: "A석",
	short_name: "I19",
	name: "A석 9열 19"
}, {
	x: 3,
	y: 13,
	seat_type: "A석",
	short_name: "J01",
	name: "A석 10열 01"
}, {
	x: 4,
	y: 13,
	seat_type: "A석",
	short_name: "J02",
	name: "A석 10열 02"
}, {
	x: 5,
	y: 13,
	seat_type: "A석",
	short_name: "J03",
	name: "A석 10열 03"
}, {
	x: 6,
	y: 13,
	seat_type: "A석",
	short_name: "J04",
	name: "A석 10열 04"
}, {
	x: 7,
	y: 13,
	seat_type: "A석",
	short_name: "J05",
	name: "A석 10열 05"
}, {
	x: 8,
	y: 13,
	seat_type: "A석",
	short_name: "J06",
	name: "A석 10열 06"
}, {
	x: 9,
	y: 13,
	seat_type: "A석",
	short_name: "J07",
	name: "A석 10열 07"
}, {
	x: 10,
	y: 13,
	seat_type: "A석",
	short_name: "J08",
	name: "A석 10열 08"
}, {
	x: 11,
	y: 13,
	seat_type: "A석",
	short_name: "J09",
	name: "A석 10열 09"
}, {
	x: 12,
	y: 13,
	seat_type: "A석",
	short_name: "J10",
	name: "A석 10열 10"
}, {
	x: 15,
	y: 13,
	seat_type: "A석",
	short_name: "J11",
	name: "A석 10열 11"
}, {
	x: 16,
	y: 13,
	seat_type: "A석",
	short_name: "J12",
	name: "A석 10열 12"
}, {
	x: 17,
	y: 13,
	seat_type: "A석",
	short_name: "J13",
	name: "A석 10열 13"
}, {
	x: 18,
	y: 13,
	seat_type: "A석",
	short_name: "J14",
	name: "A석 10열 14"
}, {
	x: 19,
	y: 13,
	seat_type: "A석",
	short_name: "J15",
	name: "A석 10열 15"
}, {
	x: 20,
	y: 13,
	seat_type: "A석",
	short_name: "J16",
	name: "A석 10열 16"
}, {
	x: 21,
	y: 13,
	seat_type: "A석",
	short_name: "J17",
	name: "A석 10열 17"
}, {
	x: 22,
	y: 13,
	seat_type: "A석",
	short_name: "J18",
	name: "A석 10열 18"
}, {
	x: 23,
	y: 13,
	seat_type: "A석",
	short_name: "J19",
	name: "A석 10열 19"
}, {
	x: 3,
	y: 14,
	seat_type: "A석",
	short_name: "K01",
	name: "A석 11열 01"
}, {
	x: 4,
	y: 14,
	seat_type: "A석",
	short_name: "K02",
	name: "A석 11열 02"
}, {
	x: 5,
	y: 14,
	seat_type: "A석",
	short_name: "K03",
	name: "A석 11열 03"
}, {
	x: 6,
	y: 14,
	seat_type: "A석",
	short_name: "K04",
	name: "A석 11열 04"
}, {
	x: 7,
	y: 14,
	seat_type: "A석",
	short_name: "K05",
	name: "A석 11열 05"
}, {
	x: 8,
	y: 14,
	seat_type: "A석",
	short_name: "K06",
	name: "A석 11열 06"
}, {
	x: 9,
	y: 14,
	seat_type: "A석",
	short_name: "K07",
	name: "A석 11열 07"
}, {
	x: 10,
	y: 14,
	seat_type: "A석",
	short_name: "K08",
	name: "A석 11열 08"
}, {
	x: 11,
	y: 14,
	seat_type: "A석",
	short_name: "K09",
	name: "A석 11열 09"
}, {
	x: 12,
	y: 14,
	seat_type: "A석",
	short_name: "K10",
	name: "A석 11열 10"
}, {
	x: 15,
	y: 14,
	seat_type: "A석",
	short_name: "K11",
	name: "A석 11열 11"
}, {
	x: 16,
	y: 14,
	seat_type: "A석",
	short_name: "K12",
	name: "A석 11열 12"
}, {
	x: 17,
	y: 14,
	seat_type: "A석",
	short_name: "K13",
	name: "A석 11열 13"
}, {
	x: 18,
	y: 14,
	seat_type: "A석",
	short_name: "K14",
	name: "A석 11열 14"
}, {
	x: 19,
	y: 14,
	seat_type: "A석",
	short_name: "K15",
	name: "A석 11열 15"
}, {
	x: 20,
	y: 14,
	seat_type: "A석",
	short_name: "K16",
	name: "A석 11열 16"
}, {
	x: 21,
	y: 14,
	seat_type: "A석",
	short_name: "K17",
	name: "A석 11열 17"
}, {
	x: 22,
	y: 14,
	seat_type: "A석",
	short_name: "K18",
	name: "A석 11열 18"
}, {
	x: 23,
	y: 14,
	seat_type: "A석",
	short_name: "K19",
	name: "A석 11열 19"
}, {
	x: 3,
	y: 15,
	seat_type: "A석",
	short_name: "L01",
	name: "A석 12열 01"
}, {
	x: 8,
	y: 15,
	seat_type: "A석",
	short_name: "L02",
	name: "A석 12열 02"
}, {
	x: 9,
	y: 15,
	seat_type: "A석",
	short_name: "L03",
	name: "A석 12열 03"
}, {
	x: 10,
	y: 15,
	seat_type: "A석",
	short_name: "L04",
	name: "A석 12열 04"
}, {
	x: 11,
	y: 15,
	seat_type: "A석",
	short_name: "L05",
	name: "A석 12열 05"
}, {
	x: 12,
	y: 15,
	seat_type: "A석",
	short_name: "L06",
	name: "A석 12열 06"
}, {
	x: 15,
	y: 15,
	seat_type: "A석",
	short_name: "L07",
	name: "A석 12열 07"
}, {
	x: 16,
	y: 15,
	seat_type: "A석",
	short_name: "L08",
	name: "A석 12열 08"
}, {
	x: 17,
	y: 15,
	seat_type: "A석",
	short_name: "L09",
	name: "A석 12열 09"
}, {
	x: 18,
	y: 15,
	seat_type: "A석",
	short_name: "L10",
	name: "A석 12열 10"
}, {
	x: 19,
	y: 15,
	seat_type: "A석",
	short_name: "L11",
	name: "A석 12열 11"
}, {
	x: 20,
	y: 15,
	seat_type: "A석",
	short_name: "L12",
	name: "A석 12열 12"
}, {
	x: 21,
	y: 15,
	seat_type: "A석",
	short_name: "L13",
	name: "A석 12열 13"
}, {
	x: 22,
	y: 15,
	seat_type: "A석",
	short_name: "L14",
	name: "A석 12열 14"
}, {
	x: 23,
	y: 15,
	seat_type: "A석",
	short_name: "L15",
	name: "A석 12열 15"
}`
	assert.Equal(t, result, ArtriumSmall().String(), "Should be Equal")
}
