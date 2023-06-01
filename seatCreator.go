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

	rect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 8},pos.Absolute{X: 3,Y: 12})
	integrated := eraser.Integrated(rect1)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         9,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('A'),
	}

	base = model.NewSeatBase(9, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         13,
		YSize:         8,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('B'),
	}

	rect2 := eraser.Rectangle(pos.Absolute{X: 27,Y: 8},pos.Absolute{X: 27, Y: 12})
	integrated2 := eraser.Integrated(rect2)
	base = model.NewSeatBase(23, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         9,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('A'),
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
	seats := ArtriumSmall() // Put Seating Here
	target := area.Area{
		Key:             "2F",
		Seats:           seats,
		XSize:           28,
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