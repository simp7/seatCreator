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

	hallH := eraser.HorizontalHallway(6)
	Arect1 := eraser.Rectangle(pos.Absolute{X: 11, Y: 7}, pos.Absolute{X: 14, Y: 7})
	Crect1 := eraser.Rectangle(pos.Absolute{X: 33, Y: 7}, pos.Absolute{X: 37, Y: 7})

	integrated1 := eraser.Integrated(hallH, Arect1, Crect1)
	base := model.NewSeatBase(3, 3, "라석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         5,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('A'),
		Reverse: true,
	}

	base = model.NewSeatBase(16, 3, "마석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         16,
		YSize:         3,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('B'),
		Reverse: true,
	}

	base = model.NewSeatBase(33, 3, "바석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         5,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('C'),
		Reverse: true,
	}

	

	block1 := group.HorizontalBlock(blockInput1)
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
		XSize:           42,
		YSize:           5,
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
