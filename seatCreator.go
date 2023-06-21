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

	hall := eraser.VerticalHallway(10, 25)
	rect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 14}, pos.Absolute{X: 11, Y: 14})
	rect2 := eraser.Rectangle(pos.Absolute{X: 23, Y: 14}, pos.Absolute{X: 32, Y: 14})
	integrated1 := eraser.Integrated(hall, rect1, rect2)
	base := model.NewSeatBase(3, 11, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         30,
		YSize:         4,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(1, 5, "BOX석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         3,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(34, 5, "BOX석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 2),
		XSize:         1,
		YSize:         3,
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
		Key:             "3F",
		Seats:           seats,
		XSize:           32,
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
