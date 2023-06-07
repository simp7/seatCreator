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

	specific := eraser.Position(pos.Absolute{X: 8, Y: 4}, 
								pos.Absolute{X: 3,Y: 8}, pos.Absolute{X: 3,Y: 9},
								pos.Absolute{X: 3, Y: 16}, pos.Absolute{X: 3, Y: 17}, pos.Absolute{X: 3, Y: 18})
	integrated := eraser.Integrated(specific)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         17,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	specific2 := eraser.Position(pos.Absolute{X: 24, Y: 5}, 
								 pos.Absolute{X: 24, Y: 7},
								 pos.Absolute{X: 24, Y: 9},
								 pos.Absolute{X: 24, Y: 11},
								 pos.Absolute{X: 24, Y: 13},
								 pos.Absolute{X: 24, Y: 15},
								 pos.Absolute{X: 24, Y: 17},
								)
	rect := eraser.Rectangle(pos.Absolute{X: 10, Y: 20}, pos.Absolute{X: 25, Y: 20})
	rect2 := eraser.Rectangle(pos.Absolute{X: 17, Y:  19}, pos.Absolute{X: 24, Y: 19})
	rect3 := eraser.Rectangle(pos.Absolute{X: 10,Y: 4},pos.Absolute{X: 13, Y: 4})
	rect4 := eraser.Rectangle(pos.Absolute{X: 21,Y: 4},pos.Absolute{X: 24, Y: 4})

	integrated2 := eraser.Integrated(specific2, rect, rect2, rect3, rect4)
	base = model.NewSeatBase(10, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         15,
		YSize:         17,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	specific3 := eraser.Position(pos.Absolute{X: 26, Y: 4},
								 pos.Absolute{X: 31, Y: 8},
								 pos.Absolute{X: 31, Y: 9},
								 pos.Absolute{X: 31, Y: 16},
								 pos.Absolute{X: 31, Y: 17},
								 pos.Absolute{X: 31, Y: 18},
								)
	integrated3 := eraser.Integrated(specific3)
	base = model.NewSeatBase(26, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         17,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}


	rect5 := eraser.Rectangle(pos.Absolute{X: 7, Y: 1}, pos.Absolute{X: 9, Y: 1})
	rect6 := eraser.Rectangle(pos.Absolute{X: 24, Y: 1}, pos.Absolute{X: 26, Y: 1})
	specific4 := eraser.Position(pos.Absolute{X: 9,Y: 2}, pos.Absolute{X: 24, Y: 2})
	integrated4 := eraser.Integrated(rect5, rect6, specific4)
	base = model.NewSeatBase(7, 1, "OP석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         20,
		YSize:         2,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}


	rect7 := eraser.Rectangle(pos.Absolute{X: 14, Y: 4}, pos.Absolute{X: 20, Y: 4})
	integrated5 := eraser.Integrated(rect7)
	base = model.NewSeatBase(12, 4, "휠체어석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         1,
		EmptyChecker:  integrated5,
		NameFormatter: nameformatter.Prefix('W'),
	}

	base = model.NewSeatBase(17, 19, "휠체어석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         4,
		YSize:         1,
		EmptyChecker:  integrated5,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	return group.Mixed(block1, block2, block3, block4, block5, block6)
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