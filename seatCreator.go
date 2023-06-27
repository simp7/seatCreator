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
	
	rect1 := eraser.Rectangle(pos.Absolute{X: 24, Y: 16}, pos.Absolute{X: 24,Y: 16})
	rect2 := eraser.Rectangle(pos.Absolute{X: 20, Y: 18}, pos.Absolute{X: 24,Y: 20})
	rect3 := eraser.Rectangle(pos.Absolute{X: 22, Y: 21}, pos.Absolute{X: 24,Y: 21})
	rect4 := eraser.Rectangle(pos.Absolute{X: 3, Y: 16}, pos.Absolute{X: 3,Y: 21})
	integrated1 := eraser.Integrated(rect1, rect2, rect3, rect4	)
	base := model.NewSeatBase(3, 16, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         22,
		YSize:         6,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(3, 2, "가석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         14,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(24, 3, "나석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         12,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.VerticalBlock(blockInput2)
	block3 := group.VerticalBlock(blockInput3)


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
		XSize:           22,
		YSize:           20,
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
