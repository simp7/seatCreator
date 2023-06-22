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

	rect1 := eraser.Rectangle(pos.Absolute{X: 10, Y: 10}, pos.Absolute{X: 11, Y: 13})
	rect2 := eraser.Rectangle(pos.Absolute{X: 8, Y: 14}, pos.Absolute{X: 14, Y: 16})
	integrated1 := eraser.Integrated(rect1, rect2)
	base := model.NewSeatBase(3, 3, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         9,
		YSize:         14,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(13, 3, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 10, 1),
		XSize:         9,
		YSize:         7,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

		base = model.NewSeatBase(15, 10, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 12, 1),
		XSize:         7,
		YSize:         4,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

		base = model.NewSeatBase(17, 14, "A석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 14, 1),
		XSize:         5,
		YSize:         3,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)

	return group.Mixed(block1, block2, block3, block4)
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
		XSize:           19,
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
