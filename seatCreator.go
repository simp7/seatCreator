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
	
	rect1 := eraser.Rectangle(pos.Absolute{X: 7, Y: 4}, pos.Absolute{X: 7, Y: 12})
	rect2 := eraser.Rectangle(pos.Absolute{X: 20, Y: 4}, pos.Absolute{X: 20, Y: 14})
	specific := eraser.Position(pos.Absolute{X: 3, Y: 4},
								pos.Absolute{X: 19, Y: 4}, pos.Absolute{X: 19, Y: 6}, pos.Absolute{X: 19, Y: 8})
	rect3 := eraser.Rectangle(pos.Absolute{X: 3, Y: 13}, pos.Absolute{X: 4, Y: 13})
	rect4 := eraser.Rectangle(pos.Absolute{X: 10, Y: 14}, pos.Absolute{X: 16, Y: 14})
	rect5 := eraser.Rectangle(pos.Absolute{X: 11, Y: 15}, pos.Absolute{X: 15, Y: 15})
	integrated1 := eraser.Integrated(rect1, rect2, specific, rect3, rect4, rect5)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         22,
		YSize:         12,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput1)

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
		YSize:           12,
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
