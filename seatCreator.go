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
	hall1 := eraser.VerticalHallway(14,35)
	hall2 := eraser.HorizontalHallway(16,17,18)

	rect1 := eraser.Rectangle(pos.Absolute{X: 2, Y: 25}, pos.Absolute{X: 14, Y: 26})
	rect2 := eraser.Rectangle(pos.Absolute{X: 35, Y: 25}, pos.Absolute{X: 47, Y: 26})

	integrated := eraser.Integrated(hall1,hall2,rect1,rect2)

	nameFormatter := nameformatter.Standard()

	base := model.NewSeatBase(2, 8, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         46,
		YSize:         17,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	integrated2 := eraser.Integrated()

	base = model.NewSeatBase(15, 25, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 13, 15),
		XSize:         20,
		YSize:         2,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	rect3 := eraser.Rectangle(pos.Absolute{X: 8, Y: 25}, pos.Absolute{X: 35, Y: 26})
	rect4 := eraser.Rectangle(pos.Absolute{X: 41, Y: 25}, pos.Absolute{X: 48, Y: 26})
	integrated3 := eraser.Integrated(rect3,rect4)

	base = model.NewSeatBase(2, 25, "휠체어석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         46,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
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
	seats := ConcertHall1F() // Put Seating Here
	target := area.Area{
		Key:             "1F",
		Seats:           seats,
		XSize:           46,
		YSize:           17,
  
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