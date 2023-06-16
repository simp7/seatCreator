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

	rect1 := eraser.Rectangle(pos.Absolute{X: 4, Y: 15}, pos.Absolute{X: 4, Y: 16})
	rect2 := eraser.Rectangle(pos.Absolute{X: 17, Y: 15}, pos.Absolute{X: 17, Y: 16})
	rect3 := eraser.Rectangle(pos.Absolute{X: 28, Y: 15}, pos.Absolute{X: 28, Y: 16})
	rect4 := eraser.Rectangle(pos.Absolute{X: 41, Y: 15}, pos.Absolute{X: 41, Y: 16})
	specific := eraser.Position(pos.Absolute{X: 5, Y: 15}, pos.Absolute{X: 16, Y: 15},
								pos.Absolute{X: 27 ,Y: 15},
								pos.Absolute{X: 29, Y: 15}, pos.Absolute{X: 40, Y: 15},
								pos.Absolute{X: 45, Y: 15})
	integrated3 := eraser.Integrated(rect1,rect2,rect3, rect4 ,specific)
	base := model.NewSeatBase(1, 15, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         45,
		YSize:         2,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Standard(),
	}

	block3 := group.HorizontalBlock(blockInput3)

	return group.Mixed(block3)
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
		XSize:           45,
		YSize:           2,
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
