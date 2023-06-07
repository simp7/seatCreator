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

func ConcertHall1F() model.Group {
	nameFormatter := nameformatter.Standard()

	hall := eraser.HorizontalHallway(7)

	// 2층 무대 앞쪽
	rectSide1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 5, Y: 6})
	rectSide2 := eraser.Rectangle(pos.Absolute{X: 49, Y: 4}, pos.Absolute{X: 52, Y: 6})

	recthall1 := eraser.Rectangle(pos.Absolute{X: 16, Y: 4}, pos.Absolute{X: 16, Y: 6})
	recthall2 := eraser.Rectangle(pos.Absolute{X: 27, Y: 4}, pos.Absolute{X: 27, Y: 6})
	recthall3 := eraser.Rectangle(pos.Absolute{X: 38, Y: 4}, pos.Absolute{X: 38, Y: 6})

	specific1 := eraser.Position(pos.Absolute{X: 6, Y: 4})
	specific2 := eraser.Position(pos.Absolute{X: 48, Y: 4})
	specificEmpty := eraser.Rectangle(pos.Absolute{X: 23, Y: 4}, pos.Absolute{X: 52, Y: 4})

	// 2층 무대 뒷쪽
	recthall4 := eraser.Rectangle(pos.Absolute{X: 8, Y: 8}, pos.Absolute{X: 8,Y: 13})
	recthall5 := eraser.Rectangle(pos.Absolute{X: 19, Y: 8}, pos.Absolute{X: 19,Y: 13})
	recthall6 := eraser.Rectangle(pos.Absolute{X: 27, Y: 8}, pos.Absolute{X: 27,Y: 13})
	recthall7 := eraser.Rectangle(pos.Absolute{X: 35, Y: 8}, pos.Absolute{X: 35,Y: 13})
	recthall8 := eraser.Rectangle(pos.Absolute{X: 46, Y: 8}, pos.Absolute{X: 46,Y: 13})

	rect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 9}, pos.Absolute{X: 3, Y: 13})
	rect2 := eraser.Rectangle(pos.Absolute{X: 4,Y: 10}, pos.Absolute{X: 4, Y: 13})
	rect3 := eraser.Rectangle(pos.Absolute{X: 5,Y: 11}, pos.Absolute{X: 5, Y: 13})
	rect4 := eraser.Rectangle(pos.Absolute{X: 6,Y: 12}, pos.Absolute{X: 6, Y: 13})
	specific3 := eraser.Position(pos.Absolute{X: 7,Y: 13})
	specificEmpty2 := eraser.Rectangle(pos.Absolute{X: 25,Y: 13}, pos.Absolute{X: 51, Y: 13})

	recthall9 := eraser.Rectangle(pos.Absolute{X: 17, Y: 8}, pos.Absolute{X: 18, Y: 12})
	recthall10 := eraser.Rectangle(pos.Absolute{X: 44, Y: 8}, pos.Absolute{X: 45, Y: 12})

	recthall11 := eraser.Rectangle(pos.Absolute{X: 48, Y: 12}, pos.Absolute{X: 51, Y: 12})
	recthall12 := eraser.Rectangle(pos.Absolute{X: 49, Y: 11}, pos.Absolute{X: 51, Y: 11})
	recthall13 := eraser.Rectangle(pos.Absolute{X: 50, Y: 10}, pos.Absolute{X: 51, Y: 10})
	specific4 := eraser.Position(pos.Absolute{X: 51,Y: 9})

	integrated := eraser.Integrated(hall, rectSide1, rectSide2, recthall1,recthall2,recthall3,
									specific1,specific2, specificEmpty,
									recthall4,recthall5,recthall6,recthall7,recthall8,
									rect1,rect2,rect3,rect4, specific3, recthall9,
									specificEmpty2, recthall10,
									recthall11,recthall12,recthall13, specific4)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         49,
		YSize:         10,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	integrated1 := eraser.Integrated(recthall3)
	base = model.NewSeatBase(32, 4, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 24, 1),
		XSize:         16,
		YSize:         1,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	integrated2 := eraser.Integrated(recthall7)
	base = model.NewSeatBase(31, 13, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 20, 9),
		XSize:         15,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
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
	seats := ConcertHall1F() // Put Seating Here
	target := area.Area{
		Key:             "2F",
		Seats:           seats,
		XSize:           45,
		YSize:           10,
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