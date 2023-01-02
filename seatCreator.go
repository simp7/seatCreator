package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/eraser"
	"github.com/simp7/seatCreator/model/group"
	"github.com/simp7/seatCreator/model/nameformatter"
	"github.com/simp7/seatCreator/model/pos"
)

func ArtriumSmall() model.Group {
	hall := eraser.VerticalHallway(13, 14)
	rect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 6, Y: 5})
	rect2 := eraser.Rectangle(pos.Absolute{X: 4, Y: 15}, pos.Absolute{X: 7, Y: 15})
	specific := eraser.Position(pos.Absolute{X: 3, Y: 6})
	integrated := eraser.Integrated(hall, rect1, rect2, specific)
	nameFormatter := nameformatter.Standard()

	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         21,
		YSize:         12,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	return group.HorizontalBlock(blockInput)
}

func OperaHouse() model.Group {
	return ArtriumSmall()
}

func ConcertHall1F() model.Group {
	vHall := eraser.VerticalHallway(14, 35)
	hHall := eraser.HorizontalHallway(16, 17, 18, 19)
	integrated := eraser.Integrated(vHall, hHall)
	nameFormatter := nameformatter.Floor(nameformatter.Standard(), 1)

	vipHall := eraser.VerticalHallway(3, 5, 7, 9, 11, 13, 36, 38, 40, 42, 44, 46)
	empty := eraser.Rectangle(pos.Absolute{X: 14, Y: 26}, pos.Absolute{X: 35, Y: 26})
	integrated2 := eraser.Integrated(vipHall, empty)

	base := model.NewSeatBase(2, 8, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         46,
		YSize:         18,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(15, 26, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 13, 15),
		XSize:         20,
		YSize:         2,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(2, 26, "VIP석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 15),
		XSize:         46,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	return group.Mixed(block1, block2, block3)
}

func ConcertHall2F() model.Group {
	top := func() model.Group {
		base := model.NewSeatBase(4, 2, "A석")
		seat := model.NewSeat(base, 1, 1)
		emptyL := eraser.Position(pos.Absolute{X: 4, Y: 3}, pos.Absolute{X: 4, Y: 4}, pos.Absolute{X: 5, Y: 4})
		emptyR := eraser.Position(pos.Absolute{X: 36, Y: 3}, pos.Absolute{X: 35, Y: 4}, pos.Absolute{X: 36, Y: 4})
		hallway := eraser.VerticalHallway(21)
		empty := eraser.Integrated(emptyL, emptyR, hallway)
		return group.HorizontalBlock(group.BlockInput{
			Criteria:      seat,
			XSize:         33,
			YSize:         3,
			NameFormatter: nameformatter.Floor(nameformatter.Prefix('T'), 2),
			EmptyChecker:  empty,
		})
	}

	left := func() model.Group {
		base := model.NewSeatBase(0, 5, "A석")
		seat := model.NewSeat(base, 1, 1)
		empty := eraser.Position(pos.Absolute{X: 0, Y: 5}, pos.Absolute{X: 0, Y: 6}, pos.Absolute{X: 1, Y: 5})
		return group.VerticalBlock(group.BlockInput{
			Criteria:      seat,
			XSize:         3,
			YSize:         21,
			NameFormatter: nameformatter.Floor(nameformatter.Prefix('L'), 2),
			EmptyChecker:  empty,
		})
	}

	right := func() model.Group {
		base := model.NewSeatBase(40, 5, "A석")
		seat := model.NewSeat(base, 1, 1)
		empty := eraser.Position(pos.Absolute{X: 41, Y: 5}, pos.Absolute{X: 42, Y: 5}, pos.Absolute{X: 42, Y: 6})
		return group.VerticalBlock(group.BlockInput{
			Criteria:      seat,
			XSize:         3,
			YSize:         21,
			NameFormatter: nameformatter.Floor(nameformatter.Prefix('R'), 2),
			EmptyChecker:  empty,
		})
	}

	bottom := func() model.Group {
		base := model.NewSeatBase(5, 24, "VIP석")
		seat := model.NewSeat(base, 1, 1)
		empty := eraser.VerticalHallway(14, 29)
		return group.HorizontalBlock(group.BlockInput{
			Criteria:      seat,
			XSize:         34,
			YSize:         7,
			NameFormatter: nameformatter.Floor(nameformatter.Standard(), 2),
			EmptyChecker:  empty,
		})
	}

	hall := group.Mixed(top(), left(), right(), bottom())

	return hall
}

func handler(c *gin.Context) {
	target := ConcertHall2F()
	fmt.Fprintf(c.Writer, "<h1>Seats Map</h1><body>%s</body>", target.Html())
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", handler)
	r.Run()
}
