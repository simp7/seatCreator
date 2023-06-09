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

	specific := eraser.Position(pos.Absolute{X: 3, Y: 5}, pos.Absolute{X: 3,Y: 7})
	rect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 9}, pos.Absolute{X: 3, Y: 13})
	rect2 := eraser.Rectangle(pos.Absolute{X: 4, Y: 11}, pos.Absolute{X: 4, Y: 12})
	rect3 := eraser.Rectangle(pos.Absolute{X: 10, Y: 12}, pos.Absolute{X: 17, Y: 13})
	rect4 := eraser.Rectangle(pos.Absolute{X: 23, Y: 11}, pos.Absolute{X: 23, Y: 12})
	rect5 := eraser.Rectangle(pos.Absolute{X: 24, Y: 9}, pos.Absolute{X: 24, Y: 13})
	integrated := eraser.Integrated(rect1, specific, rect2, rect3, rect4, rect5)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 14),
		XSize:         22,
		YSize:         10,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
		Reverse:       true,
	}

	block1 := group.HorizontalBlock(blockInput)
	return group.Mixed(block1)
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
