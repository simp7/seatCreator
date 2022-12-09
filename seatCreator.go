package main

import (
	"fmt"

	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/emptychecker"
	"github.com/simp7/seatCreator/model/group"
	"github.com/simp7/seatCreator/model/nameformatter"
	"github.com/simp7/seatCreator/model/pos"
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

func OperaHouse() model.Group {
	return ArtriumSmall()
}

func ConcertHall() model.Group {
	vHall := emptychecker.VerticalHallway(14, 35)
	hHall := emptychecker.HorizontalHallway(16, 17, 18, 19)
	integrated := emptychecker.Integrated(vHall, hHall)
	nameGenerator := nameformatter.Standard()

	vipHall := emptychecker.VerticalHallway(3, 5, 7, 9, 11, 13, 36, 38, 40, 42, 44, 46)
	empty := emptychecker.Rectangle(pos.Absolute{X: 14, Y: 26}, pos.Absolute{X: 35, Y: 26})
	integrated2 := emptychecker.Integrated(vipHall, empty)

	base := model.NewSeatBase(2, 8, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         46,
		YSize:         18,
		EmptyChecker:  integrated,
		NameGenerator: nameGenerator,
	}

	base = model.NewSeatBase(15, 26, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 13, 15),
		XSize:         20,
		YSize:         2,
		EmptyChecker:  integrated,
		NameGenerator: nameGenerator,
	}

	base = model.NewSeatBase(2, 26, "VIP석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 15),
		XSize:         46,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameGenerator: nameGenerator,
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	return group.Mixed(block1, block2, block3)
}

func main() {
	fmt.Println(ArtriumSmall())
	fmt.Println(ConcertHall())
}
