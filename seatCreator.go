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
	hall1 := eraser.VerticalHallway(12,13)
	rect1 := eraser.Rectangle(pos.Absolute{X: 2, Y: 5}, pos.Absolute{X: 2, Y: 6})
	rect2 := eraser.Rectangle(pos.Absolute{X: 24, Y: 4}, pos.Absolute{X: 24, Y: 10})
	rect3 := eraser.Rectangle(pos.Absolute{X: 23, Y: 5}, pos.Absolute{X: 23, Y: 6})
	rect4 := eraser.Rectangle(pos.Absolute{X: 2, Y: 11}, pos.Absolute{X: 2, Y: 13})
	rect5 := eraser.Rectangle(pos.Absolute{X: 23, Y: 12}, pos.Absolute{X: 24, Y: 13})

	specific := eraser.Position(pos.Absolute{X: 3,Y: 12}, pos.Absolute{X: 22,Y: 12})

	integrated := eraser.Integrated(hall1,rect1, rect2,rect3,rect4,rect5, specific)
	nameFormatter := nameformatter.Standard()

	base := model.NewSeatBase(2, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         23,
		YSize:         10,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}


	integrated2 := eraser.Integrated()
	base = model.NewSeatBase(4, 14, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 11),
		XSize:         13,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	// rect6 := eraser.Rectangle(pos.Absolute{X: 16, Y: 16}, pos.Absolute{X: 17, Y: 17})
	// specific2 := eraser.Position(pos.Absolute{X: 20,Y: 16})
	// integrated3 := eraser.Integrated(rect6,specific2)

	// base = model.NewSeatBase(6, 16, "A석")
	// blockInput3 := group.BlockInput{
	// 	Criteria:      model.NewSeat(base, 1, 12),
	// 	XSize:         15,
	// 	YSize:         2,
	// 	EmptyChecker:  integrated3,
	// 	NameFormatter: nameFormatter,
	// }

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	// block3 := group.HorizontalBlock(blockInput3)
	return group.Mixed(block1,block2)
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
		XSize:           22,
		YSize:           11,
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