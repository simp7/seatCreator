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

	hall := eraser.HorizontalHallway(7)
	rect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 3}, pos.Absolute{X: 3, Y: 6})
	rect2 := eraser.Rectangle(pos.Absolute{X: 24, Y: 3}, pos.Absolute{X: 24, Y: 6})
	rect3 := eraser.Rectangle(pos.Absolute{X: 3, Y: 8}, pos.Absolute{X: 8, Y: 8})
	rect4 := eraser.Rectangle(pos.Absolute{X: 3, Y: 9}, pos.Absolute{X: 6, Y: 9})
	integrated1 := eraser.Integrated(hall, rect1, rect2,rect3,rect4)
	base := model.NewSeatBase(3, 3, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         22,
		YSize:         10,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	integrated2 := eraser.Integrated()
	base = model.NewSeatBase(4, 9, "휠체어석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         2,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('W'),
	}


	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.HorizontalBlock(blockInput2)

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
	seats := ArtriumSmall() // Put Seating Here
	target := area.Area{
		Key:             "1F",
		Seats:           seats,
		XSize:           22,
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
