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


	Arect1 := eraser.Rectangle(pos.Absolute{X: 6, Y: 16}, pos.Absolute{X: 10, Y: 16})
	Brect1 := eraser.Rectangle(pos.Absolute{X: 19, Y: 16}, pos.Absolute{X: 23, Y: 16})
	integrated1 := eraser.Integrated(Arect1, Brect1)
	base := model.NewSeatBase(3, 3, "가석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         14,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('A'),
		Reverse: true,
	}

	base = model.NewSeatBase(16, 3, "나석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         14,
		EmptyChecker:  integrated1,
		NameFormatter: nameformatter.Prefix('B'),
	}
	
	Wrect1 := eraser.Rectangle(pos.Absolute{X: 11, Y: 16}, pos.Absolute{X: 18, Y: 16})
	Wspecific := eraser.Position(pos.Absolute{X: 7, Y: 16}, pos.Absolute{X: 9, Y: 16},pos.Absolute{X: 20, Y: 16}, pos.Absolute{X: 22, Y: 16})
	integrated2 := eraser.Integrated(Wrect1,Wspecific)
	base = model.NewSeatBase(6, 16, "휠체어석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         18,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('W'),
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
		Key:             "1F",
		Seats:           seats,
		XSize:           23,
		YSize:           16,
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
