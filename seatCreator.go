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
	
	hall := eraser.VerticalHallway(15)
	rect1 := eraser.Rectangle(pos.Absolute{X: 26, Y: 4},pos.Absolute{X: 28, Y: 5})
	rect2 := eraser.Rectangle(pos.Absolute{X: 15, Y: 11},pos.Absolute{X: 20, Y: 12})
	rect3 := eraser.Rectangle(pos.Absolute{X: 8, Y: 12},pos.Absolute{X: 9, Y: 12})
	specific := eraser.Position(pos.Absolute{X: 9, Y: 11})
	integrated := eraser.Integrated(hall, rect1, rect2,rect3, specific)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         26,
		YSize:         9,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
		Reverse:       true,
	}

	return group.HorizontalBlock(blockInput)
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
		XSize:           26,
		YSize:           9,
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
