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

	Aspecific := eraser.Position(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 3, Y: 8}, pos.Absolute{X: 3, Y: 9},
								pos.Absolute{X: 3,	Y: 16}, pos.Absolute{X: 3,Y: 17}, pos.Absolute{X: 3,Y: 18})

	Brect1 := eraser.Rectangle(pos.Absolute{X: 10, Y: 4}, pos.Absolute{X: 13, Y: 4})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 21, Y: 4}, pos.Absolute{X: 24, Y: 4})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 10, Y: 20}, pos.Absolute{X: 24, Y: 20})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 17, Y: 19}, pos.Absolute{X: 24, Y: 19} )
	Bspecific := eraser.Position(pos.Absolute{X: 24, Y: 5},
								 pos.Absolute{X: 24, Y: 7},
								 pos.Absolute{X: 24, Y: 9},
								 pos.Absolute{X: 24, Y: 11},
								 pos.Absolute{X: 24, Y: 13},
								 pos.Absolute{X: 24, Y: 15},
								 pos.Absolute{X: 24, Y: 17},
	)

	Cspecific := eraser.Position(pos.Absolute{X: 31, Y: 4}, pos.Absolute{X: 31, Y: 8}, pos.Absolute{X: 31, Y: 9},
								pos.Absolute{X: 31, Y:16 }, pos.Absolute{X: 31, Y: 17}, pos.Absolute{X: 31,Y: 18})

	integrated1 := eraser.Integrated(Aspecific, Brect1,Brect2,Brect3,Brect4, Bspecific, Cspecific)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         17,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(10, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         15,
		YSize:         17,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(26, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         17,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	OPrect1 := eraser.Rectangle(pos.Absolute{X: 7, Y: 1}, pos.Absolute{X: 8, Y: 1})
	OPrect2 := eraser.Rectangle(pos.Absolute{X: 25, Y: 1}, pos.Absolute{X: 26, Y: 1})
	hallV := eraser.VerticalHallway(9, 24)
	integrated2 := eraser.Integrated(hallV, OPrect1, OPrect2)
	base = model.NewSeatBase(7, 1, "OP석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         20,
		YSize:         2,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Wrect1 := eraser.Rectangle(pos.Absolute{X: 12, Y: 4}, pos.Absolute{X: 21,Y: 4})
	Wrect2 := eraser.Rectangle(pos.Absolute{X: 10, Y: 5}, pos.Absolute{X: 23, Y: 18})
	Wrect3 := eraser.Rectangle(pos.Absolute{X: 10, Y: 19}, pos.Absolute{X: 19, Y: 19})
	integrated3 := eraser.Integrated(Wrect1, Wrect2, Wrect3)
	base = model.NewSeatBase(10, 4, "휠체어석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         16,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('W'),
	}


	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)


	return group.Mixed(block1, block2, block3, block4, block5)
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
		XSize:           29,
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
