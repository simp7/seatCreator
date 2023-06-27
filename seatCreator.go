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
	
	hallH := eraser.HorizontalHallway(13)
	rect1 := eraser.Rectangle(pos.Absolute{X: 7, Y: 4}, pos.Absolute{X: 7, Y: 17})
	rect2 := eraser.Rectangle(pos.Absolute{X: 20, Y: 4}, pos.Absolute{X: 20, Y: 17})
	specific := eraser.Position(pos.Absolute{X: 19, Y: 4}, pos.Absolute{X: 19, Y: 6}, pos.Absolute{X: 19, Y: 8}, pos.Absolute{X: 19, Y: 10},
								pos.Absolute{X: 24, Y: 4}, pos.Absolute{X: 24, Y: 5},
								pos.Absolute{X: 24, Y: 14}, pos.Absolute{X: 24, Y: 17},
								pos.Absolute{X: 19, Y: 18})
	integrated1 := eraser.Integrated(hallH, rect1, rect2, specific)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         22,
		YSize:         15,
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
		YSize:           15,
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
