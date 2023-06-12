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
 	
	specific := eraser.Position(pos.Absolute{X: 4, Y: 10})
	rect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 10}, pos.Absolute{X: 3, Y: 13})
	integrated := eraser.Integrated(specific,rect1)
	base := model.NewSeatBase(3, 10, "D석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         16,
		YSize:         5,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	specific2 := eraser.Position(pos.Absolute{X: 34, Y: 10})
	rect2 := eraser.Rectangle(pos.Absolute{X: 35, Y: 10}, pos.Absolute{X: 35, Y: 13})
	integrated2 := eraser.Integrated(specific2, rect2)
	base = model.NewSeatBase(20, 10, "E석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         16,
		YSize:         5,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(2, 3, "L석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         6,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('L'),
	}
	
	base = model.NewSeatBase(36, 3, "R석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         6,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('R'),
	}
	
	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	return group.Mixed(block1, block2, block3, block4)
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
		XSize:           35,
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
