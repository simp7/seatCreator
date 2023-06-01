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

	integrated := eraser.Integrated()

	base := model.NewSeatBase(3, 7, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         7,
		YSize:         13,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('A'),
	}

	rect := eraser.Rectangle(pos.Absolute{X: 13,Y: 19},pos.Absolute{X: 18,Y: 19})
	integrated2 := eraser.Integrated(rect)
		base = model.NewSeatBase(12, 7, "A석")
		blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         8,
		YSize:         13,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('B'),
	}

		base = model.NewSeatBase(22, 7, "A석")
		blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         7,
		YSize:         11,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('C'),
	}

		base = model.NewSeatBase(4, 2, "A석")
		blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         3,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('D'),
	}

		base = model.NewSeatBase(12, 2, "A석")
		blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         8,
		YSize:         3,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('E'),
	}

		base = model.NewSeatBase(22, 2, "A석")
		blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         3,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('F'),
	}

	rect2 := eraser.Rectangle(pos.Absolute{X: 14,Y: 19},pos.Absolute{X: 17,Y: 19})
	integrated3 := eraser.Integrated(rect2)
		base = model.NewSeatBase(13, 19, "휠체어석")
		blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7)
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
		XSize:           24,
		YSize:           11,
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