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

	hallH := eraser.HorizontalHallway(9)
	Arect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 3}, pos.Absolute{X: 6, Y: 3})
	Arect2 := eraser.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 5, Y: 4})
	Arect3 := eraser.Rectangle(pos.Absolute{X: 3, Y: 5}, pos.Absolute{X: 4, Y: 5})
	Aspecific := eraser.Position(pos.Absolute{X: 3, Y: 6})

	Brect1 := eraser.Rectangle(pos.Absolute{X: 20, Y: 20}, pos.Absolute{X: 27, Y: 20})

	Crect1 := eraser.Rectangle(pos.Absolute{X: 41, Y: 3}, pos.Absolute{X: 44, Y: 3})
	Crect2 := eraser.Rectangle(pos.Absolute{X: 42, Y: 4}, pos.Absolute{X: 44, Y: 4})
	Crect3 := eraser.Rectangle(pos.Absolute{X: 43, Y: 5}, pos.Absolute{X: 44, Y: 5})
	Cspecific := eraser.Position(pos.Absolute{X: 44, Y: 6})

	Wrect1 := eraser.Rectangle(pos.Absolute{X: 13, Y: 2}, pos.Absolute{X: 34, Y: 2})
	integrated1 := eraser.Integrated(hallH, Arect1,Arect2,Arect3, Aspecific, 
									Brect1,
									Crect1,Crect2,Crect3,Cspecific,
									Wrect1)
	base := model.NewSeatBase(3, 3, "가석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         18,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('A'),
		Reverse: true,
	}

	base = model.NewSeatBase(16, 3, "나석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         16,
		YSize:         18,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('B'),
		Reverse: true,
	}

	base = model.NewSeatBase(33, 3, "다석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         18,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('C'),
		Reverse: true,
	}

	base = model.NewSeatBase(9, 2, "휠체어석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         30,
		YSize:         1,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('W'),
		Reverse: true,
	}

	block1 := group.HorizontalBlock(blockInput1)
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
		Key:             "1F",
		Seats:           seats,
		XSize:           42,
		YSize:           18,
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
