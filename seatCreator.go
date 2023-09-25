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
	
	hallV := eraser.VerticalHallway(7)
	rect1 := eraser.Rectangle(pos.Absolute{X: 19, Y: 4}, pos.Absolute{X: 19, Y: 15})
	specific := eraser.Position(pos.Absolute{X: 18, Y: 5}, pos.Absolute{X: 18, Y: 7}, pos.Absolute{X: 18, Y: 9})
	integrated1 := eraser.Integrated(hallV, rect1,specific)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         21,
		YSize:         4,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(5, 8, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 3, 5),
		XSize:         19,
		YSize:         1,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(4, 9, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 2, 6),
		XSize:         20,
		YSize:         1,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(3, 10, "A석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 7),
		XSize:         21,
		YSize:         6,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	specific1 := eraser.Position(pos.Absolute{X: 6, Y: 17})
	integrated2 := eraser.Integrated(specific1)
	base = model.NewSeatBase(3, 16, "A석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 13),
		XSize:         4,
		YSize:         2,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	specific2 := eraser.Position(pos.Absolute{X: 18, Y: 16})
	integrated3 := eraser.Integrated(specific2)
	base = model.NewSeatBase(9, 16, "A석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 6, 13),
		XSize:         15,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(12, 17, "A석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 8, 14),
		XSize:         12,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}



	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)

	return group.Mixed(block1, block2,block3,block4, block5, block6, block7)
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
		XSize:           21,
		YSize:           14,
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
