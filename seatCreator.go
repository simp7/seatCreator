package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/gin-gonic/gin"
	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/area"
	"github.com/simp7/seatCreator/model/eraser"
	"github.com/simp7/seatCreator/model/group"
	"github.com/simp7/seatCreator/model/nameformatter"
	"github.com/simp7/seatCreator/model/pos"
)

func ArtriumSmall() model.Group {
	nameFormatter := nameformatter.Standard()

	hall2 := eraser.HorizontalHallway(14)

	Lrect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 3}, pos.Absolute{X: 8, Y: 3})
	Lrect2 := eraser.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 5, Y: 4})
	Lrect3 := eraser.Rectangle(pos.Absolute{X: 3, Y: 5}, pos.Absolute{X: 3, Y: 6})
	Lrect4 := eraser.Rectangle(pos.Absolute{X: 3, Y: 13}, pos.Absolute{X: 10, Y: 13})

	Lhall := eraser.Rectangle(pos.Absolute{X: 9, Y: 3}, pos.Absolute{X: 9, Y: 14})

	Rrect1 := eraser.Rectangle(pos.Absolute{X: 25, Y: 3},pos.Absolute{X: 30, Y: 3})
	Rrect2 := eraser.Rectangle(pos.Absolute{X: 28, Y: 4},pos.Absolute{X: 30, Y: 4})
	Rrect3 := eraser.Rectangle(pos.Absolute{X: 30, Y: 5},pos.Absolute{X: 30, Y: 6})
	Rrect4 := eraser.Rectangle(pos.Absolute{X: 23, Y: 13},pos.Absolute{X: 30, Y: 13})

	Rhall := eraser.Rectangle(pos.Absolute{X: 24, Y: 3}, pos.Absolute{X: 24, Y: 14})

	Mrect1 := eraser.Rectangle(pos.Absolute{X: 15, Y: 15}, pos.Absolute{X: 18, Y: 16})
	Mrect2 := eraser.Rectangle(pos.Absolute{X: 10, Y: 15}, pos.Absolute{X: 10, Y: 16})
	Mrect3 := eraser.Rectangle(pos.Absolute{X: 23, Y: 15}, pos.Absolute{X: 23, Y: 16})

	specific := eraser.Position(pos.Absolute{X: 11, Y: 16}, pos.Absolute{X: 22, Y: 16})

	integrated1 := eraser.Integrated(hall2, Lrect1, Lrect2, Lrect3, Lrect4, Lhall,
									Rrect1,Rrect2,Rrect3,Rrect4, Rhall,
									Mrect1, Mrect2,Mrect3, specific)
	base := model.NewSeatBase(3, 3, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         28,
		YSize:         14,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(1, 8, "BOX석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         5,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(32, 8, "BOX석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 2),
		XSize:         1,
		YSize:         5,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.VerticalBlock(blockInput2)
	block3 := group.VerticalBlock(blockInput3)

	return group.Mixed(block1, block2, block3)
}


func copy(target model.Group) {
	err := clipboard.WriteAll(target.String())
	if err != nil {
		fmt.Println("Error occured when copying")
		return
	}
	fmt.Println("Copied!")
}

func handler(c *gin.Context) {
	seats := ArtriumSmall() // Put Seating Here
	target := area.Area{
		Key:             "1F",
		Seats:           seats,
		XSize:           33,
		YSize:           14,
		BackgroundImage: "",
		Color:           "#ff9f00",
	}
	fmt.Fprintf(c.Writer, "<h1>Seats Map</h1><body>\n%s\n</body>", target.Html())
	copy(target)
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", handler)
	r.Run()
}
