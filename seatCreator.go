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

	rect1 := eraser.Rectangle(pos.Absolute{X: 4, Y: 3}, pos.Absolute{X: 34, Y: 12})
	integrated1 := eraser.Integrated(rect1)
	base := model.NewSeatBase(3, 3, "발코니석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         33,
		YSize:         10,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('O'),
	}


	block1 := group.VerticalBlock(blockInput1)

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
		Key:             "2F",
		Seats:           seats,
		XSize:           33,
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
