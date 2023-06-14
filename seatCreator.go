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


	Arect1 := eraser.Rectangle(pos.Absolute{X: 9, Y: 16}, pos.Absolute{X: 11, Y: 24})
	Arect2 := eraser.Rectangle(pos.Absolute{X: 10, Y: 12}, pos.Absolute{X: 11, Y: 15})
	Arect3 := eraser.Rectangle(pos.Absolute{X: 11, Y: 10}, pos.Absolute{X: 11, Y: 11})
	integrated := eraser.Integrated(Arect1, Arect2, Arect3)
	base := model.NewSeatBase(5, 6, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         7,
		YSize:         19,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 11, Y: 6}, pos.Absolute{X: 11, Y: 13})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 31, Y: 6}, pos.Absolute{X: 31, Y: 13})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 12, Y: 6}, pos.Absolute{X: 12, Y: 12})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 30, Y: 6}, pos.Absolute{X: 30, Y: 10})
	Brect5 := eraser.Rectangle(pos.Absolute{X: 13, Y: 6}, pos.Absolute{X: 13, Y: 9})
	Brect6 := eraser.Rectangle(pos.Absolute{X: 29, Y: 6}, pos.Absolute{X: 29, Y: 8})
	Brect7 := eraser.Rectangle(pos.Absolute{X: 14, Y: 6}, pos.Absolute{X: 14, Y: 8})
	Brect8 := eraser.Rectangle(pos.Absolute{X: 14, Y: 6}, pos.Absolute{X: 14, Y: 8})
	Brect9 := eraser.Rectangle(pos.Absolute{X: 28, Y: 6}, pos.Absolute{X: 28, Y: 7})

	Brect10 := eraser.Rectangle(pos.Absolute{X: 11, Y: 14}, pos.Absolute{X: 12, Y: 14})
	Brect11 := eraser.Rectangle(pos.Absolute{X: 30, Y: 14}, pos.Absolute{X: 31, Y: 14})
	Brect12 := eraser.Rectangle(pos.Absolute{X: 11, Y: 15}, pos.Absolute{X: 12, Y: 16})
	Brect13 := eraser.Rectangle(pos.Absolute{X: 31, Y: 15}, pos.Absolute{X: 31, Y: 16})

	Bspecific1 := eraser.Position(pos.Absolute{X: 15,Y: 6}, pos.Absolute{X: 31, Y: 18},
								pos.Absolute{X: 31,Y: 20}, pos.Absolute{X: 31, Y: 22},
								pos.Absolute{X: 31, Y: 24},)
	integrated2 := eraser.Integrated(Brect1, Brect2, Brect3, Brect4, Brect5, Brect6, Brect7, Brect8, Bspecific1,
								Brect9, Brect10, Brect11, Brect12, Brect13)
	base = model.NewSeatBase(11, 6, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         21,
		YSize:         19,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Crect1 := eraser.Rectangle(pos.Absolute{X:31, Y: 16 }, pos.Absolute{X: 33, Y:24 })
	Crect2 := eraser.Rectangle(pos.Absolute{X:31, Y: 12 }, pos.Absolute{X: 32, Y:15 })
	Crect3 := eraser.Rectangle(pos.Absolute{X:31, Y: 10 }, pos.Absolute{X: 31, Y:11 })
	integrated3 := eraser.Integrated(Crect1, Crect2, Crect3)
	base = model.NewSeatBase(31, 6, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         7,
		YSize:         19,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	hallOP := eraser.VerticalHallway(14, 28)
	OPrect1 := eraser.Rectangle(pos.Absolute{X: 8, Y: 2}, pos.Absolute{X: 14, Y: 2})
	OPrect2 := eraser.Rectangle(pos.Absolute{X: 28, Y: 2}, pos.Absolute{X: 34, Y: 2})
	OPspecific := eraser.Position(pos.Absolute{X: 8, Y: 3}, pos.Absolute{X: 34, Y: 3})
	integrated4 := eraser.Integrated(hallOP, OPrect1, OPrect2, OPspecific)
	base = model.NewSeatBase(8, 2, "OP석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         27,
		YSize:         3,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}

	BOXrect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 18}, pos.Absolute{X: 3, Y: 20})
	BOXrect2 := eraser.Rectangle(pos.Absolute{X: 39, Y: 18}, pos.Absolute{X: 39, Y: 20})
	base = model.NewSeatBase(2, 6, "BOX석")
	integrated5 := eraser.Integrated(BOXrect1, BOXrect2)
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         2,
		YSize:         15,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(39, 6, "BOX석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 3),
		XSize:         2,
		YSize:         15,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	Wrect1 := eraser.Rectangle(pos.Absolute{X:9, Y: 25 }, pos.Absolute{X: 10, Y: 25})
	Wrect2 := eraser.Rectangle(pos.Absolute{X:12, Y: 25 }, pos.Absolute{X: 29, Y: 25})
	Wrect3 := eraser.Rectangle(pos.Absolute{X:31, Y: 25 }, pos.Absolute{X: 33, Y: 25})
	integrated7 := eraser.Integrated(Wrect1, Wrect2, Wrect3)
	base = model.NewSeatBase(7, 25, "휠체어석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 3),
		XSize:         29,
		YSize:         1,
		EmptyChecker:  integrated7,
		NameFormatter: nameformatter.Prefix('W'),
	}
	
	
	
	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.VerticalBlock(blockInput5)
	block6 := group.VerticalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7)
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
		XSize:           41,
		YSize:           20,
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
