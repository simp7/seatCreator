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
 	
	GaRect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 11, Y: 4})
	GaRect2 := eraser.Rectangle(pos.Absolute{X: 3, Y: 5}, pos.Absolute{X: 8, Y: 5})
	GaRect3 := eraser.Rectangle(pos.Absolute{X: 3, Y: 6}, pos.Absolute{X: 7, Y: 6})
	GaRect4 := eraser.Rectangle(pos.Absolute{X: 3, Y: 7}, pos.Absolute{X: 5, Y: 7})
	GaRect5 := eraser.Rectangle(pos.Absolute{X: 3, Y: 8}, pos.Absolute{X: 4, Y: 8})
	GaRect6 := eraser.Rectangle(pos.Absolute{X: 3, Y: 9}, pos.Absolute{X: 4, Y: 9})
	GaRect7 := eraser.Rectangle(pos.Absolute{X: 3, Y: 10}, pos.Absolute{X: 6, Y: 10})
	GaRect8 := eraser.Rectangle(pos.Absolute{X: 3, Y: 11}, pos.Absolute{X: 4, Y: 11})
	GaRect9 := eraser.Rectangle(pos.Absolute{X: 3, Y: 12}, pos.Absolute{X: 4, Y: 12})
	GaRect10 := eraser.Rectangle(pos.Absolute{X: 3, Y: 13}, pos.Absolute{X: 3, Y: 13})
	GaSpecific := eraser.Position(pos.Absolute{X: 15, Y: 5}, pos.Absolute{X: 15, Y: 6},
								  pos.Absolute{X: 14, Y: 7}, pos.Absolute{X: 15, Y: 7},
								  pos.Absolute{X: 14, Y: 8}, pos.Absolute{X: 15,Y: 8},
								  pos.Absolute{X: 13, Y: 9}, pos.Absolute{X: 14, Y: 9},
								  pos.Absolute{X: 13, Y: 10}, pos.Absolute{X: 14, Y: 10},
								  pos.Absolute{X: 12, Y: 11}, pos.Absolute{X: 13, Y: 11},
								  pos.Absolute{X: 12, Y: 12}, pos.Absolute{X: 13, Y: 12},
								  pos.Absolute{X: 11, Y: 13}, pos.Absolute{X: 12, Y: 13},
								)
	GaRect11 := eraser.Rectangle(pos.Absolute{X: 10, Y: 15}, pos.Absolute{X: 15, Y: 15})
	GaRect12 := eraser.Rectangle(pos.Absolute{X: 14, Y: 16}, pos.Absolute{X: 15, Y: 16})
	GaRect13 := eraser.Rectangle(pos.Absolute{X: 14, Y: 17}, pos.Absolute{X: 15, Y: 17})
	hall := eraser.HorizontalHallway(14)
	integrated := eraser.Integrated(hall, GaRect1, GaRect2, GaSpecific, GaRect3, GaRect4, GaRect5,
							GaRect6, GaRect7, GaRect8,GaRect9,GaRect10,
							GaRect11, GaRect12, GaRect13)
	base := model.NewSeatBase(3, 4, "가석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         13,
		YSize:         15,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	
	rect1 := eraser.Rectangle(pos.Absolute{X: 13,Y: 4}, pos.Absolute{X: 13, Y: 12})
	rect2 := eraser.Rectangle(pos.Absolute{X: 34,Y: 4}, pos.Absolute{X: 34, Y: 11})
	rect3 := eraser.Rectangle(pos.Absolute{X: 14,Y: 4}, pos.Absolute{X: 14, Y: 10})
	rect4 := eraser.Rectangle(pos.Absolute{X: 33,Y: 4}, pos.Absolute{X: 33, Y: 9})
	rect5 := eraser.Rectangle(pos.Absolute{X: 15,Y: 4}, pos.Absolute{X: 15, Y: 8})
	rect6 := eraser.Rectangle(pos.Absolute{X: 32,Y: 4}, pos.Absolute{X: 32, Y: 7})
	rect7 := eraser.Rectangle(pos.Absolute{X: 16,Y: 4}, pos.Absolute{X: 16, Y: 6})
	rect8 := eraser.Rectangle(pos.Absolute{X: 31,Y: 4}, pos.Absolute{X: 31, Y: 5})
	rect9 := eraser.Rectangle(pos.Absolute{X: 17,Y: 4}, pos.Absolute{X: 17, Y: 4})
	
	integrated2 := eraser.Integrated(rect1, rect2, rect3, rect4, rect5, rect6, rect7, rect8, rect9)
	base = model.NewSeatBase(13, 4, "나석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         22,
		YSize:         10,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}


	DaRect1 := eraser.Rectangle(pos.Absolute{X: 38, Y: 4}, pos.Absolute{X: 43,Y: 4})
	DaRect2 := eraser.Rectangle(pos.Absolute{X: 40, Y: 5}, pos.Absolute{X: 43,Y: 5})
	DaRect3 := eraser.Rectangle(pos.Absolute{X: 42, Y: 6}, pos.Absolute{X: 43,Y: 7})
	DaRect4 := eraser.Rectangle(pos.Absolute{X: 43, Y: 8}, pos.Absolute{X: 43,Y: 9})
	DaRect5 := eraser.Rectangle(pos.Absolute{X: 33, Y: 8}, pos.Absolute{X: 34,Y: 9})
	DaRect6 := eraser.Rectangle(pos.Absolute{X: 33, Y: 10}, pos.Absolute{X: 35, Y: 11})
	DaRect7 := eraser.Rectangle(pos.Absolute{X: 42, Y: 10}, pos.Absolute{X: 43, Y: 10})
	DaRect8 := eraser.Rectangle(pos.Absolute{X: 33, Y: 12}, pos.Absolute{X: 36, Y: 15})
	DaRect9 := eraser.Rectangle(pos.Absolute{X: 33, Y: 14}, pos.Absolute{X: 38, Y: 15})

	DaSpecific1 := eraser.Position(pos.Absolute{X: 33, Y: 6}, pos.Absolute{X: 33, Y: 7},
								pos.Absolute{X: 43, Y: 11}, pos.Absolute{X: 39, Y: 15})
	integrated3 := eraser.Integrated(DaRect1, DaRect2, DaRect3, DaSpecific1, DaRect4,DaRect5,
									DaRect6, DaRect7, DaRect8, DaRect9)
	base = model.NewSeatBase(33, 4, "다석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         12,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	RaSpecific1 := eraser.Position(pos.Absolute{X: 18,Y: 17})
	RaRect1 := eraser.Rectangle(pos.Absolute{X: 18,Y: 19}, pos.Absolute{X: 30, Y: 19})
	RaRect2 := eraser.Rectangle(pos.Absolute{X: 18,Y: 20}, pos.Absolute{X: 30, Y: 20})
	integrated4 := eraser.Integrated(RaSpecific1, RaRect1,RaRect2)
	base = model.NewSeatBase(18, 15, "라석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         17,
		YSize:         6,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(14, 15, "휠체어석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         4,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}
	

	base = model.NewSeatBase(38, 14, "휠체어석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}

	base = model.NewSeatBase(39, 15, "휠체어석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}


	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
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
		XSize:           36,
		YSize:           17,
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
