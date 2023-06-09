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

	rect1 := eraser.Rectangle(pos.Absolute{X: 5, Y: 14}, pos.Absolute{X: 5, Y: 16})
	rect2 := eraser.Rectangle(pos.Absolute{X: 6, Y: 15}, pos.Absolute{X: 9, Y: 15})
	rect3 := eraser.Rectangle(pos.Absolute{X: 6, Y: 16}, pos.Absolute{X: 7, Y: 16})
	integrated := eraser.Integrated(rect1, rect2, rect3)
	base := model.NewSeatBase(5, 13, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 14),
		XSize:         23,
		YSize:         4,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
		Reverse:       true,
	}

	base = model.NewSeatBase(3, 4, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         10,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('라'),
	}


	base = model.NewSeatBase(29, 4, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         10,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('다'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.VerticalBlock(blockInput2)
	block3 := group.VerticalBlock(blockInput3)
	return group.Mixed(block1,block2, block3)
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
		XSize:           29,
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
