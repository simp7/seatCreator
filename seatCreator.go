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

	rect1 := eraser.Rectangle(pos.Absolute{X: 5, Y: 7}, pos.Absolute{X: 5, Y: 13})
	rect2 := eraser.Rectangle(pos.Absolute{X: 23, Y: 7}, pos.Absolute{X: 23, Y: 13})
	integrated := eraser.Integrated(rect1, rect2)

	base := model.NewSeatBase(5, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         19,
		YSize:         10,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	
	base = model.NewSeatBase(3, 4, "BOX1석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         13,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(4, 17, "BOX1석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 14, 1),
		XSize:         7,
		YSize:         1,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(25, 4, "BOX2석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         13,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(18, 17, "BOX2석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 14, 1),
		XSize:         7,
		YSize:         1,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
		Reverse: true,
	}

	base = model.NewSeatBase(1, 4, "BOX3석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         15,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(27, 4, "BOX4석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         15,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	rect3 := eraser.Rectangle(pos.Absolute{X: 8,Y: 19}, pos.Absolute{X: 20, Y: 19})
	integrated2 := eraser.Integrated(rect3)
	base = model.NewSeatBase(6, 19, "휠체어석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         17,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('W'),
	}
	
	

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.VerticalBlock(blockInput2)
	block3 := group.VerticalBlock(blockInput3)
	block4 := group.VerticalBlock(blockInput4)
	block5 := group.VerticalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)
	block8 := group.HorizontalBlock(blockInput8)
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7, block8)
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
		XSize:           27,
		YSize:           19,
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
