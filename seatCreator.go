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

	hall := eraser.HorizontalHallway(18)

	rectAsection := eraser.Rectangle(pos.Absolute{X: 3,Y: 4},pos.Absolute{X: 3,Y: 13})
	rectAsection1 := eraser.Rectangle(pos.Absolute{X: 4,Y: 4},pos.Absolute{X: 4,Y: 11})
	rectAsection2 := eraser.Rectangle(pos.Absolute{X: 5,Y: 4},pos.Absolute{X: 5,Y: 10})
	rectAsection3 := eraser.Rectangle(pos.Absolute{X: 6,Y: 4},pos.Absolute{X: 6,Y: 8})
	rectAsection4 := eraser.Rectangle(pos.Absolute{X: 7,Y: 4},pos.Absolute{X: 7,Y: 6})
	rectAsection5 := eraser.Rectangle(pos.Absolute{X: 8,Y: 4},pos.Absolute{X: 8,Y: 5})
	integrated := eraser.Integrated(hall, rectAsection, rectAsection1, rectAsection2,rectAsection3,rectAsection4,rectAsection5)
	base := model.NewSeatBase(3, 5, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         23,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	specificBsection := eraser.Position(pos.Absolute{X: 36,Y: 6},
										pos.Absolute{X: 36,Y: 8},
										pos.Absolute{X: 36,Y: 10},
										pos.Absolute{X: 36,Y: 12},
										pos.Absolute{X: 36,Y: 14},
										pos.Absolute{X: 36,Y: 16},
									
										pos.Absolute{X: 36,Y: 20},
										pos.Absolute{X: 36,Y: 22},
										pos.Absolute{X: 36,Y: 24},
										pos.Absolute{X: 36,Y: 26},)
	rectBsection := eraser.Rectangle(pos.Absolute{X: 33,Y: 17}, pos.Absolute{X: 36, Y: 17})
	empty := eraser.HorizontalHallway(19)

	integrated2 := eraser.Integrated(hall, specificBsection, rectBsection, empty)
	base = model.NewSeatBase(16, 5, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         21,
		YSize:         13,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	integrated5 := eraser.Integrated()
	base = model.NewSeatBase(20, 19, "B석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 5, 14),
		XSize:         15,
		YSize:         1,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(17, 28, "B석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 2, 23),
		XSize:         16,
		YSize:         1,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	integrated7 := eraser.Integrated(specificBsection)
	base = model.NewSeatBase(16, 20, "B석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 2, 15),
		XSize:         21,
		YSize:         8,
		EmptyChecker:  integrated7,
		NameFormatter: nameFormatter,
	}

	rectCsection := eraser.Rectangle(pos.Absolute{X: 49,Y: 5},pos.Absolute{X: 49, Y: 13})
	rectCsection2 := eraser.Rectangle(pos.Absolute{X: 48,Y: 5},pos.Absolute{X: 48, Y: 11})
	rectCsection3 := eraser.Rectangle(pos.Absolute{X: 47,Y: 5},pos.Absolute{X: 47, Y: 10})
	rectCsection4 := eraser.Rectangle(pos.Absolute{X: 46,Y: 5},pos.Absolute{X: 46, Y: 8})
	rectCsection5 := eraser.Rectangle(pos.Absolute{X: 45,Y: 5},pos.Absolute{X: 45, Y: 6})
	specificCsection := eraser.Position(pos.Absolute{X: 44,Y: 5})
	integrated3 := eraser.Integrated(hall, rectCsection, rectCsection2,rectCsection3,rectCsection4,rectCsection5, specificCsection)
	base = model.NewSeatBase(38, 5, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         23,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	hallOP := eraser.VerticalHallway(18,34)
	specificOPsection := eraser.Position(pos.Absolute{X: 12,Y: 1}, pos.Absolute{X: 13, Y: 1}, pos.Absolute{X: 12, Y: 2}, 
										 pos.Absolute{X: 33,Y: 2},
										 pos.Absolute{X: 40,Y: 1}, pos.Absolute{X: 39,Y: 1}, pos.Absolute{X: 40,Y: 2})
	integrated4 := eraser.Integrated(specificOPsection, hallOP)
	base = model.NewSeatBase(12,1, "OP석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         29,
		YSize:         3,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}

	specificWsection := eraser.Position(pos.Absolute{X: 4, Y: 28},
										pos.Absolute{X: 6, Y: 28},
										pos.Absolute{X: 8, Y: 28},
										pos.Absolute{X: 10, Y: 28},
										pos.Absolute{X: 12, Y: 28},
										pos.Absolute{X: 14, Y: 28},
										pos.Absolute{X: 15, Y: 28},
										pos.Absolute{X: 37, Y: 28},
										pos.Absolute{X: 39, Y: 28},
										pos.Absolute{X: 41, Y: 28},
										pos.Absolute{X: 43, Y: 28},
										pos.Absolute{X: 45, Y: 28},
										pos.Absolute{X: 47, Y: 28},
										pos.Absolute{X: 49, Y: 28},

									)
	rectOPsection := eraser.Rectangle(pos.Absolute{X: 17,Y: 28}, pos.Absolute{X: 35,Y: 28})

	integrated6 := eraser.Integrated(specificWsection, rectOPsection)
	base = model.NewSeatBase(3,28, "휠체어석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         47,
		YSize:         1,
		EmptyChecker:  integrated6,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)
	block8 := group.HorizontalBlock(blockInput8)
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7, block8)
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
		XSize:           47,
		YSize:           28,
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