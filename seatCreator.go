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

func ConcertHall1F() model.Group {
	nameFormatter := nameformatter.Standard()
	
	hall := eraser.HorizontalHallway(9)
	hall2 := eraser.VerticalHallway(8, 21)
	rect := eraser.Rectangle(pos.Absolute{X: 9, Y: 14}, pos.Absolute{X: 11, Y: 14})
	rect2 := eraser.Rectangle(pos.Absolute{X: 18, Y: 14}, pos.Absolute{X: 20, Y: 14})
	integrated := eraser.Integrated(hall, hall2, rect, rect2)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         24,
		YSize:         11,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	rect3 := eraser.Rectangle(pos.Absolute{X: 3,Y: 15},pos.Absolute{X: 6,Y: 15})
	rect4 := eraser.Rectangle(pos.Absolute{X: 8, Y: 15}, pos.Absolute{X: 9,Y: 15})
	rect5 := eraser.Rectangle(pos.Absolute{X: 11, Y: 15}, pos.Absolute{X: 21,Y: 15})
	rect6 := eraser.Rectangle(pos.Absolute{X: 23, Y: 15}, pos.Absolute{X: 26,Y: 15})
	integrated2 := eraser.Integrated(rect3, rect4, rect5, rect6)
	base = model.NewSeatBase(3, 15, "휠체어석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         24,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('W'),
	}


	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	return group.Mixed(block1, block2)
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
	seats := ConcertHall1F() // Put Seating Here
	target := area.Area{
		Key:             "1F",
		Seats:           seats,
		XSize:           24,
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