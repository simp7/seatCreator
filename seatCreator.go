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
 	
	hallH := eraser.HorizontalHallway(10)
	hallV := eraser.VerticalHallway(13)
	specific1 := eraser.Position(pos.Absolute{X: 12,Y: 5 }, pos.Absolute{X: 12, Y: 9},
								pos.Absolute{X: 22,Y: 5}, pos.Absolute{X: 22,Y: 9})
	rect1 := eraser.Rectangle(pos.Absolute{X: 20, Y: 11}, pos.Absolute{X: 22, Y:  11})
	rect2 := eraser.Rectangle(pos.Absolute{X: 14,Y: 16}, pos.Absolute{X: 22, Y: 16})

	integrated := eraser.Integrated(hallH, hallV, specific1, rect1, rect2)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         20,
		YSize:         13,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
		Reverse:	   true,	
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
		XSize:           20,
		YSize:           13,
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
